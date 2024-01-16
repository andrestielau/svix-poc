package routertopic

import (
	"log"
	"svix-poc/app/router"
	"svix-poc/app/router/repo"
	"svix-poc/package/app"
	"svix-poc/package/topic"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type Handler struct {
	router.Dependencies
	*app.BaseActor
	router     *message.Router
	publisher  message.Publisher
	subscriber message.Subscriber
}

type NotificationProviderHandler struct {
	Topic string
	Key   string
	topic.Transformer
}

func (h *NotificationProviderHandler) Add(router *message.Router, sub message.Subscriber, pub message.Publisher) {
	router.AddHandler(h.Key+"_notification_handler", "notification", sub, h.Topic, pub, h.Handle)
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
	router.AddHandler(h.Key+"_event_handler", h.Topic, sub, "notification", pub, h.Handle)
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
