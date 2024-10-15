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

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/google/generative-ai-go/genai"
)

type VideoCleanupCommand struct {
	cor.BaseCommand
	GenaiClient genai.Client
}

func (v *VideoCleanupCommand) IsExecutable(context cor.Context) bool {
	return context != nil && context.Get("__Video_Upload__") != nil && context.Get("__Video_Upload__").(*genai.File) != nil
}

func (v *VideoCleanupCommand) Execute(context cor.Context) {
	ctx := go_ctx.Background()
	fil := context.Get("__Video_Upload__").(*genai.File)
	v.GenaiClient.DeleteFile(ctx, fil.Name)
}
