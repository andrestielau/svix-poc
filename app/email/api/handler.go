package emailapi

import (
	"svix-poc/app/email"
	"svix-poc/package/app"

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
