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
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
)

type GCSFileUpload struct {
	cor.BaseCommand
	client *storage.Client
	bucket string
}

func NewGCSFileUpload(name string, client *storage.Client, bucket string) *GCSFileUpload {
	return &GCSFileUpload{BaseCommand: *cor.NewBaseCommand(name), client: client, bucket: bucket}
}

func (c *GCSFileUpload) Execute(context cor.Context) {
	path := context.Get(c.GetInputParam()).(string)
	name := filepath.Base(path)

	original := context.Get(cloud.GetGCSObjectName()).(*cloud.GCSObject)

	dat, err := os.Open(path)
	if err != nil {
		context.AddError(c.GetName(), err)
		return
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Printf("failed to remove file from OS: %v\n", err)
		}
	}(path)

	writerBucket := c.client.Bucket(c.bucket)
	if original != nil {
		obj := writerBucket.Object(original.Name)
		writer := obj.NewWriter(context.GetContext())
		defer func(writer *storage.Writer) {
			err := writer.Close()
			if err != nil {
				log.Printf("failed to close writer: %v\n", err)
			}
		}(writer)
		written, err := io.Copy(writer, dat)
		if err != nil {
			log.Printf("failed to close writer or partial write: %d total bytes, %v\n", written, err)
			context.AddError(c.GetName(), err)
			return
		}
	} else {
		obj := writerBucket.Object(name)
		writer := obj.NewWriter(context.GetContext())
		defer func(writer *storage.Writer) {
			err := writer.Close()
			if err != nil {
				log.Printf("failed to close writer: %v\n", err)
			}
		}(writer)
		written, err := io.Copy(writer, dat)
		if err != nil {
			log.Printf("failed to close writer or partial write: %d total bytes, %v\n", written, err)
			context.AddError(c.GetName(), err)
			return
		}
	}
}
