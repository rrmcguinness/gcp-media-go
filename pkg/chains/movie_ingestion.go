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

package chains

import (
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/commands"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/google/generative-ai-go/genai"
)

func MovieIngestionChain(
	bigqueryClient *bigquery.Client,
	genaiClient *genai.Client,
	genaiModel *genai.GenerativeModel,
	storageClient *storage.Client,
	summaryPromptTemplate string,
	summaryOutputParam string,
	scenePromptTemplate string,
	sceneOutputParam string,
	movieOutputParam string,
) cor.Chain {

	out := &cor.BaseChain{}

	// Convert the Message to an Object
	out.AddCommand(&commands.MediaTriggerToGCSObject{})

	// Write a temp file
	out.AddCommand(&commands.GCSToTempFileCommand{Client: storageClient, TempFilePrefix: "movie-summary-"})

	// Upload the file to file service
	out.AddCommand(&commands.VideoUploadCommand{GenaiClient: genaiClient, TimeoutInSeconds: 300 * time.Second})

	// Generate Summary
	out.AddCommand(
		&commands.MediaPromptCommand{
			BaseCommand:        cor.BaseCommand{OutputParamName: summaryOutputParam},
			GenaiClient:        genaiClient,
			GenaiModel:         genaiModel,
			PromptTemplate:     summaryPromptTemplate,
			TemplateParamsName: cor.CTX_PROMPT_VARS,
		})

	// Convert the JSON to a struct
	out.AddCommand(&commands.SummaryResponseToStruct{})

	out.AddCommand(&commands.SceneBuilder{
		BaseCommand: cor.BaseCommand{OutputParamName: sceneOutputParam},
		GenaiClient: genaiClient,
		GenaiModel:  genaiModel,
		ScenePrompt: scenePromptTemplate})

	out.AddCommand(&commands.MovieAssembly{
		SummaryParameterName:     summaryOutputParam,
		SceneParameterName:       sceneOutputParam,
		MovieObjectParameterName: movieOutputParam,
	})

	out.AddCommand(&commands.SaveMovieToBQ{BigQueryClient: bigqueryClient,
		DataSetName:          "media_ds",
		TableName:            "movies",
		MovieObjectParamName: movieOutputParam})

	out.AddCommand(&commands.VideoCleanupCommand{GenaiClient: genaiClient})

	return out
}
