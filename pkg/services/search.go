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
	"strconv"
	"strings"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
)

type SearchService struct {
	BigqueryClient *bigquery.Client
	EmbeddingModel *genai.EmbeddingModel
	DatasetName    string
	MediaTable     string
	EmbeddingTable string
}

func (s *SearchService) FindScenes(ctx context.Context, query string, maxResults int) (out []*model.SceneMatchResult, err error) {
	out = make([]*model.SceneMatchResult, 0)
	searchEmbeddings, _ := s.EmbeddingModel.EmbedContent(ctx, genai.Text(query))

	fqEmbeddingTable := strings.Replace(s.BigqueryClient.Dataset(s.DatasetName).Table(s.EmbeddingTable).FullyQualifiedName(), ":", ".", -1)

	var stringArray []string
	for _, f := range searchEmbeddings.Embedding.Values {
		stringArray = append(stringArray, strconv.FormatFloat(float64(f), 'f', -1, 64))
	}

	queryText := fmt.Sprintf(QrySequenceKnn, fqEmbeddingTable, strings.Join(stringArray, ","), maxResults)

	q := s.BigqueryClient.Query(queryText)
	itr, err := q.Read(ctx)
	if err != nil {
		return out, err
	}

	for {
		var r = &model.SceneMatchResult{}
		err := itr.Next(r)
		if err == iterator.Done {
			break
		}
		out = append(out, r)
	}
	return out, err
}
