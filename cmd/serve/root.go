package serve

import (
	"errors"
	"fmt"
	"os"
	"svix-poc/app/router"
	routerapi "svix-poc/app/router/api"
	routergrpc "svix-poc/app/router/grpc"
	routertopic "svix-poc/app/router/topic"
	"svix-poc/app/webhook"
	webhookapi "svix-poc/app/webhook/api"
	webhookgrpc "svix-poc/app/webhook/grpc"
	webhooktopic "svix-poc/app/webhook/topic"
	"svix-poc/package/app"
	"svix-poc/package/app/cmd"
	"svix-poc/package/utils"
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
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085") // TODO: google black magic requires this or an actual env
	sys := app.NewActor(apps)
	ctx := cmd.Context()
	defer sys.Stop(ctx)
	_, err := sys.Start(ctx)
	lo.Must0(err)
	<-utils.WaitSigs(syscall.SIGINT, syscall.SIGTERM)
}

var appProviders = map[string]func() app.Actor{
	"webhook": WebHookApp,
	"router":  RouterApp,
}

func NewRouterApp(a *routerapi.Handler, g *routergrpc.Handler, t *routertopic.Handler, d router.Dependencies) app.Actor {
	return app.NewActor(app.Actors{"api": a, "grpc": g, "topic": t})
}
func NewWebHookApp(a *webhookapi.Handler, g *webhookgrpc.Handler, t *webhooktopic.Handler, d webhook.Dependencies) app.Actor {
	return app.NewActor(app.Actors{"api": a, "grpc": g, "topic": t})
}
