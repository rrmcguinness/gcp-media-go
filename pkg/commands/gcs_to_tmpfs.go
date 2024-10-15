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
	"context"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type GCSToTempFileCommand struct {
	model.BaseCommand
	Client *storage.Client
}

func NewGCSToTempFileCommand(client *storage.Client) *GCSToTempFileCommand {
	if client == nil {
		panic("Client must not be nil")
	}
	return &GCSToTempFileCommand{
		Client: client,
	}
}

func (c *GCSToTempFileCommand) IsExecutable(chCtx model.ChainContext) bool {
	if chCtx != nil && chCtx.Get(c.GetInputParam()) != nil {
		return true
	}
	return false
}

func (c *GCSToTempFileCommand) Execute(chCtx model.ChainContext) {
	ctx := context.Background()

	msg := chCtx.Get(c.GetInputParam()).(*model.GCSObject)

	readerBucket := c.Client.Bucket(msg.Bucket)
	obj := readerBucket.Object(msg.Name)
	reader, err := obj.NewReader(ctx)
	if err != nil {
		chCtx.AddError(err)
	}
	defer reader.Close()
	tempFile, err := os.CreateTemp("", "gcs-to-tmp-fs")
	io.Copy(tempFile, reader)
	// Add to the temp files to clean up after execution in a chain
	chCtx.AddTempFile(tempFile.Name())
	// Make the variable available if needed
	chCtx.Add(c.GetOutputParam(), tempFile.Name())
}
