package emailgrpc

import (
	"svix-poc/app/email"
	"svix-poc/lib/app"

	"github.com/google/wire"
)

type Handler struct {
	email.Dependencies
	*app.BaseActor
}

func Provide(d email.Dependencies) *Handler {
	return &Handler{
		Dependencies: d,
		BaseActor:    app.NewActor(nil),
	}
}

var Set = wire.NewSet(Provide)
