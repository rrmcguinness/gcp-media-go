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
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"sort"
	"strings"
	"time"
)

const (
	DefaultMovieTimeFormat = "15:04:05"
)

type MediaAssembly struct {
	cor.BaseCommand
	summaryParam     string
	sceneParam       string
	mediaObjectParam string
}

// NewMediaAssembly default constructor for MediaAssembly
func NewMediaAssembly(name string, summaryParam string, sceneParam string, mediaObjectParam string) *MediaAssembly {
	return &MediaAssembly{
		BaseCommand:      *cor.NewBaseCommand(name),
		summaryParam:     summaryParam,
		sceneParam:       sceneParam,
		mediaObjectParam: mediaObjectParam}
}

// IsExecutable overrides the default to verify the summary param and scene param are in the context
func (m *MediaAssembly) IsExecutable(context cor.Context) bool {
	return context != nil &&
		context.Get(m.summaryParam) != nil &&
		context.Get(m.sceneParam) != nil
}

func (m *MediaAssembly) Execute(context cor.Context) {
	summary := context.Get(m.summaryParam).(*model.MediaSummary)
	jsonScenes := context.Get(m.sceneParam).([]string)
	sceneValues := fmt.Sprintf("[ %s ]", strings.Join(jsonScenes, ","))

	scenes := make([]*model.Scene, 0)
	sceneErr := json.Unmarshal([]byte(sceneValues), &scenes)
	if sceneErr != nil {
		m.GetErrorCounter().Add(context.GetContext(), 1)
		context.AddError(m.GetName(), sceneErr)
		return
	}

	// Sort the scenes and sequence them
	sort.Slice(scenes, func(i, j int) bool {
		t, _ := time.Parse(DefaultMovieTimeFormat, scenes[i].Start)
		tt, _ := time.Parse(DefaultMovieTimeFormat, scenes[j].Start)
		return t.Before(tt)
	})
	for i, scene := range scenes {
		scene.SequenceNumber = i
	}

	// Call the constructor to ensure the UUID is generated
	// TODO - Base the
	media := model.NewMedia(summary.Title)
	media.Title = summary.Title
	media.Category = summary.Category
	media.Summary = summary.Summary
	media.MediaUrl = summary.MediaUrl
	media.LengthInSeconds = summary.LengthInSeconds
	media.Director = summary.Director
	media.ReleaseYear = summary.ReleaseYear
	media.Genre = summary.Genre
	media.Rating = summary.Rating
	media.Cast = append(media.Cast, summary.Cast...)
	media.Scenes = append(media.Scenes, scenes...)

	m.GetSuccessCounter().Add(context.GetContext(), 1)

	context.Add(m.mediaObjectParam, media)
	context.Add(cor.CtxOut, media)
}
