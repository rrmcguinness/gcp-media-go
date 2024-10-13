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
	"io"
	"os"
	"os/exec"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type FFMpegStreamingCommand struct {
	model.Command
	client              *storage.Client
	executableCommand   string
	executableArguments string
	destinationBucket   string
}

func (c *FFMpegStreamingCommand) GetExecutableCommandString() string {
	return c.executableCommand + " " + c.executableArguments
}

func NewFFMpegStreamingCommand(
	ctx context.Context,
	executableCommand string,
	destinationBucket string,
	format string,
	width string,
) *FFMpegStreamingCommand {

	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	scale := fmt.Sprintf("scale=w=%s:h=trunc(ow/a/2)*2", width)

	return &FFMpegStreamingCommand{
		client:              client,
		executableCommand:   executableCommand,
		executableArguments: "-analyzeduration 0 -probesize 5000000 -y -hide_banner -i %s -filter:v " + scale + " -f mp4 %s",
		destinationBucket:   destinationBucket,
	}
}

func (c *FFMpegStreamingCommand) Close() {
	c.client.Close()
}

func (c *FFMpegStreamingCommand) IsExecutable(chCtx model.ChainContext) bool {
	return chCtx.Get(model.CTX_MESSAGE).(*model.TriggerMediaWrite).Kind == model.EVENT_STORAGE_BUCKET_WRITE
}

func (c *FFMpegStreamingCommand) Execute(chCtx model.ChainContext) {

	//#############################################################################
	// Create a processing context, and get the message from the context and create temp files
	//#############################################################################
	ctx := context.Background()
	msg := chCtx.Get(model.CTX_MESSAGE).(*model.TriggerMediaWrite)

	originalFile, err := os.CreateTemp("", "tmp-vid-in-")
	defer os.Remove(originalFile.Name())

	outputFile, err := os.CreateTemp("", "tmp-vid-out-")
	defer os.Remove(outputFile.Name())

	//#############################################################################
	// Download the original file to a temp file
	//#############################################################################
	readerBucket := c.client.Bucket(msg.Bucket)
	obj := readerBucket.Object(msg.Name)
	reader, err := obj.NewReader(ctx)
	if err != nil {
		chCtx.AddError(fmt.Errorf("error creating GCS reader: %w", err))
		return
	}
	defer reader.Close()
	io.Copy(originalFile, reader)

	//#############################################################################
	// Execute FFMpeg which will write to the temp file
	//#############################################################################
	args := fmt.Sprintf(c.executableArguments, originalFile.Name(), outputFile.Name())
	cmd := exec.Command(c.executableCommand, strings.Split(args, " ")...)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		chCtx.AddError(fmt.Errorf("error running ffmpeg: %w", err))
		return
	}

	//#############################################################################
	// Copy the contents of the output file to the bucket.
	//#############################################################################
	writerBucket := c.client.Bucket(c.destinationBucket) // Use the destination bucket from the command
	obj = writerBucket.Object(msg.Name)
	writer := obj.NewWriter(ctx)
	defer writer.Close()
	io.Copy(writer, outputFile)
}
