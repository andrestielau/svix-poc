package webhookgrpc

import (
	"context"
	"svix-poc/app/webhook"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	schemav1 "svix-poc/lib/schema/grpc/v1"
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

// CreateEndpoints implements webhooksv1.WebHookServiceServer.
func (*Handler) CreateEndpoints(context.Context, *webhooksv1.CreateEndpointsRequest) (*webhooksv1.CreateEndpointsResponse, error) {
	panic("unimplemented")
}

// CreateEventTypes implements webhooksv1.WebHookServiceServer.
func (*Handler) CreateEventTypes(context.Context, *webhooksv1.CreateEventTypesRequest) (*webhooksv1.CreateEventTypesResponse, error) {
	panic("unimplemented")
}

// CreateMessages implements webhooksv1.WebHookServiceServer.
func (*Handler) CreateMessages(context.Context, *webhooksv1.CreateMessagesRequest) (*webhooksv1.CreateMessagesResponse, error) {
	panic("unimplemented")
}

// ListApps implements webhooksv1.WebHookServiceServer.
func (*Handler) ListApps(context.Context, *webhooksv1.ListAppsRequest) (*webhooksv1.ListAppsResponse, error) {
	panic("unimplemented")
}

// ListEndpoints implements webhooksv1.WebHookServiceServer.
func (*Handler) ListEndpoints(context.Context, *webhooksv1.ListEndpointsRequest) (*webhooksv1.ListEndpointsResponse, error) {
	panic("unimplemented")
}

// ListEventTypes implements webhooksv1.WebHookServiceServer.
func (*Handler) ListEventTypes(context.Context, *webhooksv1.ListEventTypesRequest) (*webhooksv1.ListEventTypesResponse, error) {
	panic("unimplemented")
}

// ListMessages implements webhooksv1.WebHookServiceServer.
func (*Handler) ListMessages(context.Context, *webhooksv1.ListMessagesRequest) (*webhooksv1.ListMessagesResponse, error) {
	panic("unimplemented")
}
