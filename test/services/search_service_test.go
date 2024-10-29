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

package services_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/services"
	"github.com/GoogleCloudPlatform/solutions/media/test"
	"github.com/zeebo/assert"
)

func TestSearchService(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// This deferral will automatically close the client that was build from
	// the same context
	defer cancel()

	// Get the config file
	config := test.GetConfig()

	cloudClients, err := cloud.NewCloudServiceClients(ctx, config)
	test.HandleErr(err, t)
	defer cloudClients.Close()

	genModel := cloudClients.AgentModels["creative-flash"]
	assert.NotNil(t, genModel)

	embeddingModel := cloudClients.EmbeddingModels["multi-lingual"]

	searchService := &services.SearchService{
		BigqueryClient: cloudClients.BiqQueryClient,
		EmbeddingModel: embeddingModel,
		DatasetName:    "media_ds",
		MediaTable:     "media",
		EmbeddingTable: "scene_embeddings",
	}

	out, err := searchService.FindScenes(ctx, "Scenes that Woody Harrelson", 5)

	if err != nil {
		t.Error(err)
	}

	assert.Nil(t, err)

	for _, o := range out {
		fmt.Printf("%s - %d\n", o.MediaId, o.SequenceNumber)
	}
}
