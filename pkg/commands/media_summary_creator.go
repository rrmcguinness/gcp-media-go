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
	"fmt"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"go.opentelemetry.io/otel/metric"
	"text/template"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/google/generative-ai-go/genai"
)

type MediaSummaryCreator struct {
	cor.BaseCommand
	config                   *cloud.Config
	generativeAIModel        *cloud.QuotaAwareGenerativeAIModel
	template                 *template.Template
	geminiInputTokenCounter  metric.Int64Counter
	geminiOutputTokenCounter metric.Int64Counter
	geminiRetryCounter       metric.Int64Counter
}

func NewMediaSummaryCreator(
	name string,
	config *cloud.Config,
	generativeAIModel *cloud.QuotaAwareGenerativeAIModel,
	template *template.Template) *MediaSummaryCreator {

	out := &MediaSummaryCreator{
		BaseCommand:       *cor.NewBaseCommand(name),
		config:            config,
		generativeAIModel: generativeAIModel,
		template:          template}

	out.geminiInputTokenCounter, _ = out.GetMeter().Int64Counter(fmt.Sprintf("%s.gemini.token.input", out.GetName()))
	out.geminiOutputTokenCounter, _ = out.GetMeter().Int64Counter(fmt.Sprintf("%s.gemini.token.ouput", out.GetName()))
	out.geminiRetryCounter, _ = out.GetMeter().Int64Counter(fmt.Sprintf("%s.gemini.token.retry", out.GetName()))

	return out
}

func (t *MediaSummaryCreator) GenerateParams(_ cor.Context) map[string]interface{} {
	params := make(map[string]interface{})

	// Create a string representation of the categories
	catStr := ""
	for key, cat := range t.config.Categories {
		catStr += key + " - " + cat.Definition + "; "
	}
	params["CATEGORIES"] = t.config.Categories

	exampleSummary, _ := json.Marshal(model.GetExampleSummary())
	params["EXAMPLE_JSON"] = string(exampleSummary)
	return params
}

func (t *MediaSummaryCreator) Execute(context cor.Context) {
	mediaFile := context.Get(t.GetInputParam()).(*genai.File)
	var buffer bytes.Buffer
	err := t.template.Execute(&buffer, t.GenerateParams(context))
	if err != nil {
		t.GetErrorCounter().Add(context.GetContext(), 1)
		context.AddError(t.GetName(), err)
		return
	}

	// Create the parts to query Gemini
	parts := make([]genai.Part, 0)
	parts = append(parts, cloud.NewFileData(mediaFile.URI, mediaFile.MIMEType))
	parts = append(parts, cloud.NewTextPart(buffer.String()))

	// Get the response
	out, err := cloud.GenerateMultiModalResponse(context.GetContext(), t.geminiInputTokenCounter, t.geminiOutputTokenCounter, t.geminiRetryCounter, 0, t.generativeAIModel, parts...)
	if err != nil {
		t.GetErrorCounter().Add(context.GetContext(), 1)
		context.AddError(t.GetName(), err)
		return
	}
	t.GetSuccessCounter().Add(context.GetContext(), 1)
	context.Add(t.GetOutputParam(), out)
}
