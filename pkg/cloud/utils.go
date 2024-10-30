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
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/google/generative-ai-go/genai"
)

// Cloud Constants
const (
	CONFIG_FILE_BASE_NAME  = ".env"
	CONFIG_FILE_EXTENSION  = ".toml"
	CONFIG_SEPARATOR       = "."
	ENV_CONFIG_FILE_PREFIX = "GCP_CONFIG_PREFIX"
	ENV_CONFIG_RUNTIME     = "GCP_RUNTIME"
	MAX_RETRIES            = 3
)

// Simple utility to see if a file exists
func fileExists(in string) bool {
	_, err := os.Stat(in)
	return !errors.Is(err, os.ErrNotExist)
}

// The configuration loader, a hierarchical loader that allows environment overrides.
func LoadConfig(baseConfig interface{}, additionalFiles ...string) {
	configurationFilePrefix := os.Getenv(ENV_CONFIG_FILE_PREFIX)
	if len(configurationFilePrefix) > 0 && !strings.HasSuffix(configurationFilePrefix, string(os.PathSeparator)) {
		configurationFilePrefix = configurationFilePrefix + string(os.PathSeparator)
	}

	runtimeEnvironment := os.Getenv(ENV_CONFIG_RUNTIME)
	if runtimeEnvironment == "" {
		runtimeEnvironment = "test"
	}

	// Read Base Config
	baseConfigFileName := configurationFilePrefix + CONFIG_FILE_BASE_NAME + CONFIG_FILE_EXTENSION
	fmt.Printf("Base Configuration File: %s\n", baseConfigFileName)

	// Override with environment config
	envConfigFileName := configurationFilePrefix + CONFIG_FILE_BASE_NAME + CONFIG_SEPARATOR + runtimeEnvironment + CONFIG_FILE_EXTENSION
	fmt.Printf("Environment Configuration File: %s\n", envConfigFileName)

	if fileExists(baseConfigFileName) {
		_, err := toml.DecodeFile(baseConfigFileName, baseConfig)
		if err != nil {
			log.Fatalf("failed to decode base configuration file %s with error: %s", baseConfigFileName, err)
		}
	}

	if fileExists(envConfigFileName) {
		//envConfig := &Config{}
		_, err := toml.DecodeFile(envConfigFileName, baseConfig)
		if err != nil {
			log.Fatalf("failed to decode environment configuration file: %s with error: %s", envConfigFileName, err)
		}
	}
}

// A helper function for debugging configuration
func PrintConfig(config interface{}) {
	fmt.Println(strings.Repeat("#", 80))
	c, err := json.MarshalIndent(config, " ", " ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(string(c))
	}
	fmt.Println(strings.Repeat("#", 80))
}

// A GenAI helper function for executing multi-modal requests with a retry limit.
func GenerateMultiModalResponse(ctx context.Context, tryCount int, model *genai.GenerativeModel, parts ...genai.Part) (value string, err error) {
	resp, err := model.GenerateContent(ctx, parts...)
	if err != nil {
		if tryCount < MAX_RETRIES {
			return GenerateMultiModalResponse(ctx, tryCount+1, model, parts...)
		} else {
			return "", err
		}
	}
	value = ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				value += fmt.Sprint(part)
			}
		}
	}
	return value, nil
}

// A simple helper method used to adjust float precision for between 64bit and 32bit models.
func GenerateTextEmbeddingAsFloat32(ctx context.Context, model genai.EmbeddingModel, in string, count int) []float32 {
	res, err := model.EmbedContent(ctx, genai.Text(in))
	if err != nil {
		if count < MAX_RETRIES {
			log.Printf("Error generating embeddings [Float32] for: %s on %v", in, err)
			GenerateTextEmbeddingAsFloat32(ctx, model, in, count+1)
		}
	}

	initialValues := res.Embedding.Values
	if len(initialValues) < 768 {
		for {
			initialValues = append(initialValues, 0)
			if len(initialValues) == 768 {
				break
			}
		}
	}

	return initialValues
}

// A simple helper function to stream Markdown generated JSON format
// sometimes returned from Gemini
func RemoveMarkdownJsonNotations(in string) string {
	in = strings.ReplaceAll(in, "```json", "")
	in = strings.ReplaceAll(in, "```", "")
	return in
}

// A helper method for creating text parts
func NewTextPart(in string) genai.Part {
	return genai.Text(in)
}

// A helper method for creating image parts
func NewImagePart(bucketURL string, mimeType string) genai.Part {
	return genai.FileData{URI: bucketURL, MIMEType: mimeType}
}

// A helper method for creating binary large objects / blobs
func NewBlobPart(in []byte, mimeType string) genai.Part {
	return genai.Blob{Data: in, MIMEType: mimeType}
}

// A helper method for creating File Data parts.
func NewFileData(in string, mimeType string) genai.Part {
	return genai.FileData{URI: in, MIMEType: mimeType}
}
