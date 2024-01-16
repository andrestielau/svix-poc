package routertopic

import (
	"context"
	"log"
	"svix-poc/app/router"
	"svix-poc/app/router/repo"
	"svix-poc/package/app"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/google/wire"
)

func Provide(d router.Dependencies) *Handler {
	return &Handler{
		Dependencies: d,
		BaseActor: app.NewActor(app.Actors{
			repo.SingletonKey: d.Repository,
		}),
	}
}

func (h *Handler) Start(ctx context.Context) (first bool, err error) {
	if first, err = h.BaseActor.Start(ctx); !first || err != nil {
		return first, err
	}
	logger := watermill.NewStdLogger(true, false)
	if h.subscriber, err = googlecloud.NewSubscriber(
		googlecloud.SubscriberConfig{
			GenerateSubscriptionName: func(topic string) string {
				return "router_" + topic
			},
			ProjectID: "demo",
		},
		logger,
	); err != nil {
		return true, err
	}
	if h.publisher, err = googlecloud.NewPublisher(googlecloud.PublisherConfig{
		ProjectID: "demo",
	}, logger); err != nil {
		return true, err
	}
	if h.router, err = message.NewRouter(message.RouterConfig{}, logger); err != nil {
		return true, err
	}
	h.router.AddMiddleware(
		middleware.CorrelationID,
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: 100 * time.Millisecond,
			Logger:          logger,
		}.Middleware,
		middleware.Recoverer,
	)
	// (Inbound) Internal Event Handlers
	(&EventNotificationHandler{Topic: "example1", Key: "example1"}).Add(h.router, h.subscriber, h.publisher)
	(&EventNotificationHandler{Topic: "example2", Key: "example2"}).Add(h.router, h.subscriber, h.publisher)
	(&EventNotificationHandler{Topic: "example3", Key: "example3"}).Add(h.router, h.subscriber, h.publisher)
	// (Outboud) Notification Provider Handlers
	(&NotificationProviderHandler{Topic: "webhook", Key: "webhook"}).Add(h.router, h.subscriber, h.publisher)
	(&NotificationProviderHandler{Topic: "email", Key: "email"}).Add(h.router, h.subscriber, h.publisher)
	go func() {
		if err := h.router.Run(ctx); err != nil {
			log.Println(err)
		}
	}()
	return true, nil
}

func (h *Handler) Stop(ctx context.Context) (last bool, err error) {
	if last, err = h.BaseActor.Stop(ctx); !last || err != nil {
		return last, err
	}
	return true, h.router.Close()
}

var Set = wire.NewSet(Provide)