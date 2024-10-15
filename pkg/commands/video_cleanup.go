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

	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/google/generative-ai-go/genai"
)

type VideoCleanupCommand struct {
	model.BaseCommand
	GenaiClient genai.Client
}

func (v *VideoCleanupCommand) IsExecutable(chCtx model.ChainContext) bool {
	return chCtx != nil && chCtx.Get("__Video_Upload__") != nil && chCtx.Get("__Video_Upload__").(*genai.File) != nil
}

func (v *VideoCleanupCommand) Execute(chCtx model.ChainContext) {
	ctx := context.Background()
	fil := chCtx.Get("__Video_Upload__").(*genai.File)
	v.GenaiClient.DeleteFile(ctx, fil.Name)
}
