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
	"io"
	"os/exec"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type FFMpegStreamingCommand struct {
	model.Command
	executableCommand   string
	executableArguments []string
	inputSource         string
	destinationBucket   string
}

func NewFFMpegStreamingCommand(
	executableCommand string,
	executableArguments []string,
	outputDestination string,
) *FFMpegStreamingCommand {

	var args []string
	args = append(args, "-i", "pipe:0")
	args = append(args, executableArguments...)
	args = append(args, "pipe:1")

	return &FFMpegStreamingCommand{
		executableCommand:   executableCommand,
		executableArguments: args,
		destinationBucket:   outputDestination,
	}
}

func (c *FFMpegStreamingCommand) IsExecutable(chCtx model.ChainContext) bool {
	return chCtx.Get(model.CTX_MESSAGE).(model.TriggerMediaWrite).Kind == model.EVENT_STORAGE_BUCKET_WRITE
}

func (c *FFMpegStreamingCommand) Execute(chCtx model.ChainContext) {
	model := chCtx.Get(model.CTX_MESSAGE).(model.TriggerMediaWrite)

	cmd := exec.Command(c.executableCommand, c.executableArguments...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		chCtx.AddError(err)
		return
	}
	defer storageClient.Close()

	// Create the source Reader
	sourceBucket := storageClient.Bucket(model.Bucket)
	sourceObject := sourceBucket.Object(model.Name)

	sourceReader, err := sourceObject.NewReader(ctx)
	if err != nil {
		chCtx.AddError(err)
		return
	}
	defer sourceReader.Close()

	// Create the destination writer
	destBucket := storageClient.Bucket(c.destinationBucket)
	destObject := destBucket.Object(model.Name)

	// Create a writer for the destination file.
	destWriter := destObject.NewWriter(ctx)
	defer destWriter.Close()

	// Start ffmpeg.
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// Copy the source file to ffmpeg's standard input.
	go func() {
		defer stdin.Close()
		if _, err := io.Copy(stdin, sourceReader); err != nil {
			panic(err)
		}
	}()

	// Copy ffmpeg's standard output to the destination file.
	if _, err := io.Copy(destWriter, stdout); err != nil {
		panic(err)
	}

	// Wait for ffmpeg to finish.
	if err := cmd.Wait(); err != nil {
		panic(err)
	}

}
