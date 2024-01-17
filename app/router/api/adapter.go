package routerapi

import (
	"os"
	"svix-poc/app/router"
	eventsv1 "svix-poc/app/router/api/v1"
	"svix-poc/app/router/repo"
	"svix-poc/package/api/http/server"
	"svix-poc/package/app"

	"github.com/google/wire"
)

func Provide(h *Handler, d router.Dependencies) *server.Adapter {
	return server.NewAdapter(server.AdapterOptions{
		Addr:    string(ProvideHost()),
		Handler: eventsv1.Handler(h),
	}, app.Actors{
		repo.SingletonKey: d.Repository,
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
