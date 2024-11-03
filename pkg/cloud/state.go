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
	"log"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// ServiceClients is the state machine for the cloud clients.
type ServiceClients struct {
	StorageClient   *storage.Client                         // The Google Cloud Storage client.
	PubsubClient    *pubsub.Client                          // The Google Cloud Pub/Sub client.
	GenAIClient     *genai.Client                           // The Google Cloud Vertex AI client.
	BiqQueryClient  *bigquery.Client                        // The Google Cloud BigQuery client.
	PubSubListeners map[string]*PubSubListener              // A map of Pub/Sub listeners, keyed by subscription name.
	EmbeddingModels map[string]*genai.EmbeddingModel        // A map of Vertex AI embedding models, keyed by model name.
	AgentModels     map[string]*QuotaAwareGenerativeAIModel // A map of Vertex AI LLM models, keyed by model name.
}

// Close A close method to ensure all clients are shut down,
// these are handled using a closable context, but here for clean testing.
func (c *ServiceClients) Close() {
	_ = c.StorageClient.Close()
	_ = c.PubsubClient.Close()
	_ = c.GenAIClient.Close()
	_ = c.BiqQueryClient.Close()
}

// NewCloudServiceClients A helper function for correctly initializing the Google Cloud Services based on the configuration.
func NewCloudServiceClients(ctx context.Context, config *Config) (cloud *ServiceClients, err error) {
	// Create a new Google Cloud Storage client.
	sc, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	// Create a new Google Cloud Pub/Sub client.
	pc, err := pubsub.NewClient(ctx, config.Application.GoogleProjectId)
	if err != nil {
		return nil, err
	}

	// Create a new Google Cloud Vertex AI client.
	gc, err := genai.NewClient(ctx, option.WithAPIKey(config.Application.GoogleAPIKey))
	if err != nil {
		log.Printf("error creating genai client: %v", err)
		return nil, err
	}

	// Create a new Google Cloud BigQuery client.
	bc, err := bigquery.NewClient(ctx, config.Application.GoogleProjectId)
	if err != nil {
		return nil, err
	}

	// Create Pub/Sub listeners based on the configuration.
	subscriptions := make(map[string]*PubSubListener)
	for sub := range config.TopicSubscriptions {
		values := config.TopicSubscriptions[sub]
		actual, err := NewPubSubListener(pc, values.Name, nil)
		if err != nil {
			return nil, err
		}
		subscriptions[sub] = actual
	}

	// Create Vertex AI embedding models based on the configuration.
	embeddingModels := make(map[string]*genai.EmbeddingModel)
	for emb := range config.EmbeddingModels {
		embeddingModels[emb] = gc.EmbeddingModel(config.EmbeddingModels[emb].Model)
	}

	// Create Vertex AI LLM models based on the configuration.
	agentModels := make(map[string]*QuotaAwareGenerativeAIModel)
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
		model.SafetySettings = DefaultSafetySettings
		model.ResponseMIMEType = values.OutputFormat
		model.Tools = []*genai.Tool{}
		wrappedAgent := NewQuotaAwareModel(model, values.RateLimit)
		agentModels[am] = wrappedAgent
	}

	// Create a new ServiceClients instance with all the initialized clients.
	cloud = &ServiceClients{
		StorageClient:   sc,
		PubsubClient:    pc,
		GenAIClient:     gc,
		BiqQueryClient:  bc,
		PubSubListeners: subscriptions,
		EmbeddingModels: embeddingModels,
		AgentModels:     agentModels,
	}

	return cloud, err
}
