package webhooktopic

import (
	"context"
	"log"
	"os"
	"svix-poc/app/webhook"
	svixclient "svix-poc/app/webhook/svix"
	"svix-poc/package/app"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/google/wire"
)

const SubscriberKey = "webhook_subscriber"

func Provide(d webhook.Dependencies) *Handler {
	return &Handler{
		Dependencies: d,
		BaseActor: app.NewActor(app.Actors{
			svixclient.SingletonKey: d.SvixClient,
		}),
	}
}
func (h *Handler) Start(ctx context.Context) (first bool, err error) {
	if first, err = h.BaseActor.Start(ctx); !first || err != nil {
		return first, err
	}
	h.subscriber, err = googlecloud.NewSubscriber(
		googlecloud.SubscriberConfig{
			GenerateSubscriptionName: func(topic string) string {
				return "test-sub_" + topic // TODO: make configurable
			},
			ProjectID: "demo", // TODO: make configurable
		},
		watermill.NewStdLogger(true, false),
	)
	if err != nil {
		return true, err
	}
	topic := string(ProvideTopic())
	log.Println("subscribed to topic: " + topic)
	if h.ch, err = h.subscriber.Subscribe(ctx, topic); err != nil {
		return true, err
	}
	go h.start()
	return true, nil
}
func (h *Handler) start() {
	for msg := range h.ch {
		if msg == nil {
			break
		}
		if err := h.Handle(msg); err != nil {
			log.Println(err)
		}
	}
	close(h.closer)
}
func (h *Handler) Stop(ctx context.Context) (last bool, err error) {
	if last, err = h.BaseActor.Stop(ctx); !last || err != nil {
		return last, err
	}
	if h.subscriber == nil {
		return true, nil
	}
	defer func() { <-h.closer }()
	err = h.subscriber.Close()
	return true, err
}

type Topic string

var DefaultTopic Topic = "webhook"

func ProvideTopic() Topic {
	if url := Topic(os.Getenv("WEBHOOK_PUBSUB_TOPIC")); url != "" {
		return url
	}
	return DefaultTopic
}

var Set = wire.NewSet(Provide)
