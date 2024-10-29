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
	"strings"
	"text/template"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/google/generative-ai-go/genai"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type SceneExtractor struct {
	cor.BaseCommand
	model  *genai.GenerativeModel
	prompt string
}

func NewSceneExtractor(name string, model *genai.GenerativeModel, prompt string) *SceneExtractor {
	return &SceneExtractor{BaseCommand: *cor.NewBaseCommand(name), model: model, prompt: prompt}
}

func (s *SceneExtractor) Execute(context cor.Context) {
	summary := context.Get(s.GetInputParam()).(*model.MediaSummary)
	videoFile := context.Get(GetVideoUploadFileParameterName()).(*genai.File)

	exampleScene := model.GetExampleScene()
	exampleJson, _ := json.Marshal(exampleScene)
	exampleText := string(exampleJson)

	// Create a human readable cast
	castString := ""
	for _, cast := range summary.Cast {
		castString += fmt.Sprintf("%s - %s\n", cast.CharacterName, cast.ActorName)
	}
	summaryText := fmt.Sprintf("Title:%s\nSummary:\n\n%s\nCast:\n\n%v\n", summary.Title, summary.Summary, castString)

	sceneData := make([]string, 0)

	template, _ := template.New("scene_templtae").Parse(s.prompt)

	for i, ts := range summary.SceneTimeStamps {
		sceneCtx, sceneSpan := s.Tracer.Start(context.GetContext(), fmt.Sprintf("%s_genai", s.GetName()))
		sceneSpan.SetAttributes(
			attribute.Int("sequence", i),
			attribute.String("start", ts.Start),
			attribute.String("end", ts.End),
		)
		// Build the vocabulary that MAY be used by the template
		vocabulary := make(map[string]string)
		vocabulary["SEQUENCE"] = fmt.Sprintf("%d", i)
		vocabulary["SUMMARY_DOCUMENT"] = summaryText
		vocabulary["TIME_START"] = ts.Start
		vocabulary["TIME_END"] = ts.End
		vocabulary["EXAMPLE_JSON"] = exampleText

		var doc bytes.Buffer
		template.Execute(&doc, vocabulary)
		tsPrompt := doc.String()

		parts := make([]genai.Part, 0)
		parts = append(parts, cloud.NewFileData(videoFile.URI, videoFile.MIMEType))
		parts = append(parts, cloud.NewTextPart(tsPrompt))

		out, err := cloud.GenerateMultiModalResponse(sceneCtx, 0, s.model, parts...)
		if err != nil {
			sceneSpan.SetStatus(codes.Error, "scene extract failed")
			context.AddError(err)
			sceneSpan.End()
			return
		}
		if len(strings.Trim(out, " ")) > 0 && out != "{}" {
			sceneData = append(sceneData, out)
		}
		sceneSpan.SetStatus(codes.Ok, "scene extract complete")
		sceneSpan.End()
	}
	context.Add(s.GetOutputParam(), sceneData)
}
