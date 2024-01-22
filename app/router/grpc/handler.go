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

// CreateNotificationTypes implements eventsv1.EventServiceServer.
func (h *Handler) CreateNotificationTypes(ctx context.Context, req *eventsv1.CreateNotificationTypesRequest) (*eventsv1.CreateNotificationTypesResponse, error) {
	res, err := h.Repository.CreateNotificationTypes(ctx, NotificationTypesFromProto(req.GetData()))
	if err != nil {
		return nil, err
	}
	return &eventsv1.CreateNotificationTypesResponse{
		Data: NewNotificationTypesToProto(res),
	}, nil
}

// CreateSubscriptions implements eventsv1.EventServiceServer.
func (h *Handler) CreateSubscriptions(ctx context.Context, req *eventsv1.CreateSubscriptionsRequest) (*eventsv1.CreateSubscriptionsResponse, error) {
	res, err := h.Repository.CreateSubscriptions(ctx, SubscriptionsFromProto(req.GetData()))
	if err != nil {
		return nil, err
	}
	return &eventsv1.CreateSubscriptionsResponse{
		Data: NewSubscriptionsToProto(res), //TODO: add errors
	}, nil
}

// DeleteEventTypes implements eventsv1.EventServiceServer.
func (h *Handler) DeleteEventTypes(ctx context.Context, req *eventsv1.DeleteEventTypesRequest) (*eventsv1.DeleteEventTypesResponse, error) {
	_, err := h.Repository.DeleteEventTypes(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return &eventsv1.DeleteEventTypesResponse{}, nil
}

// DeleteNotificationTypes implements eventsv1.EventServiceServer.
func (h *Handler) DeleteNotificationTypes(ctx context.Context, req *eventsv1.DeleteNotificationTypesRequest) (*eventsv1.DeleteNotificationTypesResponse, error) {
	_, err := h.Repository.DeleteEventTypes(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return &eventsv1.DeleteNotificationTypesResponse{}, nil
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

// GetNotificationTypes implements eventsv1.EventServiceServer.
func (h *Handler) GetNotificationTypes(ctx context.Context, req *eventsv1.GetNotificationTypesRequest) (*eventsv1.GetNotificationTypesResponse, error) {
	res, err := h.Repository.GetNotificationTypes(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return &eventsv1.GetNotificationTypesResponse{
		Data: GotNotificationTypesToProto(res),
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

// ListNotificationTypes implements eventsv1.EventServiceServer.
func (h *Handler) ListNotificationTypes(ctx context.Context, req *eventsv1.ListNotificationTypesRequest) (*eventsv1.ListNotificationTypesResponse, error) {
	res, err := h.Repository.ListNotificationTypes(ctx, Cursor(req.GetPage()), int(lo.FromPtr(req.GetPage()).Limit))
	if err != nil {
		return nil, err
	}
	return &eventsv1.ListNotificationTypesResponse{
		Data: NotificationTypeListToProto(res),
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
