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
	"encoding/json"
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	p "github.com/GoogleCloudPlatform/solutions/media/pkg/pipeline"
	"github.com/GoogleCloudPlatform/solutions/media/test"
	"github.com/stretchr/testify/assert"
)

func TestMediaReviewCommand(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// This deferral will automatically close the client that was build from
	// the same context
	defer cancel()

	// Get the config file
	config := test.GetConfig(t)

	cloudClients, err := cloud.NewCloudServiceClients(ctx, config)
	test.HandleErr(err, t)
	defer cloudClients.Close()

	// Create a "fake" message
	var out model.TriggerMediaWrite
	err = json.Unmarshal([]byte(test.GetTestLowResMessageText()), &out)
	test.HandleErr(err, t)

	// Basic assertions on nil and type
	assert.NotNil(t, out)
	assert.Equal(t, model.EVENT_STORAGE_BUCKET_WRITE, out.Kind)

	// Create the context
	chainCtx := model.NewChainContext()
	chainCtx.Add(model.CTX_MESSAGE, &out)

	chainCtx.Add(model.CTX_PROMPT_VARS, make(map[string]interface{}))

	promptTemplate := `
	Review the attached movie and fulfill the following instructions and response in JSON format:
	- Identify the movie's Title, Director, Producers, Cinematographers, and Actors.
	- Write a creative and detailed summary that matches the tone of the movie.
	- Identify all of the scenes in the movie start and end timestamps.

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
	genModel := cloudClients.AgentModels["creative-flash"]
	assert.NotNil(t, genModel)
	movieReviewCommand, err := p.NewMovieReviewCommand(genModel, promptTemplate)
	test.HandleErr(err, t)

	assert.True(t, movieReviewCommand.IsExecutable(chainCtx))

	movieReviewCommand.Execute(chainCtx)

	for _, err := range chainCtx.GetErrors() {
		fmt.Println(err.Error())
	}

	assert.False(t, chainCtx.HasErrors())

	fmt.Println(chainCtx.Get(model.CTX_PROMPT_RESP))

}
