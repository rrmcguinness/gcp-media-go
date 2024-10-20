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
	"sort"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type MediaAssembly struct {
	cor.BaseCommand
	SummaryParameterName     string
	SceneParameterName       string
	MediaObjectParameterName string
}

func (m *MediaAssembly) IsExecutable(context cor.Context) bool {
	return context != nil &&
		context.Get(m.SummaryParameterName) != nil &&
		context.Get(m.SceneParameterName) != nil
}

func (m *MediaAssembly) Execute(context cor.Context) {
	jsonSummary := context.Get(m.SummaryParameterName).(string)
	jsonScenes := context.Get(m.SceneParameterName).([]string)
	sceneValues := fmt.Sprintf("[ %s ]", strings.Join(jsonScenes, ","))

	summary := model.MediaSummary{}
	scenes := make([]*model.Scene, 0)
	_ = json.Unmarshal([]byte(jsonSummary), &summary)
	_ = json.Unmarshal([]byte(sceneValues), &scenes)

	sort.Slice(scenes, func(i, j int) bool {
		t, _ := time.Parse("15:04:05", scenes[i].Start)
		tt, _ := time.Parse("15:04:05", scenes[j].Start)
		return t.Before(tt)
	})

	for i, scene := range scenes {
		scene.SequenceNumber = i
	}

	media := model.NewMedia()
	media.Title = summary.Title
	media.Summary = summary.Summary
	media.Director = summary.Director
	media.ReleaseYear = summary.ReleaseYear
	media.Genre = summary.Genre
	media.Rating = summary.Rating
	media.Cast = append(media.Cast, summary.Cast...)
	media.Scenes = append(media.Scenes, scenes...)

	context.Add(m.MediaObjectParameterName, media)
	context.Add(cor.CTX_OUT, media)
}
