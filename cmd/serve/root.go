package serve

import (
	"errors"
	"fmt"
	"log"
	"os"
	"svix-poc/app/router"
	routertopic "svix-poc/app/router/topic"
	"svix-poc/app/webhook"
	webhooktopic "svix-poc/app/webhook/topic"
	grpcserver "svix-poc/lib/api/grpc/server"
	httpserver "svix-poc/lib/api/http/server"
	"svix-poc/lib/app"
	"svix-poc/lib/app/cmd"
	"svix-poc/lib/utils"
	"syscall"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var Root = cmd.New("serve",
	cmd.Alias("s"),
	cmd.Run(runServe),
)

func runServe(cmd *cobra.Command, modules []string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error: %+v", fmt.Errorf("%v", err))
		}
	}()
	if len(modules) == 0 {
		modules = lo.Keys(appProviders)
	}
	apps := app.Actors{}
	for _, module := range modules {
		if provider, ok := appProviders[module]; !ok {
			panic(errors.New("invalid module: " + module))
		} else {
			apps[module] = provider()
		}
	}
	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085") // TODO: google black magic requires this or an actual env
	}
	sys := app.NewActor(apps)
	ctx := cmd.Context()
	defer func() {
		if _, err := sys.Stop(ctx); err != nil {
			log.Println(err)
		}
	}()
	lo.Must(sys.Start(ctx))
	<-utils.WaitSigs(syscall.SIGINT, syscall.SIGTERM)
}

var appProviders = map[string]func() app.Actor{
	"webhook": WebHookApp,
	"router":  RouterApp,
}

func NewRouterApp(a *httpserver.Adapter, g *grpcserver.Adapter, t *routertopic.Adapter, d router.Dependencies) app.Actor {
	return app.NewActor(app.Actors{"api": a, "grpc": g, "topic": t})
}
func NewWebHookApp(a *httpserver.Adapter, g *grpcserver.Adapter, t *webhooktopic.Adapter, d webhook.Dependencies) app.Actor {
	return app.NewActor(app.Actors{"api": a, "grpc": g, "topic": t})
}
