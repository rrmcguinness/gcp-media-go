// Copyright 2024 Google, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pipeline

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"text/template"
	"time"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/google/generative-ai-go/genai"
)

type MovieReviewCommand struct {
	model.Command
	genaiClient   *genai.Client
	storageClient *storage.Client
	model         *genai.GenerativeModel
	template      *template.Template
}

func NewMovieReviewCommand(genaiClient *genai.Client, storageClient *storage.Client, model *genai.GenerativeModel, promptTemplate string) (cmd *MovieReviewCommand, err error) {
	template, err := template.New("why").Parse(promptTemplate)
	if err != nil {
		return nil, err
	}
	return &MovieReviewCommand{
		genaiClient:   genaiClient,
		storageClient: storageClient,
		model:         model,
		template:      template}, nil
}

func (m MovieReviewCommand) IsExecutable(chCtx model.ChainContext) bool {
	return chCtx.Get(model.CTX_MESSAGE).(*model.TriggerMediaWrite).Kind == model.EVENT_STORAGE_BUCKET_WRITE &&
		chCtx.Get(model.CTX_PROMPT_VARS).(map[string]interface{}) != nil
}

func (m MovieReviewCommand) Execute(chCtx model.ChainContext) {
	ctx := context.Background()

	promptVars := chCtx.Get(model.CTX_PROMPT_VARS).(map[string]interface{})
	var buffer bytes.Buffer
	err := m.template.Execute(&buffer, promptVars)
	if err != nil {
		chCtx.AddError(err)
		return
	}

	msg := chCtx.Get(model.CTX_MESSAGE).(*model.TriggerMediaWrite)

	outputFile, err := os.CreateTemp("", "vid-to-genai-")
	defer os.Remove(outputFile.Name())
	err = cloud.CopyFromGcsToTmp(ctx, m.storageClient, outputFile, msg.Bucket, msg.Name)
	if err != nil {
		chCtx.AddError(err)
		return
	}
	genFil, err := m.genaiClient.UploadFileFromPath(ctx, outputFile.Name(), &genai.UploadFileOptions{DisplayName: msg.Name, MIMEType: msg.ContentType})
	if err != nil {
		chCtx.AddError(err)
		return
	}
	defer m.genaiClient.DeleteFile(ctx, genFil.Name)
	fmt.Printf("Generated File: %s", genFil.URI)

	// Videos need to be processed before you can use them.
	for genFil.State == genai.FileStateProcessing {
		time.Sleep(5 * time.Second)
		var err error
		if genFil, err = m.genaiClient.GetFile(ctx, genFil.Name); err != nil {
			chCtx.AddError(err)
		}
	}

	parts := make([]genai.Part, 0)
	parts = append(parts, cloud.NewFileData(genFil.URI, msg.ContentType))
	parts = append(parts, cloud.NewTextPart(buffer.String()))

	fmt.Printf("%v", parts)

	out, err := cloud.GenerateMultiModalResponse(ctx, 0, *m.model, parts...)
	if err != nil {
		chCtx.AddError(err)
		return
	}
	chCtx.Add(model.CTX_PROMPT_RESP, out)
}
