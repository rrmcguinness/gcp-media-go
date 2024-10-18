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
	go_ctx "context"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type GCSToTempFile struct {
	cor.BaseCommand
	Client         *storage.Client
	TempFilePrefix string
}

func NewGCSToTempFile(client *storage.Client, tempFilePrefix string) *GCSToTempFile {
	if client == nil {
		panic("Client must not be nil")
	}
	return &GCSToTempFile{
		Client:         client,
		TempFilePrefix: tempFilePrefix,
	}
}

func (c *GCSToTempFile) Execute(context cor.Context) {
	ctx := go_ctx.Background()

	msg := context.Get(c.GetInputParam()).(*model.GCSObject)

	readerBucket := c.Client.Bucket(msg.Bucket)
	obj := readerBucket.Object(msg.Name)
	reader, err := obj.NewReader(ctx)
	if err != nil {
		context.AddError(err)
	}
	defer reader.Close()
	tempFile, err := os.CreateTemp("", c.TempFilePrefix)
	io.Copy(tempFile, reader)
	// Add to the temp files to clean up after execution in a chain
	context.AddTempFile(tempFile.Name())
	// Make the variable available if needed
	context.Add(c.GetOutputParam(), tempFile.Name())
}
