package cmd

import (
	"svix-poc/x/app/cmd"

	svix "github.com/svix/svix-webhooks/go"
)

var Root = cmd.New("svix-poc",
	cmd.Add(Demo, Mock, Serve),
)

type App struct {
	Def       svix.ApplicationIn
	Endpoints []svix.EndpointIn
}
type Config struct {
	Apps   []App
	Events []svix.EventTypeIn
}
