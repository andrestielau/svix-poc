package cmd

import (
	"svix-poc/cmd/serve"
	"svix-poc/package/app/cmd"

	svix "github.com/svix/svix-webhooks/go"
)

var Root = cmd.New("svix-poc",
	cmd.Add(Demo, Mock, serve.Root),
)

type App struct {
	Def       svix.ApplicationIn
	Endpoints []svix.EndpointIn
}
type Config struct {
	Apps   []App
	Events []svix.EventTypeIn
}
