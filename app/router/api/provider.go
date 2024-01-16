package routerapi

import (
	"os"
	"svix-poc/app/router"
	"svix-poc/package/app"

	"github.com/google/wire"
)

func Provide(d router.Dependencies) *Handler {
	return &Handler{
		Dependencies: d,
		BaseActor:    app.NewActor(),
	}
}

type Host string

var DefaultHost Host = ""

func ProvideHost() Host {
	if url := Host(os.Getenv("ROUTER_HOST")); url != "" {
		return url
	}
	return DefaultHost
}

var Set = wire.NewSet(Provide)
