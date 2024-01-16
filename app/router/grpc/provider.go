package routergrpc

import (
	"os"
	"svix-poc/app/router"
	"svix-poc/app/router/repo"
	"svix-poc/package/app"

	"github.com/google/wire"
)

func Provide(d router.Dependencies) *Handler {
	return &Handler{
		Dependencies: d,
		BaseActor: app.NewActor(app.Actors{
			repo.SingletonKey: d.Repository,
		}),
	}
}

type Host string

var DefaultHost Host = ":2573"

func ProvideHost() Host {
	if url := Host(os.Getenv("ROUTER_GRPC_HOST")); url != "" {
		return url
	}
	return DefaultHost
}

var Set = wire.NewSet(Provide)
