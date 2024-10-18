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
	"time"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/google/generative-ai-go/genai"
)

type VideoUploadCommand struct {
	cor.BaseCommand
	GenaiClient      *genai.Client
	TimeoutInSeconds time.Duration
}

func GetVideoUploadFileParameterName() string {
	return "__VIDEO_UPLOAD_FILE__"
}

func (v *VideoUploadCommand) Execute(context cor.Context) {
	ctx, cancel := go_ctx.WithTimeout(go_ctx.Background(), v.TimeoutInSeconds)
	defer cancel()

	gcsFile := context.Get(model.GetGCSObjectName()).(*model.GCSObject)
	fileName := context.Get(v.GetInputParam()).(string)

	genFil, err := v.GenaiClient.UploadFileFromPath(ctx, fileName, &genai.UploadFileOptions{DisplayName: gcsFile.Name, MIMEType: gcsFile.MIMEType})
	if err != nil {
		context.AddError(err)
		return
	}

	// Videos need to be processed before you can use them.
	for genFil.State == genai.FileStateProcessing {
		time.Sleep(5 * time.Second)
		var err error
		if genFil, err = v.GenaiClient.GetFile(ctx, genFil.Name); err != nil {
			context.AddError(err)
		}
	}

	context.Add(GetVideoUploadFileParameterName(), genFil)
	context.Add(v.GetOutputParam(), genFil)
}
