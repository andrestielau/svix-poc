package routergrpc

import (
	"context"
	"svix-poc/app/router"
	eventsv1 "svix-poc/app/router/grpc/v1"
	schemav1 "svix-poc/lib/schema/grpc/v1"
)

type Handler struct {
	router.Dependencies
}

var _ eventsv1.EventServiceServer = &Handler{}
var _ schemav1.SchemaServiceServer = &Handler{}

// CreateEventTypes implements eventsv1.EventServiceServer.
func (*Handler) CreateEventTypes(context.Context, *eventsv1.CreateEventTypesRequest) (*eventsv1.CreateEventTypesResponse, error) {
	panic("unimplemented")
}

// CreateSubscriptions implements eventsv1.EventServiceServer.
func (*Handler) CreateSubscriptions(context.Context, *eventsv1.CreateSubscriptionsRequest) (*eventsv1.CreateSubscriptionsResponse, error) {
	panic("unimplemented")
}

// ListEventTypes implements eventsv1.EventServiceServer.
func (*Handler) ListEventTypes(context.Context, *eventsv1.ListEventTypesRequest) (*eventsv1.ListEventTypesResponse, error) {
	panic("unimplemented")
}

// ListSubscriptions implements eventsv1.EventServiceServer.
func (*Handler) ListSubscriptions(context.Context, *eventsv1.ListSubscriptionsRequest) (*eventsv1.ListSubscriptionsResponse, error) {
	panic("unimplemented")
}
