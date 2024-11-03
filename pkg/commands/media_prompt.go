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
	"encoding/json"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"text/template"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/google/generative-ai-go/genai"
)

type MediaPrompt struct {
	cor.BaseCommand
	generativeAIModel *cloud.QuotaAwareGenerativeAIModel
	template          *template.Template
	params            map[string]interface{}
}

func NewMediaPrompt(name string,
	generativeAIModel *cloud.QuotaAwareGenerativeAIModel,
	template *template.Template) *MediaPrompt {

	exampleSummary, _ := json.Marshal(model.GetExampleSummary())

	out := &MediaPrompt{
		BaseCommand:       *cor.NewBaseCommand(name),
		generativeAIModel: generativeAIModel,
		template:          template,
		params:            make(map[string]interface{})}

	out.params["EXAMPLE_JSON"] = string(exampleSummary)
	return out
}

func (t *MediaPrompt) Execute(context cor.Context) {
	mediaFile := context.Get(t.GetInputParam()).(*genai.File)

	var buffer bytes.Buffer
	err := t.template.Execute(&buffer, t.params)
	if err != nil {
		context.AddError(err)
		return
	}

	// Create the parts to query Gemini
	parts := make([]genai.Part, 0)
	parts = append(parts, cloud.NewFileData(mediaFile.URI, mediaFile.MIMEType))
	parts = append(parts, cloud.NewTextPart(buffer.String()))

	// Get the response
	out, err := cloud.GenerateMultiModalResponse(context.GetContext(), 0, t.generativeAIModel, parts...)
	if err != nil {
		context.AddError(err)
		return
	}
	context.Add(t.GetOutputParam(), out)
}
