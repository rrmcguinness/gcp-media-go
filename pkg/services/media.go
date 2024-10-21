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

package services

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type MediaService struct {
	BigqueryClient *bigquery.Client
	DatasetName    string
	MediaTable     string
}

// GetFQN returns the fully qualified BQ Table Name
func (s *MediaService) GetFQN() string {
	return strings.Replace(s.BigqueryClient.Dataset(s.DatasetName).Table(s.MediaTable).FullyQualifiedName(), ":", ".", -1)
}

// Get returns a media object by id, or an error if it doesn't exist
func (s *MediaService) Get(id string) (media *model.Media, err error) {
	ctx := context.Background()

	queryText := fmt.Sprintf(QRY_FIND_MEDIA_BY_ID, s.GetFQN(), id)
	q := s.BigqueryClient.Query(queryText)
	itr, err := q.Read(ctx)
	if err != nil {
		return media, err
	}
	// Since this should only return a single result
	media = &model.Media{}
	err = itr.Next(media)
	return media, err
}

// GetScene returns a scene in a specified media type by its sequence number
func (s *MediaService) GetScene(id string, sceneSequence int) (scene *model.Scene, err error) {
	fqMediaTableName := strings.Replace(s.BigqueryClient.Dataset(s.DatasetName).Table(s.MediaTable).FullyQualifiedName(), ":", ".", -1)
	ctx := context.Background()
	queryText := fmt.Sprintf(QRY_GET_SCENE, fqMediaTableName, id, sceneSequence)
	q := s.BigqueryClient.Query(queryText)
	itr, err := q.Read(ctx)
	if err != nil {
		return scene, err
	}
	scene = &model.Scene{}
	// Since this should only return a single result
	err = itr.Next(scene)
	return scene, err
}
