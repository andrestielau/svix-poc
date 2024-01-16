package routergrpc

import (
	"svix-poc/app/router"
	"svix-poc/package/app"

	"github.com/google/wire"
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
