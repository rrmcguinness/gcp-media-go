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
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/commands"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/google/generative-ai-go/genai"
)

func MediaIngestion(
	bigqueryClient *bigquery.Client,
	genaiClient *genai.Client,
	genaiModel *cloud.QuotaAwareModel,
	storageClient *storage.Client,
	summaryPromptTemplate string,
	summaryOutputParam string,
	scenePromptTemplate string,
	sceneOutputParam string,
	mediaOutputParam string,
	numberOfWorkers int,
) cor.Chain {

	out := cor.NewBaseChain("media-ingestion-workflow")

	// Convert the Message to an Object
	out.AddCommand(&commands.MediaTriggerToGCSObject{})

	// Write a temp file
	out.AddCommand(commands.NewGCSToTempFile("gcs-to-temp-file", storageClient, "media-summary-"))

	// Upload the file to file service
	out.AddCommand(commands.NewMediaUpload("media-upload", genaiClient, 300*time.Second))

	// Generate Summary
	out.AddCommand(commands.NewMediaPrompt("generate-media-summary", genaiModel, summaryPromptTemplate, cor.CTX_PROMPT_VARS))

	// Convert the JSON to a struct and save to the summaryOutputParam
	out.AddCommand(commands.NewMediaSummaryJsonToStruct("convert-media-summary", summaryOutputParam))

	// Create the scene extraction command
	sceneExtractor := commands.NewSceneExtractor("extract-media-scenes", genaiModel, scenePromptTemplate, numberOfWorkers)
	sceneExtractor.BaseCommand.OutputParamName = sceneOutputParam
	out.AddCommand(sceneExtractor)

	// Assemble the output into a single media object
	out.AddCommand(commands.NewMediaAssembly("assemble-media-scenes", summaryOutputParam, sceneOutputParam, mediaOutputParam))

	// Save media object to big query for async embedding job
	out.AddCommand(commands.NewMediaPersistToBigQuery("write-to-bigquery", bigqueryClient, "media_ds", "media", mediaOutputParam))

	// Clean up the temporary media created by the job
	out.AddCommand(commands.NewMediaCleanup("cleanup-file-system", genaiClient))

	// Return the chain for multiple executions
	return out
}
