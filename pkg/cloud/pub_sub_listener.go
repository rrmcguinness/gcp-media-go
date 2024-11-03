package cloud

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/GoogleCloudPlatform/solutions/media/pkg/cor"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"log"
)

// PubSubListener is a simple stateful wrapper around a subscription object.
// this allows for the easy configuration of multiple listeners. Since listeners
// life-cycles are outside the command life-cycle they are considered cloud components.
type PubSubListener struct {
	client       *pubsub.Client
	subscription *pubsub.Subscription
	command      cor.Command
}

// NewPubSubListener the constructor for PubSubListener
func NewPubSubListener(
	pubsubClient *pubsub.Client,
	subscriptionID string,
	command cor.Command,
) (cmd *PubSubListener, err error) {

	sub := pubsubClient.Subscription(subscriptionID)

	cmd = &PubSubListener{
		client:       pubsubClient,
		subscription: sub,
		command:      command,
	}
	return cmd, nil
}

// SetCommand A setter for the underlying handler command.
func (m *PubSubListener) SetCommand(command cor.Command) {
	if m.command == nil {
		m.command = command
	}
}

// Listen starts the async function for listening and should be instantiated
// using the same context of the cloud service but may be configured independently
// for a different recovery life-cycle.
func (m *PubSubListener) Listen(ctx context.Context) {
	log.Printf("listening: %s", m.subscription)

	go func() {
		tracer := otel.Tracer("message-listener")
		err := m.subscription.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
			spanCtx, span := tracer.Start(ctx, "receive-message")
			span.SetName("receive-message")
			span.SetAttributes(attribute.String("msg", string(msg.Data)))
			log.Println("received message")
			chainCtx := cor.NewBaseContext()
			chainCtx.SetContext(spanCtx)
			chainCtx.Add(cor.CtxIn, string(msg.Data))
			m.command.Execute(chainCtx)
			// Only take the message from the topic if the chain executes successfully
			if !chainCtx.HasErrors() {
				span.SetStatus(codes.Ok, "success")
				msg.Ack()
			} else {
				span.SetStatus(codes.Error, "failed")
				for _, e := range chainCtx.GetErrors() {
					log.Printf("error executing chain: %v", e)
				}
			}
			span.End()
		})
		if err != nil {
			log.Printf("error receiving data: %v", err)
		}
	}()
}
