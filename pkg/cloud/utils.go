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
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/google/generative-ai-go/genai"
)

// Cloud Constants
const (
	ConfigFileBaseName  = ".env"
	ConfigFileExtension = ".toml"
	ConfigSeparator     = "."
	EnvConfigFilePrefix = "GCP_CONFIG_PREFIX"
	EnvConfigRuntime    = "GCP_RUNTIME"
	MaxRetries          = 3
)

// Simple utility to see if a file exists
func fileExists(in string) bool {
	_, err := os.Stat(in)
	return !errors.Is(err, os.ErrNotExist)
}

// LoadConfig The configuration loader, a hierarchical loader that allows environment overrides.
func LoadConfig(baseConfig interface{}) {
	configurationFilePrefix := os.Getenv(EnvConfigFilePrefix)
	if len(configurationFilePrefix) > 0 && !strings.HasSuffix(configurationFilePrefix, string(os.PathSeparator)) {
		configurationFilePrefix = configurationFilePrefix + string(os.PathSeparator)
	}

	runtimeEnvironment := os.Getenv(EnvConfigRuntime)
	if runtimeEnvironment == "" {
		runtimeEnvironment = "test"
	}

	// Read Base Config
	baseConfigFileName := configurationFilePrefix + ConfigFileBaseName + ConfigFileExtension
	fmt.Printf("Base Configuration File: %s\n", baseConfigFileName)

	// Override with environment config
	envConfigFileName := configurationFilePrefix + ConfigFileBaseName + ConfigSeparator + runtimeEnvironment + ConfigFileExtension
	fmt.Printf("Environment Configuration File: %s\n", envConfigFileName)

	if fileExists(baseConfigFileName) {
		_, err := toml.DecodeFile(baseConfigFileName, baseConfig)
		if err != nil {
			log.Fatalf("failed to decode base configuration file %s with error: %s", baseConfigFileName, err)
		}
	}

	if fileExists(envConfigFileName) {
		_, err := toml.DecodeFile(envConfigFileName, baseConfig)
		if err != nil {
			log.Fatalf("failed to decode environment configuration file: %s with error: %s", envConfigFileName, err)
		}
	}
}

// GenerateMultiModalResponse A GenAI helper function for executing multi-modal requests with a retry limit.
func GenerateMultiModalResponse(ctx context.Context, tryCount int, model *QuotaAwareGenerativeAIModel, parts ...genai.Part) (value string, err error) {
	resp, err := model.GenerateContent(ctx, parts...)
	if err != nil {
		if tryCount < MaxRetries {
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

// NewTextPart A delegate method for creating text parts
func NewTextPart(in string) genai.Part {
	return genai.Text(in)
}

// NewFileData A delegate method for creating File Data parts.
func NewFileData(in string, mimeType string) genai.Part {
	return genai.FileData{URI: in, MIMEType: mimeType}
}
