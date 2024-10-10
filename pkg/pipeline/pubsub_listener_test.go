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

package pipeline

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/stretchr/testify/assert"
)

func handleErr(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error reading config file: %v", err)
	}
}

type MediaMessageCommand struct {
	model.Command
}

func (c *MediaMessageCommand) IsExecutable(context model.ChainContext) bool {
	return context.Get("message").(model.TriggerMediaWrite).Kind == "storage#object"
}

func (c *MediaMessageCommand) Execute(context model.ChainContext) {
	model := context.Get("message").(model.TriggerMediaWrite)
	log.Println(fmt.Sprintf("Message:\n%v\n", model))
}

func TestMessageHandler(t *testing.T) {
	config, err := model.ReadConfig("configs/app.yaml")
	handleErr(err, t)
	assert.NotNil(t, config)

	ctx, cancel := context.WithCancel(context.Background())

	// Create the external controller group.
	var wg sync.WaitGroup

	pubsubListener, err := NewPubSubListener(&wg, ctx, config.Project.Id, config.HighResSubscription.Id, &MediaMessageCommand{})
	assert.Nil(t, err)
	assert.NotNil(t, pubsubListener)
	pubsubListener.Listen()

	go func() {
		time.Sleep(20 * time.Second)
		cancel()
	}()

	wg.Wait()
}
