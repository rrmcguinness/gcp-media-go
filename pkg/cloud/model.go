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
	"errors"
	"time"

	"github.com/google/generative-ai-go/genai"
	"golang.org/x/time/rate"
)

// DefaultSafetySettings Default System Settings for GenAI agents
var DefaultSafetySettings = []*genai.SafetySetting{
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

type BigQueryDataSource struct {
	DatasetName    string `toml:"dataset"`
	MediaTable     string `toml:"media_table"`
	EmbeddingTable string `toml:"embedding_table"`
}

type PromptTemplates struct {
	SummaryPrompt string `toml:"summary"`
	ScenePrompt   string `toml:"scene"`
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
}

type Storage struct {
	HiResInputBucket   string `toml:"high_res_input_bucket"`
	LowResOutputBucket string `toml:"low_res_output_bucket"`
}

type Config struct {
	Application struct {
		Name            string `toml:"name"`
		GoogleProjectId string `toml:"google_project_id"`
		GoogleLocation  string `toml:"location"`
		GoogleAPIKey    string `toml:"google_api_key"`
		ThreadPoolSize  int    `toml:"thread_pool_size"`
	} `toml:"application"`
	Storage            Storage                           `toml:"storage"`
	BigQueryDataSource BigQueryDataSource                `toml:"big_query_data_source"`
	PromptTemplates    PromptTemplates                   `toml:"prompt_templates"`
	TopicSubscriptions map[string]TopicSubscription      `toml:"topic_subscriptions"`
	EmbeddingModels    map[string]VertexAiEmbeddingModel `toml:"embedding_models"`
	AgentModels        map[string]VertexAiLLMModel       `toml:"agent_models"`
}

func NewConfig() *Config {
	return &Config{
		TopicSubscriptions: make(map[string]TopicSubscription),
		EmbeddingModels:    make(map[string]VertexAiEmbeddingModel),
		AgentModels:        make(map[string]VertexAiLLMModel),
	}
}

type QuotaAwareGenerativeAIModel struct {
	*genai.GenerativeModel
	RateLimit rate.Limiter
}

func NewQuotaAwareModel(wrapped *genai.GenerativeModel, requestsPerSecond int) *QuotaAwareGenerativeAIModel {
	return &QuotaAwareGenerativeAIModel{
		GenerativeModel: wrapped,
		RateLimit:       *rate.NewLimiter(rate.Every(time.Second/1), requestsPerSecond),
	}
}

func (q *QuotaAwareGenerativeAIModel) GenerateContent(ctx context.Context, parts ...genai.Part) (resp *genai.GenerateContentResponse, err error) {
	if q.RateLimit.Allow() {
		resp, err = q.GenerativeModel.GenerateContent(ctx, parts...)
		if err != nil {
			retryCount := ctx.Value("retry").(int)
			if retryCount > 3 {
				return nil, errors.New("failed generation on max retries")
			}
			// Wait for one minute and try again
			errCtx := context.WithValue(ctx, "retry", retryCount+1)
			time.Sleep(time.Minute * 1)
			return q.GenerateContent(errCtx, parts...)
		}
		return resp, err
	} else {
		time.Sleep(time.Second * 5)
		return q.GenerateContent(ctx, parts...)
	}
}
