package webhookgrpc

import (
	"context"
	"svix-poc/app/webhook"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	schemav1 "svix-poc/package/schema/grpc/v1"
)

// TODO: wrap handler with a provider instead of havingt server inside
type Handler struct {
	webhook.Dependencies
}

var _ webhooksv1.WebHookServiceServer = &Handler{}
var _ schemav1.SchemaServiceServer = &Handler{}

// CreateApps implements webhooksv1.WebHookServiceServer.
func (*Handler) CreateApps(context.Context, *webhooksv1.CreateAppsRequest) (*webhooksv1.CreateAppsResponse, error) {
	panic("unimplemented")
}
