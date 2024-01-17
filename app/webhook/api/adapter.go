package webhookapi

import (
	"os"
	"svix-poc/app/webhook"
	webhooksv1 "svix-poc/app/webhook/api/v1"
	svixclient "svix-poc/app/webhook/svix"
	"svix-poc/package/api/http/server"
	"svix-poc/package/app"

	"github.com/google/wire"
)

func Provide(h *Handler, d webhook.Dependencies) *server.Adapter {
	return server.NewAdapter(server.AdapterOptions{
		Addr:    string(ProvideHost()),
		Handler: webhooksv1.Handler(h),
	}, app.Actors{
		svixclient.SingletonKey: d.SvixClient,
	})
}

type Host string

var DefaultHost Host = ":2635"

func ProvideHost() Host {
	if url := Host(os.Getenv("WEBHOOK_API_HOST")); url != "" {
		return url
	}
	return DefaultHost
}

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	Provide,
)
