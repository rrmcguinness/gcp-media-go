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
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	cloud "github.com/GoogleCloudPlatform/solutions/media/pkg/cloud"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
	"github.com/GoogleCloudPlatform/solutions/media/test"
	"github.com/stretchr/testify/assert"
)

type MediaMessageCommand struct {
	cor.Command
}

func (c *MediaMessageCommand) IsExecutable(context cor.Context) bool {
	return context != nil && context.Get("message").(model.TriggerMediaWrite).Kind == "storage#object"
}

func (c *MediaMessageCommand) Execute(context cor.Context) {
	model := context.Get("message").(model.TriggerMediaWrite)
	log.Println(fmt.Sprintf("Message:\n%v\n", model))
}

func TestMessageHandler(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	config := test.GetConfig(t)

	cloudClients, err := cloud.NewCloudServiceClients(ctx, config)
	test.HandleErr(err, t)
	defer cloudClients.Close()

	// Create the external controller group.
	var wg sync.WaitGroup
	wg.Add(1)

	pubsubListener := cloudClients.PubSubListeners["HiResTopic"]
	pubsubListener.SetCommand(&MediaMessageCommand{})

	assert.NotNil(t, pubsubListener)
	pubsubListener.Listen(ctx)

	go func() {
		time.Sleep(10 * time.Second)
		// By calling cancel here, we shut down the Message Listener
		// which in turn signals the WaitGroup that the work is complete.
		wg.Done()
		cancel()
	}()

	wg.Wait()
}
