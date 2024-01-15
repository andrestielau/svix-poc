package topic

import (
	"encoding/json"
	"svix-poc/provide/topic"

	"github.com/ThreeDotsLabs/watermill/message"
	svix "github.com/svix/svix-webhooks/go"
)

type Handler struct {
	Client *svix.Svix
	topic.Transformer
}

func (h *Handler) Handle(msg *message.Message) ([]*message.Message, error) {
	var payload map[string]any
	if err := json.Unmarshal(msg.Payload, &payload); err != nil { // TODO: replace with simple http client to avoid redundant conversion
		return nil, err
	} else if _, err = h.Client.Message.Create(msg.Context(), msg.Metadata["app"], &svix.MessageIn{ // Send WebHookS
		EventType: msg.Metadata["EventType"],
		Payload:   payload,
	}); err != nil {
		return nil, err
	}
	return []*message.Message{}, nil
}
