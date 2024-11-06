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
	"time"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/google/generative-ai-go/genai"
)

type MediaUpload struct {
	cor.BaseCommand
	client           *genai.Client
	timeoutInSeconds time.Duration
}

func NewMediaUpload(name string, genaiClient *genai.Client, timeoutInSeconds time.Duration) *MediaUpload {
	return &MediaUpload{BaseCommand: *cor.NewBaseCommand(name), client: genaiClient, timeoutInSeconds: timeoutInSeconds}
}

func GetVideoUploadFileParameterName() string {
	return "__VIDEO_UPLOAD_FILE__"
}

func (v *MediaUpload) Execute(context cor.Context) {
	gcsFile := context.Get(cloud.GetGCSObjectName()).(*cloud.GCSObject)
	fileName := context.Get(v.GetInputParam()).(string)

	genFil, err := v.client.UploadFileFromPath(context.GetContext(), fileName, &genai.UploadFileOptions{DisplayName: gcsFile.Name, MIMEType: gcsFile.MIMEType})
	if err != nil {
		v.GetErrorCounter().Add(context.GetContext(), 1)
		context.AddError(v.GetName(), err)
		return
	}

	// Videos need to be processed before you can use them.
	for genFil.State == genai.FileStateProcessing {
		time.Sleep(5 * time.Second)
		var err error
		if genFil, err = v.client.GetFile(context.GetContext(), genFil.Name); err != nil {
			v.GetErrorCounter().Add(context.GetContext(), 1)
			context.AddError(v.GetName(), err)
			return
		}
	}

	v.GetSuccessCounter().Add(context.GetContext(), 1)
	context.Add(GetVideoUploadFileParameterName(), genFil)
	context.Add(v.GetOutputParam(), genFil)
}
