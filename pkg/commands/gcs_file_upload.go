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
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type GCSFileUpload struct {
	cor.BaseCommand
	Client storage.Client
	Bucket string
}

func (c *GCSFileUpload) Execute(context cor.Context) {
	ctx := go_ctx.Background()
	path := context.Get(c.GetInputParam()).(string)
	name := filepath.Base(path)

	original := context.Get(model.GetGCSObjectName()).(*model.GCSObject)

	dat, err := os.Open(path)
	if err != nil {
		context.AddError(err)
		return
	}
	defer os.Remove(path)

	writerBucket := c.Client.Bucket(c.Bucket)
	if original != nil {
		obj := writerBucket.Object(original.Name)
		writer := obj.NewWriter(ctx)
		defer writer.Close()
		io.Copy(writer, dat)
	} else {
		obj := writerBucket.Object(name)
		writer := obj.NewWriter(ctx)
		defer writer.Close()
		io.Copy(writer, dat)
	}
}
