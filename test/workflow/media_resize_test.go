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
	"context"
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/workflow"
	"github.com/GoogleCloudPlatform/solutions/media/test"
	"github.com/stretchr/testify/assert"
)

func TestFFMpegCommand(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// This deferral will automatically close the client that was build from
	// the same context
	defer cancel()

	// Get the config file
	config := test.GetConfig(t)

	cloud, err := cloud.NewCloudServiceClients(ctx, config)
	test.HandleErr(err, t)
	defer cloud.Close()

	// Create the context
	chainCtx := cor.NewBaseContext()
	chainCtx.Add(cor.CTX_IN, test.GetTestHighResMessageText())
	mediaResizeWorkflow := workflow.MediaResize("bin/ffmpeg", &model.MediaFormatFilter{Width: "240"}, cloud.StorageClient, "media_low_res_resources")

	// This assertion insures the command can be executed
	assert.True(t, mediaResizeWorkflow.IsExecutable(chainCtx))
	mediaResizeWorkflow.Execute(chainCtx)

	for _, err := range chainCtx.GetErrors() {
		fmt.Println(err.Error())
	}

	assert.False(t, chainCtx.HasErrors())
}
