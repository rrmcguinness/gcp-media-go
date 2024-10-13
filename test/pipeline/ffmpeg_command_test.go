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

	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	p "github.com/GoogleCloudPlatform/solutions/media/pkg/pipeline"
	"github.com/stretchr/testify/assert"
)

func TestFFMpegCommand(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// This deferral will automatically close the client that was build from
	// the same context
	defer cancel()

	// Get the config file
	config := GetConfig(t)

	// Create a "fake" message
	var out model.TriggerMediaWrite
	err := json.Unmarshal([]byte(GetTestHighResMessageText()), &out)
	HandleErr(err, t)

	// Basic assertions on nil and type
	assert.NotNil(t, out)
	assert.Equal(t, model.EVENT_STORAGE_BUCKET_WRITE, out.Kind)

	// Create the context
	chainCtx := model.NewChainContext()
	chainCtx.Add(model.CTX_MESSAGE, &out)

	// Create the command
	cmd := p.NewFFMpegStreamingCommand(
		ctx,
		"pkg/pipeline/bin/ffmpeg",
		config.HighResToLowResCommand.Bucket,
		config.HighResToLowResCommand.Format,
		config.HighResToLowResCommand.Width)

	// This assertion insures the command can be executed
	assert.True(t, cmd.IsExecutable(chainCtx))
	cmd.Execute(chainCtx)

	for _, err := range chainCtx.GetErrors() {
		fmt.Println(err.Error())
	}

	assert.False(t, chainCtx.HasErrors())
}
