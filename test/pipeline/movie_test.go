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

package pipeline

import (
	"context"
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	p "github.com/GoogleCloudPlatform/solutions/media/pkg/pipeline"
	"github.com/GoogleCloudPlatform/solutions/media/test"
	"github.com/stretchr/testify/assert"
)

const DEFAULT_PROMPT = `
Review the attached movie and fulfill the following instructions and response in JSON format:
- Identify the movie's Title, Director, Producers, Cinematographers, and Actors.
- Write a creative and detailed summary that matches the tone of the movie.
- Identify all of the scenes in the movie start and end timestamps in the format of HH:mm:ss where hours (HH), minutes (mm), and seconds (ss) are only two digits.

Example:
	{
		"title": "",
		"summary": "",
		"duration_in_minutes": 60,
		"director": "",
		"executive_producers": [""],
		"producers": [""],
		"cinematographers": [""],
		"actors": [""],
		"scenes": [
			{"start": "00:00:00", "end": "00:00:00" }
		]
	}
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

	chain := p.MovieChain(cloudClients.GenAIClient, genModel, cloudClients.StorageClient, DEFAULT_PROMPT)

	chainCtx := model.NewChainContext()
	chainCtx.Add(model.CTX_IN, test.GetTestLowResMessageText())
	chainCtx.Add(model.CTX_PROMPT_VARS, make(map[string]interface{}))

	assert.True(t, chain.IsExecutable(chainCtx))

	chain.Execute(chainCtx)

	assert.False(t, chainCtx.HasErrors())

	fmt.Println(chainCtx.Get("DOC_SUMMARY"))
}
