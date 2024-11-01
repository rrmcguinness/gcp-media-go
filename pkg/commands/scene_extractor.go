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
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"text/template"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/google/generative-ai-go/genai"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type SceneExtractor struct {
	cor.BaseCommand
	model           *cloud.QuotaAwareModel
	prompt          string
	numberOfWorkers int
}

func NewSceneExtractor(name string, model *cloud.QuotaAwareModel, prompt string, numberOfWorkers int) *SceneExtractor {
	return &SceneExtractor{BaseCommand: *cor.NewBaseCommand(name), model: model, prompt: prompt, numberOfWorkers: numberOfWorkers}
}

func (s *SceneExtractor) IsExecutable(context cor.Context) bool {
	return context != nil &&
		context.Get(s.GetInputParam()) != nil &&
		context.Get(GetVideoUploadFileParameterName()) != nil
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
	template, _ := template.New("scene_template").Parse(s.prompt)

	var wg sync.WaitGroup
	jobs := make(chan *SceneJob, len(summary.SceneTimeStamps))
	results := make(chan *SceneResponse, len(summary.SceneTimeStamps))

	// Create worker pool
	for w := 1; w <= s.numberOfWorkers; w++ {
		wg.Add(1)
		go scene_worker(jobs, results, &wg)
	}

	// Execute all scenes against the worker pool
	for i, ts := range summary.SceneTimeStamps {
		job := CreateJob(context.GetContext(), s.Tracer, i, s.GetName(), summaryText, exampleText, *template, videoFile, s.model, ts)
		jobs <- job
	}

	close(jobs)
	wg.Wait()
	close(results)

	// Aggregate the responses
	sceneData := make([]string, 0)
	for r := range results {
		if r.err != nil {
			context.AddError(r.err)
		} else {
			sceneData = append(sceneData, r.value)
		}
	}
	context.Add(s.GetOutputParam(), sceneData)
	context.Add(cor.CTX_OUT, sceneData)
}

type SceneResponse struct {
	value string
	err   error
}

type SceneJob struct {
	workerId int
	ctx      go_ctx.Context
	timeSpan *model.TimeSpan
	span     trace.Span
	parts    []genai.Part
	model    *cloud.QuotaAwareModel
}

func (s *SceneJob) Close(status codes.Code, description string) {
	s.span.SetStatus(status, description)
	s.span.End()
}

func CreateJob(
	ctx go_ctx.Context,
	tracer trace.Tracer,
	workerId int,
	commandName string,
	summaryText string,
	exampleText string,
	template template.Template,
	videoFile *genai.File,
	model *cloud.QuotaAwareModel,
	timeSpan *model.TimeSpan,
) *SceneJob {
	sceneCtx, sceneSpan := tracer.Start(ctx, fmt.Sprintf("%s_genai", commandName))
	sceneSpan.SetAttributes(
		attribute.Int("sequence", workerId),
		attribute.String("start", timeSpan.Start),
		attribute.String("end", timeSpan.End),
	)

	vocabulary := make(map[string]string)
	vocabulary["SEQUENCE"] = fmt.Sprintf("%d", workerId)
	vocabulary["SUMMARY_DOCUMENT"] = summaryText
	vocabulary["TIME_START"] = timeSpan.Start
	vocabulary["TIME_END"] = timeSpan.End
	vocabulary["EXAMPLE_JSON"] = exampleText

	var doc bytes.Buffer
	template.Execute(&doc, vocabulary)
	tsPrompt := doc.String()

	parts := make([]genai.Part, 0)
	parts = append(parts, cloud.NewFileData(videoFile.URI, videoFile.MIMEType))
	parts = append(parts, cloud.NewTextPart(tsPrompt))

	return &SceneJob{workerId: workerId, ctx: sceneCtx, timeSpan: timeSpan, span: sceneSpan, parts: parts, model: model}
}

// Create a worker function for parallel work streams
func scene_worker(jobs <-chan *SceneJob, results chan<- *SceneResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		out, err := cloud.GenerateMultiModalResponse(j.ctx, 0, j.model, j.parts...)
		if err != nil {
			j.Close(codes.Error, "scene extract failed")
			results <- &SceneResponse{err: err}
			return
		}
		if len(strings.Trim(out, " ")) > 0 && out != "{}" {
			results <- &SceneResponse{value: out, err: nil}
		}
		j.Close(codes.Ok, "completed scene")
	}
}
