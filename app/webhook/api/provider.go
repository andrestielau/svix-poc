package webhookapi

import (
	"os"
	"svix-poc/app/webhook"
	svixclient "svix-poc/app/webhook/svix"
	"svix-poc/package/app"

	"github.com/google/wire"
)

func Provide(d webhook.Dependencies) *Handler {
	return &Handler{
		Dependencies: d,
		BaseActor: app.NewActor(app.Actors{
			svixclient.SingletonKey: d.SvixClient,
		}),
	}
}

type Host string

var DefaultHost Host = ":2635"

func ProvideHost() Host {
	if url := Host(os.Getenv("WEBHOOK_API_HOST")); url != "" {
		return url
	}
	return DefaultHost
}

var Set = wire.NewSet(Provide)
