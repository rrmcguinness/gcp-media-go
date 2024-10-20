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

package model

import "encoding/json"

// These objects are used in memory via workflows, but are not persisted to the dataset

func GetMediaSummaryName() string {
	return "__MEDIA_SUMMARY__"
}

func GetMediaFormatFilterName() string {
	return "__VIDEO_FORMAT__"
}

// MediaFormatFilter is a simple video format object expressing the intended output
// and the destination width
type MediaFormatFilter struct {
	Format string
	Width  string
}

type TimeSpan struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type MediaSummary struct {
	Title           string        `json:"title"`
	Summary         string        `json:"summary"`
	Director        string        `json:"director"`
	ReleaseYear     int           `json:"release_year"`
	Genre           string        `json:"genre"`
	Rating          string        `json:"rating"`
	Cast            []*CastMember `json:"cast"`
	SceneTimeStamps []*TimeSpan   `json:"scene_time_stamps"`
}

func MediaSummaryFromJSON(value string) (summary *MediaSummary, err error) {
	summary = &MediaSummary{}
	err = json.Unmarshal([]byte(value), &summary)
	return summary, err
}

type SceneMatchResult struct {
	MediaId        string `json:"media_id" bigquery:"media_id"`
	SequenceNumber int    `json:"sequence_number" bigquery:"sequence_number"`
}
