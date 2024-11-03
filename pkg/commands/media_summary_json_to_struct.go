// Package commands Copyright 2024 Google, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package commands

import (
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type MediaSummaryJsonToStruct struct {
	cor.BaseCommand
}

func NewMediaSummaryJsonToStruct(name string, outputParamName string) *MediaSummaryJsonToStruct {
	out := MediaSummaryJsonToStruct{BaseCommand: *cor.NewBaseCommand(name)}
	out.OutputParamName = outputParamName
	return &out
}

func (s *MediaSummaryJsonToStruct) Execute(context cor.Context) {
	in := context.Get(s.GetInputParam()).(string)
	doc, err := model.MediaSummaryFromJSON(in)
	if err != nil {
		context.AddError(err)
		return
	}
	context.Add(s.GetOutputParam(), doc)
	context.Add(cor.CtxOut, doc)
}
