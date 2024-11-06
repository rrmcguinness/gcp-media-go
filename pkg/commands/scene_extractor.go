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
	goctx "context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel/metric"
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
	generativeAIModel        *cloud.QuotaAwareGenerativeAIModel
	promptTemplate           *template.Template
	numberOfWorkers          int
	geminiInputTokenCounter  metric.Int64Counter
	geminiOutputTokenCounter metric.Int64Counter
	geminiRetryCounter       metric.Int64Counter
}

func NewSceneExtractor(
	name string,
	model *cloud.QuotaAwareGenerativeAIModel,
	prompt *template.Template,
	numberOfWorkers int) *SceneExtractor {
	out := &SceneExtractor{
		BaseCommand:       *cor.NewBaseCommand(name),
		generativeAIModel: model,
		promptTemplate:    prompt,
		numberOfWorkers:   numberOfWorkers}

	out.geminiInputTokenCounter, _ = out.GetMeter().Int64Counter(fmt.Sprintf("%s.gemini.token.input", out.GetName()))
	out.geminiOutputTokenCounter, _ = out.GetMeter().Int64Counter(fmt.Sprintf("%s.gemini.token.ouput", out.GetName()))
	out.geminiRetryCounter, _ = out.GetMeter().Int64Counter(fmt.Sprintf("%s.gemini.token.retry", out.GetName()))

	return out
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

	// Create a human-readable cast
	castString := ""
	for _, cast := range summary.Cast {
		castString += fmt.Sprintf("%s - %s\n", cast.CharacterName, cast.ActorName)
	}
	summaryText := fmt.Sprintf("Title:%s\nSummary:\n\n%s\nCast:\n\n%v\n", summary.Title, summary.Summary, castString)

	var wg sync.WaitGroup
	jobs := make(chan *SceneJob, len(summary.SceneTimeStamps))
	results := make(chan *SceneResponse, len(summary.SceneTimeStamps))

	// Create worker pool
	for w := 1; w <= s.numberOfWorkers; w++ {
		wg.Add(1)
		go sceneWorker(jobs, results, &wg)
	}

	// Execute all scenes against the worker pool
	for i, ts := range summary.SceneTimeStamps {
		job := CreateJob(context.GetContext(), s.Tracer, s.geminiInputTokenCounter, s.geminiOutputTokenCounter, s.geminiRetryCounter, i, s.GetName(), summaryText, exampleText, *s.promptTemplate, videoFile, s.generativeAIModel, ts)
		jobs <- job
	}

	close(jobs)
	wg.Wait()
	close(results)

	// Aggregate the responses
	sceneData := make([]string, 0)
	for r := range results {
		if r.err != nil {
			s.GetErrorCounter().Add(context.GetContext(), 1)
			context.AddError(s.GetName(), r.err)
		} else {

			sceneData = append(sceneData, r.value)
		}
	}

	if !context.HasErrors() {
		s.GetSuccessCounter().Add(context.GetContext(), 1)
	}

	context.Add(s.GetOutputParam(), sceneData)
	context.Add(cor.CtxOut, sceneData)
}

type SceneResponse struct {
	value string
	err   error
}

type SceneJob struct {
	workerId                 int
	ctx                      goctx.Context
	geminiInputTokenCounter  metric.Int64Counter
	geminiOutputTokenCounter metric.Int64Counter
	geminiRetryCounter       metric.Int64Counter
	timeSpan                 *model.TimeSpan
	span                     trace.Span
	parts                    []genai.Part
	model                    *cloud.QuotaAwareGenerativeAIModel
	err                      error
}

func (s *SceneJob) Close(status codes.Code, description string) {
	s.span.SetStatus(status, description)
	s.span.End()
}

func CreateJob(
	ctx goctx.Context,
	tracer trace.Tracer,
	geminiInputTokenCounter metric.Int64Counter,
	geminiOutputTokenCounter metric.Int64Counter,
	geminiRetryCounter metric.Int64Counter,
	workerId int,
	commandName string,
	summaryText string,
	exampleText string,
	template template.Template,
	videoFile *genai.File,
	model *cloud.QuotaAwareGenerativeAIModel,
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
	err := template.Execute(&doc, vocabulary)
	if err != nil {
		return &SceneJob{err: err}
	}
	tsPrompt := doc.String()

	parts := make([]genai.Part, 0)
	parts = append(parts, cloud.NewFileData(videoFile.URI, videoFile.MIMEType))
	parts = append(parts, cloud.NewTextPart(tsPrompt))

	return &SceneJob{workerId: workerId,
		ctx:                      sceneCtx,
		geminiInputTokenCounter:  geminiInputTokenCounter,
		geminiOutputTokenCounter: geminiOutputTokenCounter,
		geminiRetryCounter:       geminiRetryCounter,
		timeSpan:                 timeSpan, span: sceneSpan, parts: parts, model: model}
}

// Create a worker function for parallel work streams
func sceneWorker(jobs <-chan *SceneJob, results chan<- *SceneResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		if j.err == nil {
			out, err := cloud.GenerateMultiModalResponse(j.ctx, j.geminiInputTokenCounter, j.geminiOutputTokenCounter, j.geminiRetryCounter, 0, j.model, j.parts...)
			if err != nil {
				j.Close(codes.Error, "scene extract failed")
				results <- &SceneResponse{err: err}
				return
			}
			if len(strings.Trim(out, " ")) > 0 && out != "{}" {
				results <- &SceneResponse{value: out, err: nil}
			}
			j.Close(codes.Ok, "completed scene")
		} else {
			results <- &SceneResponse{value: "", err: j.err}
		}
	}
}
