package webhooktopic

import (
	"encoding/json"
	"log"
	"svix-poc/app/webhook"

	"github.com/ThreeDotsLabs/watermill/message"
	svix "github.com/svix/svix-webhooks/go"
)

type Handler struct {
	webhook.Dependencies
}

func (h *Handler) Handle(msg *message.Message) error {
	var payload map[string]any
	log.Println(string(msg.Payload))
	if err := json.Unmarshal(msg.Payload, &payload); err != nil { // TODO: replace with simple http client to avoid redundant conversion
		return err
	} else if _, err = h.Message.Create(msg.Context(), msg.Metadata["app"], &svix.MessageIn{ // Send WebHookS
		EventType: msg.Metadata["EventType"],
		Payload:   payload,
	}); err != nil {
		return err
	}
	return nil
}
