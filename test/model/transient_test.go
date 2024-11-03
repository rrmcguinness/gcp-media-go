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

package model_test

import (
	"encoding/json"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMediaSummaryFromJSON(t *testing.T) {
	// Test case with valid JSON
	validJSON := `{
		"title": "Test Title",
		"summary": "This is a test summary.",
		"director": "Test Director",
		"release_year": 2023,
		"genre": "Test Genre",
		"rating": "PG-13",
		"cast": [
			{"actor_name": "Actor 1", "character_name": "Character 1"},
			{"actor_name": "Actor 2", "character_name": "Character 2"}
		],
		"scene_time_stamps": [
			{"start": "00:00:00", "end": "00:01:00"},
			{"start": "00:02:00", "end": "00:03:00"}
		]
	}`

	summary, err := model.MediaSummaryFromJSON(validJSON)

	assert.NoError(t, err, "Expected no error for valid JSON")
	assert.Equal(t, "Test Title", summary.Title, "Title should match")
	assert.Equal(t, "This is a test summary.", summary.Summary, "Summary should match")
	assert.Equal(t, "Test Director", summary.Director, "Director should match")
	assert.Equal(t, 2023, summary.ReleaseYear, "ReleaseYear should match")
	assert.Equal(t, "Test Genre", summary.Genre, "Genre should match")
	assert.Equal(t, "PG-13", summary.Rating, "Rating should match")
	assert.Equal(t, 2, len(summary.Cast), "Cast should have 2 members")
	assert.Equal(t, "Actor 1", summary.Cast[0].ActorName, "Cast member 1 name should match")
	assert.Equal(t, 2, len(summary.SceneTimeStamps), "SceneTimeStamps should have 2 entries")
	assert.Equal(t, "00:00:00", summary.SceneTimeStamps[0].Start, "SceneTimeStamp 1 start should match")

	// Test case with invalid JSON
	invalidJSON := `{"title": "Test Title", "summary": "This is a test summary.`

	_, err = model.MediaSummaryFromJSON(invalidJSON)

	assert.Error(t, err, "Expected error for invalid JSON")
}

func TestTimeSpan_UnmarshalJSON(t *testing.T) {
	jsonString := `{"start": "00:00:00", "end": "00:01:00"}`

	timeSpan := &model.TimeSpan{}
	err := json.Unmarshal([]byte(jsonString), timeSpan)

	assert.NoError(t, err)
	assert.Equal(t, "00:00:00", timeSpan.Start)
	assert.Equal(t, "00:01:00", timeSpan.End)
}
