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
	"text/template"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/google/generative-ai-go/genai"
)

type MovieReviewCommand struct {
	model.Command
	model    *genai.GenerativeModel
	template *template.Template
}

func NewMovieReviewCommand(model *genai.GenerativeModel, promptTemplate string) (cmd *MovieReviewCommand, err error) {
	template, err := template.New("why").Parse(promptTemplate)
	if err != nil {
		return nil, err
	}
	return &MovieReviewCommand{
		model:    model,
		template: template}, nil
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

	parts := make([]genai.Part, 0)
	parts = append(parts, cloud.NewFileData(msg.GetGSUri(), msg.ContentType))
	parts = append(parts, cloud.NewTextPart(buffer.String()))

	fmt.Printf("%v", parts)

	out, err := cloud.GenerateMultiModalResponse(ctx, 0, *m.model, parts...)
	if err != nil {
		chCtx.AddError(err)
		return
	}
	chCtx.Add(model.CTX_PROMPT_RESP, out)
}
