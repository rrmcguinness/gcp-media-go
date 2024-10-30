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

package main

import (
	"context"
	"os"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/services"
)

type StateManager struct {
	config        *cloud.CloudConfig
	cloud         *cloud.CloudServiceClients
	searchService *services.SearchService
	mediaService  *services.MediaService
}

var state = &StateManager{}

func SetupOS() {
	os.Setenv(cloud.ENV_CONFIG_FILE_PREFIX, "configs")
	os.Setenv(cloud.ENV_CONFIG_RUNTIME, "test")
}

func GetConfig() *cloud.CloudConfig {
	if state.config == nil {
		SetupOS()
		// Create a default cloud config
		config := cloud.NewCloudConfig()
		// Load it from the TOML files
		cloud.LoadConfig(&config)
		state.config = config
	}
	return state.config
}

func InitState(ctx context.Context) {
	// Get the config file
	config := GetConfig()

	cloudClients, err := cloud.NewCloudServiceClients(ctx, config)
	if err != nil {
		panic(err)
	}

	state.cloud = cloudClients

	datasetName := "media_ds"
	mediaTableName := "media"
	embeddingTableName := "scene_embeddings"

	state.searchService = &services.SearchService{
		BigqueryClient: cloudClients.BiqQueryClient,
		EmbeddingModel: cloudClients.EmbeddingModels["multi-lingual"],
		DatasetName:    datasetName,
		MediaTable:     mediaTableName,
		EmbeddingTable: embeddingTableName,
	}

	state.mediaService = &services.MediaService{
		BigqueryClient: cloudClients.BiqQueryClient,
		DatasetName:    datasetName,
		MediaTable:     mediaTableName,
	}

	SetupListeners(config, cloudClients, ctx)

}
