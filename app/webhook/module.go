package webhook

import (
	svixclient "svix-poc/app/webhook/svix"
	"svix-poc/package/app"

	"github.com/google/wire"
)

type Dependencies struct {
	*svixclient.Provider
}

type Module struct {
	*app.BaseActor
	Dependencies
}

func New(d Dependencies) *Module {
	m := &Module{
		BaseActor:    app.NewActor(),
		Dependencies: d,
	}
	m.Spawn("svix_client", d.Provider)
	return m
}

var Set = wire.NewSet(
	svixclient.Provide,
	wire.Struct(new(Dependencies), "*"),
)
