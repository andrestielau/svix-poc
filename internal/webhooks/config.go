package webhooks

import svix "github.com/svix/svix-webhooks/go"

type App struct {
	Def svix.ApplicationIn
	Endpoints []svix.EndpointIn
}
type Config struct {
	Apps   []App
	Events []svix.EventTypeIn
}
