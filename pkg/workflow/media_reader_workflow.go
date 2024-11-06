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

package workflow

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/commands"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/google/generative-ai-go/genai"
	"text/template"
	"time"
)

type MediaReaderWorkflow struct {
	cor.BaseCommand
	config          *cloud.Config
	bigqueryClient  *bigquery.Client
	genaiClient     *genai.Client
	genaiModel      *cloud.QuotaAwareGenerativeAIModel
	storageClient   *storage.Client
	numberOfWorkers int
	summaryTemplate *template.Template
	sceneTemplate   *template.Template
	chain           cor.Chain
}

func (m *MediaReaderWorkflow) Execute(context cor.Context) {
	m.chain.Execute(context)
}

func (m *MediaReaderWorkflow) initializeChain() {
	const SummaryOutputParamName = "__summary_output__"
	const SceneOutputParamName = "__scene_output__"
	const MediaOutputParamName = "__media_output__"

	out := cor.NewBaseChain(m.GetName())

	// Convert the Message to an Object
	out.AddCommand(commands.NewMediaTriggerToGCSObject("media-trigger-to-gcs-object"))

	// Write a temp file
	out.AddCommand(commands.NewGCSToTempFile("gcs-to-temp-file", m.storageClient, "media-summary-"))

	// Upload the file to file service
	out.AddCommand(commands.NewMediaUpload("media-upload", m.genaiClient, 300*time.Second))

	// Generate Summary
	out.AddCommand(commands.NewMediaSummaryCreator("generate-media-summary", m.config, m.genaiModel, m.summaryTemplate))

	// Convert the JSON to a struct and save to the summaryOutputParam
	out.AddCommand(commands.NewMediaSummaryJsonToStruct("convert-media-summary", SummaryOutputParamName))

	// Create the scene extraction command
	sceneExtractor := commands.NewSceneExtractor("extract-media-scenes", m.genaiModel, m.sceneTemplate, m.numberOfWorkers)
	sceneExtractor.BaseCommand.OutputParamName = SceneOutputParamName
	out.AddCommand(sceneExtractor)

	// Assemble the output into a single media object
	out.AddCommand(commands.NewMediaAssembly("assemble-media-scenes", SummaryOutputParamName, SceneOutputParamName, MediaOutputParamName))

	// Save media object to big query for async embedding job
	out.AddCommand(commands.NewMediaPersistToBigQuery(
		"write-to-bigquery",
		m.bigqueryClient,
		m.config.BigQueryDataSource.DatasetName,
		m.config.BigQueryDataSource.MediaTable, MediaOutputParamName))

	// Clean up the temporary media created by the job
	out.AddCommand(commands.NewMediaCleanup("cleanup-file-system", m.genaiClient))
	m.chain = out
}

func NewMediaReaderPipeline(
	config *cloud.Config,
	serviceClients *cloud.ServiceClients,
	agentModelName string) *MediaReaderWorkflow {

	summaryTemplate, err := template.New("summary-template").Parse(config.PromptTemplates.SummaryPrompt)
	if err != nil {
		panic(err)
	}
	sceneTemplate, err := template.New("scene-template").Parse(config.PromptTemplates.ScenePrompt)
	if err != nil {
		panic(err)
	}

	pipeline := &MediaReaderWorkflow{
		BaseCommand:     *cor.NewBaseCommand("media-reader-pipeline"),
		config:          config,
		bigqueryClient:  serviceClients.BiqQueryClient,
		genaiClient:     serviceClients.GenAIClient,
		genaiModel:      serviceClients.AgentModels[agentModelName],
		storageClient:   serviceClients.StorageClient,
		numberOfWorkers: config.Application.ThreadPoolSize,
		summaryTemplate: summaryTemplate,
		sceneTemplate:   sceneTemplate,
	}
	pipeline.initializeChain()
	return pipeline
}
