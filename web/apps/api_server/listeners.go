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

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/workflow"
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
Given the following media file, summary, actors, and characters, extract the following details between time frames {{ .TIME_START }} - {{ .TIME_END }} in json format.
- sequence_number: {{ .SEQUENCE }} as a number
- start: {{ .TIME_START }} as a string
- end: {{ .TIME_END }} as a string
- script, the details of the scene using a modified screen play script format that includes the actors name on the character line. Example: CHARACTER NAME (V.O.) - (ACTOR NAME) as a string

Media Summary:
{{ .SUMMARY_DOCUMENT }}

Example Output:
{{ .EXAMPLE_JSON }}
`

func SetupListeners(config *cloud.CloudConfig, cloudClients *cloud.CloudServiceClients, ctx context.Context) {
	// TODO - Externalize the destination topic and ffmpeg command
	mediaResizeWorkflow := workflow.MediaResize("bin/ffmpeg", &model.MediaFormatFilter{Width: "240"}, cloudClients.StorageClient, "media_low_res_resources")
	cloudClients.PubSubListeners["HiResTopic"].SetCommand(mediaResizeWorkflow)
	cloudClients.PubSubListeners["HiResTopic"].Listen(ctx)

	jsonData, _ := json.Marshal(model.GetExampleSummary())
	prompt := fmt.Sprintf(DEFAULT_PROMPT, jsonData)

	// TODO - Externalize prompt
	mediaIngestWorkflow := workflow.MediaIngestion(
		cloudClients.BiqQueryClient,
		cloudClients.GenAIClient,
		cloudClients.AgentModels["creative-flash"],
		cloudClients.StorageClient,
		prompt, "DOC_SUMMARY",
		SCENE_PROMPT,
		"SCENES",
		"MEDIA",
		config.Application.ThreadPoolSize)

	cloudClients.PubSubListeners["LowResTopic"].SetCommand(mediaIngestWorkflow)
	cloudClients.PubSubListeners["LowResTopic"].Listen(ctx)
}
