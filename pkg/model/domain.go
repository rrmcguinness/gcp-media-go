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

// Actor is used to represent the public details of an actor or actress.
type Actor struct {
	Name         string    `json:"name"`
	DateOfBirth  time.Time `json:"date_of_birth"` // Or time.Time if you need more precise handling
	DateOfDeath  time.Time `json:"date_of_death"`
	PlaceOfBirth string    `json:"place_of_birth"`
	Biography    string    `json:"biography"`
	Aliases      []string  `json:"aliases"`
	Awards       []string  `json:"awards"`
	Nominations  []string  `json:"nominations"`
	KnownFor     []string  `json:"known_for"` // Titles of notable movies/shows
	ImageURL     string    `json:"image_url"`
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

// Scene is a representation of a time span and it's sequence in a movie
// giving granular detail for the agent objects to interrogate
type Scene struct {
	SequenceNumber int    `json:"seq" bigquery:"sequence"`
	Start          string `json:"start" bigquery:"start"`
	End            string `json:"end" bigquery:"end"`
	Script         string `json:"script" bigquery:"script"`
}

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

// GetContextMovieIdentityName used for referring to movie identity consistently
func GetContextMovieIdentityName() string {
	return "__MOVIE_IDENTITY__"
}

// MovieIdentity is used to identity a movie file, not by title or file name, but by fingerprint.
type MovieIdentity struct {
	Id         string `json:"id" bigquery:"id"`
	FileSHA256 string `json:"sha256" bigquery:"sha256"`
}

// SceneEmbeddings captures the embeddings related to a specific scene good for additional research searches.
type SceneEmbeddings struct {
	SequenceNumber int
	Embeddings     []float64 `json:"embeddings" bigquery:"embeddings"`
}

// MovieEmbedding captures the summary embedding of a movie, good for general searches.
type MovieEmbedding struct {
	Id              string            `json:"id" bigquery:"id"`
	Model           string            `json:"model_name" bigquery:"model_name"`
	Embeddings      []float64         `json:"embeddings" bigquery:"embeddings"`
	SceneEmbeddings []SceneEmbeddings `json:"scene_embeddings" bigquery:"scene_embeddings"`
}

func NewMovieEmbedding() *MovieEmbedding {
	uuid, _ := uuid.NewV7()

	return &MovieEmbedding{
		Id:              uuid.String(),
		Embeddings:      make([]float64, 0),
		SceneEmbeddings: make([]SceneEmbeddings, 0),
	}
}

// GetContextMovieName used for referring to movie consistently
func GetContextMovieName() string {
	return "__MOVIE__"
}

// Movie capture the highest level of metadata about a movie.
type Movie struct {
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

func NewMovie() *Movie {
	uuid, _ := uuid.NewV7()

	return &Movie{
		Id:         uuid.String(),
		CreateDate: time.Now(),
		Cast:       make([]*CastMember, 0),
		Scenes:     make([]*Scene, 0),
	}
}
