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

import (
	"time"

	"github.com/google/uuid"
)

// GetContextMediaName used for referring to media consistently
func GetContextMediaName() string {
	return "__MEDIA__"
}

// Actor is used to represent the public details of an actor or actress.
type Actor struct {
	Id           string    `json:"id" bigquery:"id"`
	CreateDate   time.Time `json:"create_date" bigquery:"create_date"`
	Name         string    `json:"name" bigquery:"name"`
	DateOfBirth  time.Time `json:"dob" bigquery:"dob"` // Or time.Time if you need more precise handling
	DateOfDeath  time.Time `json:"dod" bigquery:"dod"`
	PlaceOfBirth string    `json:"pob" bigquery:"pob"`
	Biography    string    `json:"bio" bigquery:"bio"`
	Aliases      []string  `json:"aliases" bigquery:"aliases"`
	Awards       []string  `json:"awards" bigquery:"awards"`
	Nominations  []string  `json:"nominations" bigquery:"nominations"`
	ImageURL     string    `json:"ima_url" bigquery:"img_url"`
}

// Media capture the highest level of metadata about a media file.
type Media struct {
	Id          string        `json:"id" bigquery:"id"`
	CreateDate  time.Time     `json:"create_date" bigquery:"create_date"`
	Title       string        `json:"title" bigquery:"title"`
	Summary     string        `json:"summary" bigquery:"summary"`
	Director    string        `json:"director" bigquery:"director"`
	ReleaseYear int           `json:"release_year" bigquery:"release_year"`
	Genre       string        `json:"genre" bigquery:"genre"`
	Rating      string        `json:"rating" bigquery:"rating"`
	Cast        []*CastMember `json:"cast" bigquery:"cast"`
	Scenes      []*Scene      `json:"scenes" bigquery:"scenes"`
}

func NewMedia() *Media {
	uuid, _ := uuid.NewV7()

	return &Media{
		Id:         uuid.String(),
		CreateDate: time.Now(),
		Cast:       make([]*CastMember, 0),
		Scenes:     make([]*Scene, 0),
	}
}

// Scene is a representation of a time span and it's sequence in a media object
// giving granular detail for the agent objects to interrogate
type Scene struct {
	SequenceNumber int    `json:"sequence" bigquery:"sequence"`
	Start          string `json:"start" bigquery:"start"`
	End            string `json:"end" bigquery:"end"`
	Script         string `json:"script" bigquery:"script"`
}

// CastMember is a mapping object from a character to an actor
type CastMember struct {
	CharacterName string `json:"character_name" bigquery:"character_name"`
	ActorName     string `json:"actor_name" bigquery:"actor_name"`
}

// CastDialog is a mapping from a character to the spoken word in a scene
type CastDialog struct {
	CharacterName string `json:"character_name" bigquery:"character_name"`
	Dialog        string `json:"dialog" bigquery:"dialog"`
}

// SceneEmbedding captures the summary embedding of a media file, good for general searches.
type SceneEmbedding struct {
	Id             string    `json:"id" bigquery:"media_id"`
	SequenceNumber int       `json:"sequence_number" bigquery:"sequence_number"`
	ModelName      string    `json:"model_name" bigquery:"model_name"`
	Embeddings     []float64 `json:"embeddings" bigquery:"embeddings"`
}

func NewSceneEmbedding(
	mediaId string,
	sequenceNumber int,
	modelName string) *SceneEmbedding {

	return &SceneEmbedding{
		Id:             mediaId,
		SequenceNumber: sequenceNumber,
		ModelName:      modelName,
		Embeddings:     make([]float64, 0),
	}
}
