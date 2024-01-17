package emailtopic

import (
	"svix-poc/app/email"
	"svix-poc/app/email/smtp"
	"svix-poc/lib/app"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/wire"
)

type Handler struct {
	email.Dependencies
	*app.BaseActor
}

func Provide(d email.Dependencies) *Handler {
	return &Handler{
		Dependencies: d,
		BaseActor:    app.NewActor(),
	}
}

var Set = wire.NewSet(Provide)

func (h *Handler) Handle(msg *message.Message) ([]*message.Message, error) {
	if _, err := h.Transform(msg); err != nil { // Apply Transformations
		return nil, err
	} else if err = h.Client.Send(msg.Context(), smtp.NewEmail{}); err != nil { // Send Email
		return nil, err
	}
	return []*message.Message{}, nil
}
