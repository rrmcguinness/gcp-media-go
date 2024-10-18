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

// These objects are used during the generative model phase and represent
// parts of the whole

func GetMovieSummaryName() string {
	return "__MOVIE_SUMMARY__"
}

type TimeSpan struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type MovieSummary struct {
	Title           string        `json:"title"`
	Summary         string        `json:"summary"`
	Director        string        `json:"director"`
	ReleaseYear     int           `json:"release_year"`
	Genre           string        `json:"genre"`
	Rating          string        `json:"rating"`
	Cast            []*CastMember `json:"cast"`
	SceneTimeStamps []*TimeSpan   `json:"scene_time_stamps"`
}

func GetExampleSummary() *MovieSummary {
	s := &MovieSummary{
		Title:           "Serenity",
		Summary:         "The crew of the ship Serenity try to evade an assassin sent to recapture telepath River.",
		Director:        "Joss Whedon",
		ReleaseYear:     2005,
		Genre:           "Science Fiction",
		Rating:          "PG-13",
		SceneTimeStamps: make([]*TimeSpan, 0),
		Cast:            make([]*CastMember, 0),
	}
	s.SceneTimeStamps = append(s.SceneTimeStamps, &TimeSpan{Start: "00:00:00", End: "00:00:05"}, &TimeSpan{Start: "00:00:06", End: "00:00:10"})
	s.Cast = append(s.Cast, &CastMember{CharacterName: "Malcolm Reynolds", ActorName: "Nathan Fillion"})
	return s
}
