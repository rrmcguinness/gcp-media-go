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
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/google/generative-ai-go/genai"
)

type MediaCleanup struct {
	cor.BaseCommand
	client *genai.Client
}

func NewMediaCleanup(name string, client *genai.Client) *MediaCleanup {
	return &MediaCleanup{BaseCommand: *cor.NewBaseCommand(name), client: client}
}

func (v *MediaCleanup) IsExecutable(context cor.Context) bool {
	return context != nil && context.Get(GetVideoUploadFileParameterName()) != nil &&
		context.Get(GetVideoUploadFileParameterName()).(*genai.File) != nil
}

func (v *MediaCleanup) Execute(context cor.Context) {
	fil := context.Get(GetVideoUploadFileParameterName()).(*genai.File)
	err := v.client.DeleteFile(context.GetContext(), fil.Name)
	if err != nil {
		v.GetErrorCounter().Add(context.GetContext(), 1)
		context.AddError(v.GetName(), err)
		return
	}
	v.GetSuccessCounter().Add(context.GetContext(), 1)
}
