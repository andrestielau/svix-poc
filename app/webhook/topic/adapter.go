package webhooktopic

import (
	"context"
	"log"
	"os"
	"svix-poc/app/webhook"
	svixclient "svix-poc/app/webhook/svix"
	"svix-poc/lib/app"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/wire"
)

const SubscriberKey = "webhook_subscriber"

type Adapter struct {
	*app.BaseActor
	h          *Handler
	subscriber message.Subscriber
	closer     chan struct{}
}

func Provide(d webhook.Dependencies) *Adapter {
	return &Adapter{
		h: &Handler{Dependencies: d},
		BaseActor: app.NewActor(app.Actors{
			svixclient.SingletonKey: d.SvixClient,
		}),
	}
}

func (p *Adapter) Start(ctx context.Context) (first bool, err error) {
	if first, err = p.BaseActor.Start(ctx); !first || err != nil {
		return first, err
	}
	p.subscriber, err = googlecloud.NewSubscriber(
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
	ch, err := p.subscriber.Subscribe(ctx, topic)
	if err != nil {
		return true, err
	}
	defer close(p.closer)
	for msg := range ch {
		if msg == nil {
			break
		} else if err := p.h.Handle(msg); err != nil {
			log.Println(err)
		}
	}
	return true, nil
}

func (h *Adapter) Stop(ctx context.Context) (last bool, err error) {
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
