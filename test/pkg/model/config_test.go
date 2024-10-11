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

package model

import (
	"testing"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/stretchr/testify/assert"
)

const test_config_file = "configs/app-test.yaml"

func TestConfig(t *testing.T) {
	config, err := model.ReadConfig(test_config_file)
	if err != nil {
		t.Errorf("Error reading config file: %v", err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "test", config.Project.Id)
	assert.Equal(t, "test", config.Project.Location)
	assert.Equal(t, "test", config.HighResSubscription.Id)
	assert.Equal(t, "test", config.HighResToLowResCommand.Bucket)
	assert.Equal(t, "240", config.HighResToLowResCommand.Width)
	assert.Equal(t, "mp4", config.HighResToLowResCommand.Format)
	assert.Equal(t, 1, len(config.HighResToLowResCommand.AdditionalArgs))
}
