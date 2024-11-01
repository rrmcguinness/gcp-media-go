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
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type MediaPersistToBigQuery struct {
	cor.BaseCommand
	client     *bigquery.Client
	dataset    string
	table      string
	mediaParam string
}

func NewMediaPersistToBigQuery(name string, client *bigquery.Client, dataset string, table string, mediaParam string) *MediaPersistToBigQuery {
	return &MediaPersistToBigQuery{BaseCommand: *cor.NewBaseCommand(name), client: client, dataset: dataset, table: table, mediaParam: mediaParam}
}

func (s *MediaPersistToBigQuery) IsExecutable(context cor.Context) bool {
	executable := context != nil && context.Get(s.mediaParam) != nil
	log.Printf("can persists data: %v\n", executable)
	return executable
}

func (s *MediaPersistToBigQuery) Execute(context cor.Context) {
	log.Println("Persisting data")
	media := context.Get(s.mediaParam).(*model.Media)
	log.Printf("saving media to database: title %s id %s", media.Title, media.Id)
	i := s.client.Dataset(s.dataset).Table(s.table).Inserter()
	if err := i.Put(context.GetContext(), media); err != nil {
		log.Printf("failed to write media to database. title %s error %s", media.Title, err)
		fmt.Println("Error writing to database", err)
		context.AddError(err)
		return
	}
	context.Add(cor.CTX_OUT, media)
}
