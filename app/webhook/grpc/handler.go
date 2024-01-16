package webhookgrpc

import (
	"svix-poc/app/webhook"
	"svix-poc/package/app"

	"github.com/google/wire"
)

type Handler struct {
	webhook.Dependencies
	*app.BaseActor
}

func Provide(d webhook.Dependencies) *Handler {
	return &Handler{
		Dependencies: d,
		BaseActor:    app.NewActor(),
	}
}

var Set = wire.NewSet(Provide)
