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

package cloud

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type CloudServiceClients struct {
	StorageClient   *storage.Client
	PubsubClient    *pubsub.Client
	GenAIClient     *genai.Client
	PubSubListeners map[string]*PubSubListener
	EmbeddingModels map[string]*genai.EmbeddingModel
	AgentModels     map[string]*genai.GenerativeModel
}

func (c *CloudServiceClients) Close() {
	c.StorageClient.Close()
	c.PubsubClient.Close()
	c.GenAIClient.Close()
}

func NewCloudServiceClients(ctx context.Context, config *CloudConfig) (cloud *CloudServiceClients, err error) {
	sc, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	pc, err := pubsub.NewClient(ctx, config.Application.GoogleProjectId)
	if err != nil {
		return nil, err
	}

	gc, err := genai.NewClient(ctx, option.WithAPIKey(config.Application.GoogleAPIKey))
	if err != nil {
		fmt.Printf("Error creating GenAI client: %v\n", err)
		//return nil, err
	}

	subscriptions := make(map[string]*PubSubListener)
	for sub := range config.TopicSubscriptions {
		values := config.TopicSubscriptions[sub]
		actual, err := NewPubSubListener(pc, values.Name, nil)
		if err != nil {
			return nil, err
		}
		subscriptions[sub] = actual
	}

	embeddingModels := make(map[string]*genai.EmbeddingModel)

	for emb := range config.EmbeddingModels {
		embeddingModels[emb] = gc.EmbeddingModel(config.EmbeddingModels[emb].Model)
	}

	agentModels := make(map[string]*genai.GenerativeModel)
	for am := range config.AgentModels {
		values := config.AgentModels[am]
		model := gc.GenerativeModel(values.Model)
		model.SetTemperature(values.Temperature)
		model.SetTopK(values.TopK)
		model.SetTopP(values.TopP)
		model.SetMaxOutputTokens(values.MaxTokens)
		model.SystemInstruction = &genai.Content{
			Parts: []genai.Part{genai.Text(values.SystemInstructions)},
		}
		model.ResponseMIMEType = values.OutputFormat
		model.Tools = []*genai.Tool{}
		agentModels[am] = model
	}

	cloud = &CloudServiceClients{
		StorageClient:   sc,
		PubsubClient:    pc,
		GenAIClient:     gc,
		PubSubListeners: subscriptions,
		EmbeddingModels: embeddingModels,
		AgentModels:     agentModels,
	}

	return cloud, err
}
