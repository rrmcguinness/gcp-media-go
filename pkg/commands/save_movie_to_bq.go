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
	go_ctx "context"
	"log"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type SaveMovieToBQ struct {
	cor.BaseCommand
	BigQueryClient       *bigquery.Client
	MovieObjectParamName string
	DataSetName          string
	TableName            string
}

func (s *SaveMovieToBQ) IsExecutable(context cor.Context) bool {
	return context != nil && context.Get(s.MovieObjectParamName) != nil
}

func (s *SaveMovieToBQ) Execute(context cor.Context) {
	movie := context.Get(s.MovieObjectParamName).(*model.Movie)
	i := s.BigQueryClient.Dataset(s.DataSetName).Table(s.TableName).Inserter()
	if err := i.Put(go_ctx.Background(), movie); err != nil {
		log.Printf("Failed to write movie: %s with error %v\n", movie.Title, err)
	}
	context.Add(cor.CTX_OUT, movie)
}
