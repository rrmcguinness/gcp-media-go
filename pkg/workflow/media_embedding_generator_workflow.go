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
	go_ctx "context"
	"fmt"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
)

type MediaEmbeddingGeneratorWorkflow struct {
	cor.BaseCommand
	genaiEmbedding         *genai.EmbeddingModel
	bigqueryClient         *bigquery.Client
	dataset                string
	mediaTable             string
	embeddingTable         string
	findEligibleMediaQuery string
}

func (m *MediaEmbeddingGeneratorWorkflow) StartTimer() {
	tracer := otel.Tracer("embedding-batch")
	ticker := time.NewTicker(60 * time.Second)
	closeTicker := make(chan struct{})

	// Create a timer to run embedding checks every 60 seconds
	go func(m *MediaEmbeddingGeneratorWorkflow) {
		for {
			select {
			case <-ticker.C:
				traceCtx, span := tracer.Start(go_ctx.Background(), "media-embeddings")
				chainCtx := cor.NewBaseContext()
				chainCtx.SetContext(traceCtx)
				m.Execute(chainCtx)
				if chainCtx.HasErrors() {
					span.SetStatus(codes.Error, "failed to execute embedding chain")
				} else {
					span.SetStatus(codes.Ok, "executed embeddings")
				}
				span.End()
			case <-closeTicker:
				ticker.Stop()
				return
			}
		}
	}(m)
}

func NewMediaEmbeddingGeneratorWorkflow(config *cloud.Config, serviceClients *cloud.ServiceClients) *MediaEmbeddingGeneratorWorkflow {

	fqMediaTableName := strings.Replace(serviceClients.BiqQueryClient.Dataset(config.BigQueryDataSource.DatasetName).Table(config.BigQueryDataSource.MediaTable).FullyQualifiedName(), ":", ".", -1)
	fqEmbeddingTable := strings.Replace(serviceClients.BiqQueryClient.Dataset(config.BigQueryDataSource.DatasetName).Table(config.BigQueryDataSource.EmbeddingTable).FullyQualifiedName(), ":", ".", -1)
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE ID NOT IN (SELECT MEDIA_ID FROM `%s`)", fqMediaTableName, fqEmbeddingTable)

	return &MediaEmbeddingGeneratorWorkflow{
		BaseCommand:            *cor.NewBaseCommand("media-embedding-generator"),
		genaiEmbedding:         serviceClients.EmbeddingModels["multi-lingual"],
		bigqueryClient:         serviceClients.BiqQueryClient,
		dataset:                config.BigQueryDataSource.DatasetName,
		mediaTable:             config.BigQueryDataSource.MediaTable,
		embeddingTable:         config.BigQueryDataSource.EmbeddingTable,
		findEligibleMediaQuery: query,
	}
}

func (m *MediaEmbeddingGeneratorWorkflow) IsExecutable(_ cor.Context) bool {
	return true
}

func (m *MediaEmbeddingGeneratorWorkflow) Execute(context cor.Context) {
	q := m.bigqueryClient.Query(m.findEligibleMediaQuery)
	it, err := q.Read(context.GetContext())
	if err != nil {
		context.AddError(err)
		return
	}

	for {
		var value model.Media
		fmt.Println(value.Title)
		err := it.Next(&value)
		if err == iterator.Done {
			break
		}
		if err != nil {
			context.AddError(err)
			return
		}

		toInsert := make([]*model.SceneEmbedding, 0)

		for _, scene := range value.Scenes {
			in := model.NewSceneEmbedding(value.Id, scene.SequenceNumber, m.genaiEmbedding.Name())
			resp, err := m.genaiEmbedding.EmbedContent(context.GetContext(), genai.Text(scene.Script))
			if err != nil {
				context.AddError(err)
				return
			}
			for _, f := range resp.Embedding.Values {
				in.Embeddings = append(in.Embeddings, float64(f))
			}
			toInsert = append(toInsert, in)
		}

		inserter := m.bigqueryClient.Dataset(m.dataset).Table(m.embeddingTable).Inserter()
		if err := inserter.Put(context.GetContext(), toInsert); err != nil {
			context.AddError(err)
			return
		}
	}
}
