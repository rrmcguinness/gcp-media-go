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

package workflow_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/telemetry"
	"github.com/GoogleCloudPlatform/solutions/media/test"
	"github.com/google/generative-ai-go/genai"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
)

var err error
var genModel *genai.GenerativeModel
var embeddingModel *genai.EmbeddingModel
var cloudClients *cloud.CloudServiceClients
var ctx context.Context

const tName = "cloud.google.com/media/tests/workflow"

var (
	tracer = otel.Tracer(tName)
	meter  = otel.Meter(tName)
	logger = otelslog.NewLogger(tName)
)

func TestMain(m *testing.M) {
	ctx, cancel := context.WithCancel(context.Background())
	// This deferral will automatically close the client that was build from
	// the same context
	defer cancel()

	// Get the config file
	config := test.GetConfig()

	telemetry.SetupLogging()
	shutdown, err := telemetry.SetupOpenTelemetry(ctx, config)
	if err != nil {
		panic(err)
	}

	cloudClients, err = cloud.NewCloudServiceClients(ctx, config)
	if err != nil {
		panic(err)
	}
	defer cloudClients.Close()

	genModel = cloudClients.AgentModels["creative-flash"]
	if genModel == nil {
		panic(errors.New("failed to create Generative model"))
	}

	embeddingModel = cloudClients.EmbeddingModels["multi-lingual"]

	logger.Info("completed test setup")
	fmt.Println("Completed Setup")

	exitCode := m.Run()
	shutdown(ctx)
	os.Exit(exitCode)
}
