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

package commands

import (
	"bytes"
	go_ctx "context"
	"text/template"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/google/generative-ai-go/genai"
)

type MediaPrompt struct {
	cor.BaseCommand
	GenaiClient        *genai.Client
	GenaiModel         *genai.GenerativeModel
	PromptTemplate     string
	TemplateParamsName string
}

func (t *MediaPrompt) Execute(context cor.Context) {

	ctx := go_ctx.Background()

	mediaFile := context.Get(t.GetInputParam()).(*genai.File)
	params := context.Get(t.TemplateParamsName).(map[string]interface{})
	template, err := template.New("why").Parse(t.PromptTemplate)
	if err != nil {
		context.AddError(err)
		return
	}

	var buffer bytes.Buffer
	err = template.Execute(&buffer, params)
	if err != nil {
		context.AddError(err)
		return
	}

	parts := make([]genai.Part, 0)
	parts = append(parts, cloud.NewFileData(mediaFile.URI, mediaFile.MIMEType))
	parts = append(parts, cloud.NewTextPart(buffer.String()))

	out, err := cloud.GenerateMultiModalResponse(ctx, 0, t.GenaiModel, parts...)
	if err != nil {
		context.AddError(err)
		return
	}
	// Make retrievable for later use
	context.Add(t.GetOutputParam(), out)
	// Still pipe to next command
	context.Add(cor.CTX_OUT, out)
}
