package routerapi

import (
	"os"
	"svix-poc/app/router"
	eventsv1 "svix-poc/app/router/api/v1"
	"svix-poc/app/router/repo"
	"svix-poc/lib/api/http/server"
	"svix-poc/lib/app"

	"github.com/google/wire"
)

func Provide(d router.Dependencies) *server.Adapter {
	return server.NewAdapter(server.AdapterOptions{
		Handler: eventsv1.Handler(&Handler{Dependencies: d}),
		Addr:    string(ProvideHost()),
	}, app.Actors{
		repo.SingletonKey: d.Provider,
	})
}

type Host string

var DefaultHost Host = ":3524"

func ProvideHost() Host {
	if url := Host(os.Getenv("ROUTER_API_HOST")); url != "" {
		return url
	}
	return DefaultHost
}

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	Provide,
)
