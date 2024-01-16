package routertopic

import (
	"context"
	"log"
	"svix-poc/app/router"
	"svix-poc/app/router/repo"
	"svix-poc/package/app"
	"svix-poc/package/topic"
	"svix-poc/package/topic/gps"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/google/wire"
	"github.com/samber/lo"
)

type Handler struct {
	router.Dependencies
	*app.BaseActor
}

func Provide(d router.Dependencies) *Handler {
	return &Handler{
		Dependencies: d,
		BaseActor:    app.NewActor(),
	}
}

var Set = wire.NewSet(Provide)

// TODO make dynamic
func Handle() {
	logger := watermill.NewStdLogger(true, false)
	router := lo.Must(message.NewRouter(message.RouterConfig{}, logger))
	router.AddMiddleware(
		middleware.CorrelationID,
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: 100 * time.Millisecond,
			Logger:          logger,
		}.Middleware,
		middleware.Recoverer,
	)
	// Pubsub
	pub := lo.Must(gps.NewPublisher(logger))
	sub := lo.Must(gps.NewSubscriber(logger))
	// (Inbound) Internal Event Handlers
	(&EventNotificationHandler{Topic: "example_topic1", Key: "example1"}).Add(router, sub, pub)
	(&EventNotificationHandler{Topic: "example_topic2", Key: "example2"}).Add(router, sub, pub)
	(&EventNotificationHandler{Topic: "example_topic3", Key: "example3"}).Add(router, sub, pub)
	// (Outboud) Notification Provider Handlers
	(&NotificationProviderHandler{Topic: "webhooks_topic", Key: "webhook"}).Add(router, sub, pub)
	(&NotificationProviderHandler{Topic: "emails_topic", Key: "email"}).Add(router, sub, pub)
	lo.Must0(router.Run(context.Background()))
}

type NotificationProviderHandler struct {
	Topic string
	Key   string
	topic.Transformer
}

func (h *NotificationProviderHandler) Add(router *message.Router, sub message.Subscriber, pub message.Publisher) {
	router.AddHandler(h.Key+"_notification_handler", "notifications_topic", sub, h.Topic, pub, h.Handle)
}

func (h *NotificationProviderHandler) Handle(msg *message.Message) ([]*message.Message, error) {
	// Check Metadata if should be skipped
	log.Println("structHandler received message", msg.UUID)
	msg = message.NewMessage(watermill.NewUUID(), []byte("message produced by provider handler: "+h.Key))
	return message.Messages{msg}, nil
}

type EventNotificationHandler struct {
	router.Dependencies
	Topic string
	Key   string
}

func (h *EventNotificationHandler) Add(router *message.Router, sub message.Subscriber, pub message.Publisher) {
	router.AddHandler(h.Key+"_event_handler", h.Topic, sub, "notifications_topic", pub, h.Handle)
}

func (h *EventNotificationHandler) Handle(msg *message.Message) ([]*message.Message, error) {
	subscriptions, err := h.ListEventSubscriptions(msg.Context(), repo.ListEventSubscriptionsParams{ // Get Subscriptions
		Event: msg.Metadata["EventType"],
	})
	if err != nil {
		return nil, err
	}
	log.Println("event notification handler received message", msg.UUID, subscriptions)
	msg = message.NewMessage(watermill.NewUUID(), []byte("message produced by provider handler: "+h.Key))
	return message.Messages{msg}, nil
}
