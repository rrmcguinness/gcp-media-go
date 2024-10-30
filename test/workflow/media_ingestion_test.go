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

package workflow_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/workflow"
	"github.com/GoogleCloudPlatform/solutions/media/test"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/codes"
)

const DEFAULT_PROMPT = `
Review the attached media file and extract the following information
- Title as title
- Summary - a detailed summary of the media contents, plot, and cinematic themes in markdown format
- Director as director
- Release Year as release_year, a four digit year
- Genre as genre
- Rating as rating with one of the following values: G, PG, PG-13, R, NC-17
- Cast as cast, an array of Cast Members including Character Name as character_name, and associated actor name as actor_name
- Extract the scenes and order by start and end times in the format of HH:MM:SS or hours:minutes:seconds as two digits each
- Add a sequence number to each scene starting from 1 and incrementing in order of the timestamp

Example Output:
%s
`

const SCENE_PROMPT = `
Given the following media file, summary, actors, and characters, extract the following details between time frames {{ .TIME_START }} - {{ .TIME_END }} in json format:
- sequence_number: {{ .SEQUENCE }}
- start: {{ .TIME_START }}
- end: {{ .TIME_END }}
- script, the details of the scene using a modified screen play script format that includes the actors name on the character line. Example: CHARACTER NAME (V.O.) - (ACTOR NAME)

Media Summary:
{{ .SUMMARY_DOCUMENT }}

Example Output:
{{ .EXAMPLE_JSON }}
`

func TestMediaChain(t *testing.T) {
	traceCtx, span := tracer.Start(ctx, "media-ingestion-test")
	defer span.End()

	jsonData, _ := json.Marshal(model.GetExampleSummary())
	prompt := fmt.Sprintf(DEFAULT_PROMPT, jsonData)

	mediaIngestWorkflow := workflow.MediaIngestion(
		cloudClients.BiqQueryClient,
		cloudClients.GenAIClient,
		genModel,
		cloudClients.StorageClient,
		prompt, "DOC_SUMMARY",
		SCENE_PROMPT,
		"SCENES",
		"MEDIA",
		config.Application.ThreadPoolSize)

	chainCtx := cor.NewBaseContext()
	chainCtx.SetContext(traceCtx)
	chainCtx.Add(cor.CTX_IN, test.GetTestLowResMessageText())
	chainCtx.Add(cor.CTX_PROMPT_VARS, make(map[string]interface{}))

	assert.True(t, mediaIngestWorkflow.IsExecutable(chainCtx))

	mediaIngestWorkflow.Execute(chainCtx)

	for _, err := range chainCtx.GetErrors() {
		fmt.Println(err.Error())
	}

	if chainCtx.HasErrors() {
		span.SetStatus(codes.Error, "failed to execute media ingestion test")
	}

	assert.False(t, chainCtx.HasErrors())

	span.SetStatus(codes.Ok, "passed - media ingestion test")

	fmt.Println(chainCtx.Get("MEDIA"))
}
