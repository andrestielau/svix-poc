package api

import (
	eventsv1 "svix-poc/module/router/api/v1"
	webhookapi "svix-poc/module/webhook/api"
)

type Handler struct {
	webhookapi.WebHooks
	Events
}

var _ eventsv1.ServerInterface = &Handler{}
