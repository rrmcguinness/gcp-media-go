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
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Project struct {
		Id       string `yaml:"id"`
		Location string `yaml:"location"`
	} `yaml:"project"`

	HighResSubscription struct {
		Id string `yaml:"id"`
	} `yaml:"high_res_subscription"`

	HighResToLowResCommand struct {
		Bucket         string   `yaml:"bucket"`
		Format         string   `yaml:"format"`
		Width          string   `yaml:"width"`
		AdditionalArgs []string `yaml:"additional_args"`
	} `yaml:"high_res_to_low_res_command"`
}

func ReadConfig(configFilePath string) (*Config, error) {
	f, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Decode the YAML file into a Config struct
	var config Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
