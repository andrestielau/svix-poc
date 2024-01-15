package gps

import (
	"svix-poc/module/email/smtp"
	"svix-poc/provide/topic"

	"github.com/ThreeDotsLabs/watermill/message"
)

type Handler struct {
	Client smtp.Client
	topic.Transformer
}

func (h *Handler) Handle(msg *message.Message) ([]*message.Message, error) {
	if _, err := h.Transform(msg); err != nil { // Apply Transformations
		return nil, err
	} else if err = h.Client.Send(msg.Context(), smtp.NewEmail{}); err != nil { // Send Email
		return nil, err
	}
	return []*message.Message{}, nil
}
