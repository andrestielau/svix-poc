//go:build wireinject
// +build wireinject

package serve

import (
	"svix-poc/app/router"
	routerapi "svix-poc/app/router/api"
	routergrpc "svix-poc/app/router/grpc"
	routertopic "svix-poc/app/router/topic"
	"svix-poc/app/webhook"
	webhookapi "svix-poc/app/webhook/api"
	webhookgrpc "svix-poc/app/webhook/grpc"
	webhooktopic "svix-poc/app/webhook/topic"
	"svix-poc/package/app"

	"github.com/google/wire"
)

func RouterApp() app.Actor {
	panic(wire.Build(
		router.Set,
		routerapi.Set,
		routergrpc.Set,
		routertopic.Set,
		NewRouterApp,
	))
}

func WebHookApp() app.Actor {
	panic(wire.Build(
		webhook.Set,
		webhookapi.Set,
		webhookgrpc.Set,
		webhooktopic.Set,
		NewWebHookApp,
	))
}
