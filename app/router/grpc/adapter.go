package routergrpc

import (
	"os"
	"svix-poc/app/router"
	eventsv1 "svix-poc/app/router/grpc/v1"
	"svix-poc/app/router/repo"
	"svix-poc/lib/api/grpc/server"
	"svix-poc/lib/app"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

func Provide(d router.Dependencies) *server.Adapter {
	return server.NewAdapter(server.AdapterOptions{
		Addr: string(ProvideHost()),
		Register: func(s *grpc.Server) {
			eventsv1.RegisterEventServiceServer(s, &Handler{Dependencies: d})
		},
	}, app.Actors{
		repo.SingletonKey: d.Provider,
	})
}

type Host string

var DefaultHost Host = ":2573"

func ProvideHost() Host {
	if url := Host(os.Getenv("ROUTER_GRPC_HOST")); url != "" {
		return url
	}
	return DefaultHost
}

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	Provide,
)
