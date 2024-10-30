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

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/google/generative-ai-go/genai"
)

// Default System Settings for GenAI agents
var DEFAULT_SAFETY_SETTINGS = []*genai.SafetySetting{
	{
		Category:  genai.HarmCategoryDangerousContent,
		Threshold: genai.HarmBlockOnlyHigh,
	},
	{
		Category:  genai.HarmCategoryHarassment,
		Threshold: genai.HarmBlockMediumAndAbove,
	},
	{
		Category:  genai.HarmCategoryHateSpeech,
		Threshold: genai.HarmBlockMediumAndAbove,
	},
	{
		Category:  genai.HarmCategorySexuallyExplicit,
		Threshold: genai.HarmBlockOnlyHigh,
	},
}

// PubSubListener is a simple stateful wrapper around a subscription object.
// this allows for the easy configuration of multiple listeners. Since listeners
// life-cycles are outside of the command life-cycle they are considered cloud components.
type PubSubListener struct {
	client       *pubsub.Client
	subscription *pubsub.Subscription
	command      cor.Command
}

// NewPubSubListener the constructor for PubSubListener
func NewPubSubListener(
	pubsubClient *pubsub.Client,
	subscriptionID string,
	command cor.Command,
) (cmd *PubSubListener, err error) {

	sub := pubsubClient.Subscription(subscriptionID)

	cmd = &PubSubListener{
		client:       pubsubClient,
		subscription: sub,
		command:      command,
	}
	return cmd, nil
}

// A setter for the underlying handler command.
func (m *PubSubListener) SetCommand(command cor.Command) {
	if m.command == nil {
		m.command = command
	}
}

// Listen starts the async function for listening and should be instantiated
// using the same context of the cloud service but may be configured independently
// for a different recovery life-cycle.
func (m *PubSubListener) Listen(ctx context.Context) {
	log.Printf("listening: %s", m.subscription)
	go func() {
		err := m.subscription.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
			chainCtx := cor.NewBaseContext()
			chainCtx.SetContext(ctx)
			chainCtx.Add(cor.CTX_IN, string(msg.Data))
			m.command.Execute(chainCtx)
			// Only take the message from the topic if the chain executes successfully
			if !chainCtx.HasErrors() {
				msg.Ack()
			}
		})
		if err != nil {
			log.Println(err)
		}
	}()
}

type BigQueryDataSource struct {
	DatasetName string            `toml:"dataset"`
	TableNames  map[string]string `toml:"table_names"`
}

type VertexAiEmbeddingModel struct {
	Model                string `toml:"model"`
	MaxRequestsPerMinute int    `toml:"max_requests_per_minute"`
}

type VertexAiLLMModel struct {
	Model              string  `toml:"model"`
	SystemInstructions string  `toml:"system_instructions"`
	Temperature        float32 `toml:"temperature"`
	TopP               float32 `toml:"top_p"`
	TopK               int32   `toml:"top_k"`
	MaxTokens          int32   `toml:"max_tokens"`
	OutputFormat       string  `toml:"output_format"`
	EnableGoogle       bool    `toml:"enable_google"`
}

type TopicSubscription struct {
	Name             string `toml:"name"`
	DeadLetterTopic  string `toml:"dead_letter_topic"`
	TimeoutInSeconds int    `toml:"timeout_in_seconds"`
	CommandName      string `toml:"command_name"`
}

type CloudConfig struct {
	Application struct {
		Name            string `toml:"name"`
		GoogleProjectId string `toml:"google_project_id"`
		GoogleLocation  string `toml:"location"`
		GoogleAPIKey    string `toml:"google_api_key"`
		ThreadPoolSize  int    `toml:"thread_pool_size"`
	} `toml:"application"`

	TopicSubscriptions  map[string]TopicSubscription      `toml:"topic_subscriptions"`
	BigQueryDataSources map[string]BigQueryDataSource     `toml:"big_query_data_sources"`
	EmbeddingModels     map[string]VertexAiEmbeddingModel `toml:"embedding_models"`
	AgentModels         map[string]VertexAiLLMModel       `toml:"agent_models"`
}

func NewCloudConfig() *CloudConfig {
	return &CloudConfig{
		BigQueryDataSources: make(map[string]BigQueryDataSource),
		EmbeddingModels:     make(map[string]VertexAiEmbeddingModel),
		AgentModels:         make(map[string]VertexAiLLMModel),
	}
}
