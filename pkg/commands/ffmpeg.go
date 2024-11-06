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

const (
	DefaultFfmpegArgs = "-analyzeduration 0 -probesize 5000000 -y -hide_banner -i %s -filter:v scale=w=%s:h=trunc(ow/a/2)*2 -f mp4 %s"
	TempFilePrefix    = "ffmpeg-output-"
	CommandSeparator  = " "
)

// FFMpegCommand is a simple command used for
// downloading a media file embedded in the message, resizing it
// and uploading the resized version to the destination bucket.
// The scale uses a dynamic scale to keep the aspect ratio of the original.
type FFMpegCommand struct {
	cor.BaseCommand
	commandPath string
	targetWidth string
}

func NewFFMpegCommand(name string, commandPath string, targetWidth string) *FFMpegCommand {
	return &FFMpegCommand{
		BaseCommand: *cor.NewBaseCommand(name),
		commandPath: commandPath,
		targetWidth: targetWidth}
}

// Execute executes the business logic of the command
func (c *FFMpegCommand) Execute(context cor.Context) {
	inputFileName := context.Get(c.GetInputParam()).(string)
	file, err := os.Open(inputFileName)
	if err != nil {
		c.GetErrorCounter().Add(context.GetContext(), 1)
		context.AddError(c.GetName(), err)
		return
	}
	tempFile, err := os.CreateTemp("", TempFilePrefix)

	args := fmt.Sprintf(DefaultFfmpegArgs, file.Name(), c.targetWidth, tempFile.Name())
	cmd := exec.Command(c.commandPath, strings.Split(args, CommandSeparator)...)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		c.GetErrorCounter().Add(context.GetContext(), 1)
		context.AddError(c.GetName(), fmt.Errorf("error running ffmpeg: %w", err))
		return
	}

	c.GetSuccessCounter().Add(context.GetContext(), 1)
	context.AddTempFile(tempFile.Name())
	context.Add(cor.CtxOut, tempFile.Name())
}
