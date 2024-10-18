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

// GetExampleScene is used to provide an example to the generative contexts.
func GetExampleScene() *Scene {
	out := &Scene{SequenceNumber: 1, Start: "00:00:00", End: "00:01:00", Script: `
	INT. BATTLEFIELD - DAY

A fierce battle is raging. Soldiers are fighting and dying all around.

VOICEOVER (V.O.)
I aim to misbehave.

We see a young woman, RIVER TAM (16), running through the battlefield. She is terrified and covered in blood.

RIVER (V.O.)
They were right. They were always right.

River stumbles and falls. She looks up to see a man standing over her. He is SIMON TAM (26), her older brother.

SIMON
It's all right, River. I'm here.

Simon helps River to her feet. They run away together.`}
	return out
}

func GetExampleSummary() *MediaSummary {
	s := &MediaSummary{
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
