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
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/workflow"
	"github.com/GoogleCloudPlatform/solutions/media/test"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/codes"
)

func TestFFMpegCommand(t *testing.T) {
	traceContext, span := tracer.Start(ctx, "media-resize-test")
	defer span.End()

	// Create the context
	chainCtx := cor.NewBaseContext()
	chainCtx.SetContext(traceContext)
	chainCtx.Add(cor.CTX_IN, test.GetTestHighResMessageText())
	mediaResizeWorkflow := workflow.MediaResize(
		"bin/ffmpeg",
		&model.MediaFormatFilter{Width: "240"},
		cloudClients.StorageClient,
		"media_low_res_resources")

	// This assertion insures the command can be executed
	assert.True(t, mediaResizeWorkflow.IsExecutable(chainCtx))
	mediaResizeWorkflow.Execute(chainCtx)

	for _, err := range chainCtx.GetErrors() {
		fmt.Println(err.Error())
	}

	if chainCtx.HasErrors() {
		span.SetStatus(codes.Error, "failed - media-resize-test")
	}

	assert.False(t, chainCtx.HasErrors())
	span.SetStatus(codes.Ok, "passed - media-resize-test")
}
