package webhooktopic

import (
	"encoding/json"
	"svix-poc/app/webhook"
	svixclient "svix-poc/app/webhook/svix"
	"svix-poc/package/app"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/wire"
	svix "github.com/svix/svix-webhooks/go"
)

type Handler struct {
	webhook.Dependencies
	*app.BaseActor
}

func Provide(d webhook.Dependencies) *Handler {
	return &Handler{
		Dependencies: d,
		BaseActor: app.NewActor(app.Actors{
			svixclient.SingletonKey: d.SvixClient,
		}),
	}
}

var Set = wire.NewSet(Provide)

func (h *Handler) Handle(msg *message.Message) ([]*message.Message, error) {
	var payload map[string]any
	if err := json.Unmarshal(msg.Payload, &payload); err != nil { // TODO: replace with simple http client to avoid redundant conversion
		return nil, err
	} else if _, err = h.Message.Create(msg.Context(), msg.Metadata["app"], &svix.MessageIn{ // Send WebHookS
		EventType: msg.Metadata["EventType"],
		Payload:   payload,
	}); err != nil {
		return nil, err
	}
	return []*message.Message{}, nil
}

func WithSubscriber() {

}
