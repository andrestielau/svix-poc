package webhookgrpc

import (
	"os"
	"svix-poc/app/webhook"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	svixclient "svix-poc/app/webhook/svix"
	"svix-poc/package/api/grpc/server"
	"svix-poc/package/app"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

type Host string

var DefaultHost Host = ":4315"

func ProvideHost() Host {
	if url := Host(os.Getenv("ROUTER_GRPC_HOST")); url != "" {
		return url
	}
	return DefaultHost
}

func Provide(d webhook.Dependencies) *server.Adapter {
	return server.NewAdapter(server.AdapterOptions{
		Addr: string(ProvideHost()),
		Register: func(s *grpc.Server) {
			webhooksv1.RegisterWebHookServiceServer(s, &Handler{Dependencies: d})
		},
	}, app.Actors{
		svixclient.SingletonKey: d.SvixClient,
	})
}

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	Provide,
)
