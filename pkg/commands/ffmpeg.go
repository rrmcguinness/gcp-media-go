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

package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
)

const DEFAULT_FFMPEG_ARGS = "-analyzeduration 0 -probesize 5000000 -y -hide_banner -i %s -filter:v scale=w=%s:h=trunc(ow/a/2)*2 -f mp4 %s"

// FFMpegCommand is a simple command used for
// downloading a media file embedded in the message, resizing it
// and uploading the resized version to the destination bucket.
// The scale uses a dynamic scale to keep the aspect ratio of the original.
type FFMpegCommand struct {
	cor.BaseCommand
	ExecutableCommand string
	TargetWidth       string
}

// IsExecutable determines if the command should execute
func (c *FFMpegCommand) IsExecutable(context cor.Context) bool {
	return context.Get(c.GetInputParam()) != nil
}

// Execute executes the business logic of the command
func (c *FFMpegCommand) Execute(context cor.Context) {
	inputFileName := context.Get(c.GetInputParam()).(string)
	file, err := os.Open(inputFileName)
	if err != nil {
		context.AddError(err)
		return
	}
	tempFile, err := os.CreateTemp("", "ffmpeg-output-")

	args := fmt.Sprintf(DEFAULT_FFMPEG_ARGS, file.Name(), c.TargetWidth, tempFile.Name())
	cmd := exec.Command(c.ExecutableCommand, strings.Split(args, " ")...)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		context.AddError(fmt.Errorf("error running ffmpeg: %w", err))
		return
	}
	context.AddTempFile(tempFile.Name())
	context.Add(cor.CTX_OUT, tempFile.Name())
}
