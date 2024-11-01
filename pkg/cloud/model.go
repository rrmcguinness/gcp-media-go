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
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/google/generative-ai-go/genai"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"golang.org/x/time/rate"
)

// Default System Settings for GenAI agents
var DEFAULT_SAFETY_SETTINGS = []*genai.SafetySetting{
	{
		Category:  genai.HarmCategoryDangerousContent,
		Threshold: genai.HarmBlockNone,
	},
	{
		Category:  genai.HarmCategoryHarassment,
		Threshold: genai.HarmBlockNone,
	},
	{
		Category:  genai.HarmCategoryHateSpeech,
		Threshold: genai.HarmBlockNone,
	},
	{
		Category:  genai.HarmCategorySexuallyExplicit,
		Threshold: genai.HarmBlockNone,
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
		tracer := otel.Tracer("message-listener")
		err := m.subscription.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
			spanCtx, span := tracer.Start(ctx, "receive-message")
			span.SetName("receive-message")
			span.SetAttributes(attribute.String("msg", string(msg.Data)))
			log.Println("received message")
			chainCtx := cor.NewBaseContext()
			chainCtx.SetContext(spanCtx)
			chainCtx.Add(cor.CTX_IN, string(msg.Data))
			m.command.Execute(chainCtx)
			// Only take the message from the topic if the chain executes successfully
			if !chainCtx.HasErrors() {
				span.SetStatus(codes.Ok, "success")
				msg.Ack()
			} else {
				span.SetStatus(codes.Error, "failed")
				for _, e := range chainCtx.GetErrors() {
					log.Printf("error executing chain: %v", e)
				}
			}
			span.End()
		})
		if err != nil {
			log.Printf("error receiving data: %v", err)
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
	RateLimit          int     `toml:"rate_limit"`
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

type QuotaAwareModel struct {
	*genai.GenerativeModel
	RateLimit rate.Limiter
}

func NewQuotaAwareModel(wrapped *genai.GenerativeModel, requestsPerMinute int) *QuotaAwareModel {
	return &QuotaAwareModel{
		GenerativeModel: wrapped,
		RateLimit:       *rate.NewLimiter(rate.Every(time.Minute/1), requestsPerMinute),
	}
}

func (q *QuotaAwareModel) GenerateContent(ctx context.Context, parts ...genai.Part) (resp *genai.GenerateContentResponse, err error) {
	if q.RateLimit.Allow() {
		return q.GenerativeModel.GenerateContent(ctx, parts...)
	} else {
		time.Sleep(time.Second * 5)
		return q.GenerateContent(ctx, parts...)
	}
}
