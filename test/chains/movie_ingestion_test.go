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
	"context"
	"encoding/json"
	"fmt"
	"testing"

	parent_chains "github.com/GoogleCloudPlatform/solutions/media/pkg/chains"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/GoogleCloudPlatform/solutions/media/test"
	"github.com/stretchr/testify/assert"
)

const DEFAULT_PROMPT = `
Review the attached movie and extract the following information
- Title as title
- Summary - a detailed summary of the movie, plot, and cinematic universe of the movie in markdown format
- Director as director
- Release Year as release_year, a four digit year
- Genre as genre
- Rating as rating with one of the following values: G, PG, PG-13, R, NC-17
- Cast as cast, an array of Cast Members including Character Name as character_name, and associated actor name as actor_name
- Extract the scenes and their start and end times in the format of HH:MM:SS or hours:minutes:seconds as two digits each

Example Output:
%s
`

const SCENE_PROMPT = `
Given the following movie, movie summary, actors, and characters, extract the details of the scene using screen play script format between time frames {{ .TIME_START }} - {{ .TIME_END }}.

Movie Sequence: {{ .SEQUENCE }}
Movie Summary:
{{ .SUMMARY_DOCUMENT }}

Example Output:
{{ .EXAMPLE_JSON }}
`

func TestMediaChain(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// This deferral will automatically close the client that was build from
	// the same context
	defer cancel()

	// Get the config file
	config := test.GetConfig(t)

	cloudClients, err := cloud.NewCloudServiceClients(ctx, config)
	test.HandleErr(err, t)
	defer cloudClients.Close()

	genModel := cloudClients.AgentModels["creative-flash"]
	assert.NotNil(t, genModel)

	jsonData, _ := json.Marshal(model.GetExampleSummary())
	prompt := fmt.Sprintf(DEFAULT_PROMPT, jsonData)

	chain := parent_chains.MovieIngestionChain(
		cloudClients.BiqQueryClient,
		cloudClients.GenAIClient,
		genModel,
		cloudClients.StorageClient,
		prompt, "DOC_SUMMARY",
		SCENE_PROMPT,
		"SCENES",
		"MOVIE")

	chainCtx := cor.NewBaseContext()
	chainCtx.Add(cor.CTX_IN, test.GetTestLowResMessageText())
	chainCtx.Add(cor.CTX_PROMPT_VARS, make(map[string]interface{}))

	assert.True(t, chain.IsExecutable(chainCtx))

	chain.Execute(chainCtx)

	for _, err := range chainCtx.GetErrors() {
		fmt.Println(err.Error())
	}

	assert.False(t, chainCtx.HasErrors())

	fmt.Println(chainCtx.Get("MOVIE"))
}
