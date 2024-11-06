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
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
)

type GCSToTempFile struct {
	cor.BaseCommand
	client         *storage.Client
	tempFilePrefix string
}

func NewGCSToTempFile(name string, client *storage.Client, tempFilePrefix string) *GCSToTempFile {
	return &GCSToTempFile{
		BaseCommand:    *cor.NewBaseCommand(name),
		client:         client,
		tempFilePrefix: tempFilePrefix,
	}
}

func (c *GCSToTempFile) Execute(context cor.Context) {
	msg := context.Get(c.GetInputParam()).(*cloud.GCSObject)

	readerBucket := c.client.Bucket(msg.Bucket)
	obj := readerBucket.Object(msg.Name)
	reader, err := obj.NewReader(context.GetContext())
	if err != nil {
		c.GetErrorCounter().Add(context.GetContext(), 1)
		context.AddError(c.GetName(), err)
		return
	}
	defer func(reader *storage.Reader, context cor.Context) {
		err := reader.Close()
		if err != nil {
			c.GetErrorCounter().Add(context.GetContext(), 1)
			log.Printf("failed to close reader: %v\n", err)
		}
	}(reader, context)

	tempFile, err := os.CreateTemp("", c.tempFilePrefix)
	written, err := io.Copy(tempFile, reader)
	if err != nil {
		c.GetErrorCounter().Add(context.GetContext(), 1)
		log.Printf("failed to copy io, %d written: %v\n", written, err)
		context.AddError(c.GetName(), err)
		return
	}
	c.GetSuccessCounter().Add(context.GetContext(), 1)
	// Add to the temp files to clean up after execution in a chain
	context.AddTempFile(tempFile.Name())
	// Make the variable available if needed
	context.Add(c.GetOutputParam(), tempFile.Name())
}
