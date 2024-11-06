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

// QuotaAwareGenerativeAIModel wraps a genai.GenerativeModel with rate limiting.
type QuotaAwareGenerativeAIModel struct {
	*genai.GenerativeModel              // The wrapped Vertex AI LLM.
	RateLimit              rate.Limiter // The rate limiter for the LLM.
}

// NewQuotaAwareModel creates a new QuotaAwareGenerativeAIModel with the given rate limit.
func NewQuotaAwareModel(wrapped *genai.GenerativeModel, requestsPerSecond int) *QuotaAwareGenerativeAIModel {
	return &QuotaAwareGenerativeAIModel{
		GenerativeModel: wrapped,
		RateLimit:       *rate.NewLimiter(rate.Every(time.Second/1), requestsPerSecond),
	}
}

// GenerateContent generates content using the wrapped LLM with rate limiting.
func (q *QuotaAwareGenerativeAIModel) GenerateContent(ctx context.Context, parts ...genai.Part) (resp *genai.GenerateContentResponse, err error) {
	// Check if the rate limit allows a request.
	if q.RateLimit.Allow() {
		// If allowed, make the request to the LLM.
		resp, err = q.GenerativeModel.GenerateContent(ctx, parts...)
		if err != nil {
			// If there's an error, check the retry count from the context.
			retryCount := ctx.Value("retry").(int)
			if retryCount > 3 {
				// If retry count exceeds the limit, return an error.
				return nil, errors.New("failed generation on max retries")
			}
			// If retries are allowed, wait for one minute and try again.
			errCtx := context.WithValue(ctx, "retry", retryCount+1)
			time.Sleep(time.Minute * 1)
			return q.GenerateContent(errCtx, parts...)
		}
		// If successful, return the response.
		return resp, err
	} else {
		// If rate limit is exceeded, wait for 5 seconds and try again.
		time.Sleep(time.Second * 5)
		return q.GenerateContent(ctx, parts...)
	}
}
