package routergrpc

import (
	"context"
	"svix-poc/app/router"
	eventsv1 "svix-poc/app/router/grpc/v1"
	schemav1 "svix-poc/lib/schema/grpc/v1"

	"github.com/samber/lo"
)

type Handler struct {
	router.Dependencies
}

var _ eventsv1.EventServiceServer = &Handler{}
var _ schemav1.SchemaServiceServer = &Handler{}

// CreateEventTypes implements eventsv1.EventServiceServer.
func (h *Handler) CreateEventTypes(ctx context.Context, req *eventsv1.CreateEventTypesRequest) (*eventsv1.CreateEventTypesResponse, error) {
	res, err := h.Repository.CreateEventTypes(ctx, EventTypesFromProto(req.GetData()))
	if err != nil {
		return nil, err
	}
	return &eventsv1.CreateEventTypesResponse{
		Data: NewEventTypesToProto(res),
	}, nil
}

// CreateSubscriptions implements eventsv1.EventServiceServer.
func (h *Handler) CreateSubscriptions(ctx context.Context, req *eventsv1.CreateSubscriptionsRequest) (*eventsv1.CreateSubscriptionsResponse, error) {
	res, err := h.Repository.CreateSubscriptions(ctx, SubscriptionsFromProto(req.GetData()))
	if err != nil {
		return nil, err
	}
	return &eventsv1.CreateSubscriptionsResponse{
		Data: NewSubscriptionsToProto(res),
	}, nil
}

// GetEventTypes implements eventsv1.EventServiceServer.
func (h *Handler) GetEventTypes(ctx context.Context, req *eventsv1.GetEventTypesRequest) (*eventsv1.GetEventTypesResponse, error) {
	res, err := h.Repository.GetEventTypes(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return &eventsv1.GetEventTypesResponse{
		Data: GotEventTypesToProto(res),
	}, nil
}

// GetProviders implements eventsv1.EventServiceServer.
func (h *Handler) GetProviders(ctx context.Context, req *eventsv1.GetProvidersRequest) (*eventsv1.GetProvidersResponse, error) {
	res, err := h.Repository.GetProviders(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return &eventsv1.GetProvidersResponse{
		Data: GotProvidersToProto(res),
	}, nil
}

// GetSubscriptions implements eventsv1.EventServiceServer.
func (h *Handler) GetSubscriptions(ctx context.Context, req *eventsv1.GetSubscriptionsRequest) (*eventsv1.GetSubscriptionsResponse, error) {
	res, err := h.Repository.GetSubscriptions(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return &eventsv1.GetSubscriptionsResponse{
		Data: GotSubscriptionsToProto(res),
	}, nil
}

// ListEventTypes implements eventsv1.EventServiceServer.
func (h *Handler) ListEventTypes(ctx context.Context, req *eventsv1.ListEventTypesRequest) (*eventsv1.ListEventTypesResponse, error) {
	res, err := h.Repository.ListEventTypes(ctx, Cursor(req.GetPage()), int(lo.FromPtr(req.GetPage()).Limit))
	if err != nil {
		return nil, err
	}
	return &eventsv1.ListEventTypesResponse{
		Data: EventTypeListToProto(res),
		Page: &eventsv1.PageResponse{},
	}, nil
}

// ListProviders implements eventsv1.EventServiceServer.
func (h *Handler) ListProviders(ctx context.Context, req *eventsv1.ListProvidersRequest) (*eventsv1.ListProvidersResponse, error) {
	res, err := h.Repository.ListProviders(ctx, Cursor(req.GetPage()), int(lo.FromPtr(req.GetPage()).Limit))
	if err != nil {
		return nil, err
	}
	return &eventsv1.ListProvidersResponse{
		Data: ProviderListToProto(res),
		Page: &eventsv1.PageResponse{},
	}, nil
}

// ListSubscriptions implements eventsv1.EventServiceServer.
func (h *Handler) ListSubscriptions(ctx context.Context, req *eventsv1.ListSubscriptionsRequest) (*eventsv1.ListSubscriptionsResponse, error) {
	res, err := h.Repository.ListSubscriptions(ctx, Cursor(req.GetPage()), int(lo.FromPtr(req.GetPage()).Limit))
	if err != nil {
		return nil, err
	}
	return &eventsv1.ListSubscriptionsResponse{
		Data: SubscriptionListToProto(res),
		Page: &eventsv1.PageResponse{},
	}, nil
}
