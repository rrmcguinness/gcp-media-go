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

package workflow

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
)

func GenerateEmbeddings(
	genaiEmbedding *genai.EmbeddingModel,
	bigqueryClient *bigquery.Client,
	datasetName string,
	mediaTable string,
	embeddingTable string,
	modelName string) (err error) {

	ctx := context.Background()

	fqMediaTableName := strings.Replace(bigqueryClient.Dataset(datasetName).Table(mediaTable).FullyQualifiedName(), ":", ".", -1)
	fqEmbeddingTable := strings.Replace(bigqueryClient.Dataset(datasetName).Table(embeddingTable).FullyQualifiedName(), ":", ".", -1)

	queryString := fmt.Sprintf("SELECT * FROM `%s` WHERE ID NOT IN (SELECT MEDIA_ID FROM `%s`)", fqMediaTableName, fqEmbeddingTable)

	q := bigqueryClient.Query(queryString)

	it, err := q.Read(ctx)
	if err != nil {
		return err
	}

	for {
		var value model.Media
		err := it.Next(&value)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		toInsert := make([]*model.SceneEmbedding, 0)

		for _, scene := range value.Scenes {
			in := model.NewSceneEmbedding(value.Id, scene.SequenceNumber, modelName)
			resp, err := genaiEmbedding.EmbedContent(ctx, genai.Text(scene.Script))
			if err != nil {
				return err
			}
			for _, f := range resp.Embedding.Values {
				in.Embeddings = append(in.Embeddings, float64(f))
			}
			toInsert = append(toInsert, in)
		}

		inserter := bigqueryClient.Dataset(datasetName).Table(embeddingTable).Inserter()
		if err := inserter.Put(ctx, toInsert); err != nil {
			return err
		}
	}
	return nil
}
