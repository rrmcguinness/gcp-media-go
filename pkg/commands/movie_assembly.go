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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type MovieAssembly struct {
	cor.BaseCommand
	SummaryParameterName     string
	SceneParameterName       string
	MovieObjectParameterName string
}

func (m *MovieAssembly) IsExecutable(context cor.Context) bool {
	return context != nil &&
		context.Get(m.SummaryParameterName) != nil &&
		context.Get(m.SceneParameterName) != nil
}

func (m *MovieAssembly) Execute(context cor.Context) {
	jsonSummary := context.Get(m.SummaryParameterName).(string)
	jsonScenes := context.Get(m.SceneParameterName).([]string)
	sceneValues := fmt.Sprintf("[ %s ]", strings.Join(jsonScenes, ","))

	summary := model.MovieSummary{}
	scenes := make([]*model.Scene, 0)
	_ = json.Unmarshal([]byte(jsonSummary), &summary)
	_ = json.Unmarshal([]byte(sceneValues), &scenes)

	movie := model.NewMovie()
	movie.Title = summary.Title
	movie.Summary = summary.Summary
	movie.Director = summary.Director
	movie.ReleaseYear = summary.ReleaseYear
	movie.Genre = summary.Genre
	movie.Rating = summary.Rating
	movie.Cast = append(movie.Cast, summary.Cast...)
	movie.Scenes = append(movie.Scenes, scenes...)

	context.Add(m.MovieObjectParameterName, movie)
	context.Add(cor.CTX_OUT, movie)
}
