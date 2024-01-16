package router

import (
	"svix-poc/app/router/repo"
	"svix-poc/package/app"

	"github.com/google/wire"
)

type Dependencies struct {
	*repo.Provider
}
type Module struct {
	*app.BaseActor
	Dependencies
}

func New(d Dependencies) *Module {
	return &Module{
		BaseActor:    app.NewActor(),
		Dependencies: d,
	}
}

var Set = wire.NewSet(
	repo.Provide,
	wire.Struct(new(Dependencies), "*"),
)
