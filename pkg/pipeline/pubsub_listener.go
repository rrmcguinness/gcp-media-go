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
	"encoding/json"
	"log"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/model"
)

type PubSubListener struct {
	wg           *sync.WaitGroup
	ctx          context.Context
	client       *pubsub.Client
	subscription *pubsub.Subscription
	command      model.Command
}

func (m *PubSubListener) Listen() {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		err := m.subscription.Receive(m.ctx, func(_ context.Context, msg *pubsub.Message) {
			var out model.TriggerMediaWrite
			err := json.Unmarshal(msg.Data, &out)
			if err == nil {
				chainCtx := model.NewChainContext()
				chainCtx.Add("message", out)
				m.command.Execute(chainCtx)
				// Only take the message from the topic if the chain executes successfully
				if !chainCtx.HasErrors() {
					msg.Ack()
				}
			} else {
				log.Println(err)
				log.Printf("failed to json decode object: %s", string(msg.Data))
			}
		})
		if err != nil {
			log.Println(err)
		}
	}()
}

func NewPubSubListener(
	wg *sync.WaitGroup,
	ctx context.Context,
	projectID string,
	subscriptionID string,
	command model.Command,
) (cmd *PubSubListener, err error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	sub := client.Subscription(subscriptionID)
	cmd = &PubSubListener{
		wg:           wg,
		ctx:          ctx,
		client:       client,
		subscription: sub,
		command:      command,
	}
	return cmd, nil
}
