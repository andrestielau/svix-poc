package webhookgrpc

import (
	"context"
	"errors"
	"strconv"
	"svix-poc/app/webhook"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	schemav1 "svix-poc/lib/schema/grpc/v1"

	"github.com/samber/lo"
	svix "github.com/svix/svix-webhooks/go"
)

// TODO: wrap handler with a provider instead of havingt server inside
type Handler struct {
	webhook.Dependencies
}

var _ webhooksv1.WebHookServiceServer = &Handler{}
var _ schemav1.SchemaServiceServer = &Handler{}

// CreateApps implements webhooksv1.WebHookServiceServer.
func (h *Handler) CreateApps(ctx context.Context, req *webhooksv1.CreateAppsRequest) (*webhooksv1.CreateAppsResponse, error) {
	ret, errs := Batch(req.Data, func(in *webhooksv1.App, i int) (*webhooksv1.App, *webhooksv1.Error) {
		res, err := h.Application.Create(ctx, AppFromProto(in))
		if err != nil {
			return nil, Err(err, i)
		}
		return AppToProto(res), nil
	})
	return &webhooksv1.CreateAppsResponse{Data: ret, Errors: errs}, nil
}

// CreateEndpoints implements webhooksv1.WebHookServiceServer.
func (h *Handler) CreateEndpoints(ctx context.Context, req *webhooksv1.CreateEndpointsRequest) (*webhooksv1.CreateEndpointsResponse, error) {
	ret, errs := Batch(req.Data, func(in *webhooksv1.Endpoint, i int) (*webhooksv1.Endpoint, *webhooksv1.Error) {
		res, err := h.Endpoint.Create(ctx, "", EndpointFromProto(in))
		if err != nil {
			return nil, Err(err, i)
		}
		return EndpointToProto(res), nil
	})
	return &webhooksv1.CreateEndpointsResponse{Data: ret, Errors: errs}, nil
}

// CreateEventTypes implements webhooksv1.WebHookServiceServer.
func (h *Handler) CreateEventTypes(ctx context.Context, req *webhooksv1.CreateEventTypesRequest) (*webhooksv1.CreateEventTypesResponse, error) {
	ret, errs := Batch(req.Data, func(in *webhooksv1.EventType, i int) (*webhooksv1.EventType, *webhooksv1.Error) {
		res, err := h.EventType.Create(ctx, EventTypeFromProto(in))
		if err != nil {
			return nil, Err(err, i)
		}
		return EventTypeToProto(res), nil
	})
	return &webhooksv1.CreateEventTypesResponse{Data: ret, Errors: errs}, nil
}

// CreateMessages implements webhooksv1.WebHookServiceServer.
func (h *Handler) CreateMessages(ctx context.Context, req *webhooksv1.CreateMessagesRequest) (*webhooksv1.CreateMessagesResponse, error) {
	ret, errs := Batch(req.Data, func(in *webhooksv1.Message, i int) (*webhooksv1.Message, *webhooksv1.Error) {
		res, err := h.Message.Create(ctx, "", MessageFromProto(in))
		if err != nil {
			return nil, Err(err, i)
		}
		return MessageToProto(res), nil
	})
	return &webhooksv1.CreateMessagesResponse{Data: ret, Errors: errs}, nil
}

// GetApps implements webhooksv1.WebHookServiceServer.
func (h *Handler) GetApps(ctx context.Context, req *webhooksv1.GetAppsRequest) (*webhooksv1.GetAppsResponse, error) {
	ret, errs := BatchMap(req.Ids, func(in string, i int) (*webhooksv1.App, *webhooksv1.Error) {
		res, err := h.Application.Get(ctx, in)
		if err != nil {
			return nil, Err(err, i)
		}
		return AppToProto(res), nil
	})
	return &webhooksv1.GetAppsResponse{Data: ret, Errors: errs}, nil
}

// GetEndpoints implements webhooksv1.WebHookServiceServer.
func (h *Handler) GetEndpoints(ctx context.Context, req *webhooksv1.GetEndpointsRequest) (*webhooksv1.GetEndpointsResponse, error) {
	ret, errs := BatchMap(req.Ids, func(in string, i int) (*webhooksv1.Endpoint, *webhooksv1.Error) {
		res, err := h.Endpoint.Get(ctx, req.TenantId, in)
		if err != nil {
			return nil, Err(err, i)
		}
		return EndpointToProto(res), nil
	})
	return &webhooksv1.GetEndpointsResponse{Data: ret, Errors: errs}, nil
}

// GetEventTypes implements webhooksv1.WebHookServiceServer.
func (h *Handler) GetEventTypes(ctx context.Context, req *webhooksv1.GetEventTypesRequest) (*webhooksv1.GetEventTypesResponse, error) {
	ret, errs := BatchMap(req.Ids, func(in string, i int) (*webhooksv1.EventType, *webhooksv1.Error) {
		res, err := h.EventType.Get(ctx, in)
		if err != nil {
			return nil, Err(err, i)
		}
		return EventTypeToProto(res), nil
	})
	return &webhooksv1.GetEventTypesResponse{Data: ret, Errors: errs}, nil
}

// GetMessages implements webhooksv1.WebHookServiceServer.
func (h *Handler) GetMessages(ctx context.Context, req *webhooksv1.GetMessagesRequest) (*webhooksv1.GetMessagesResponse, error) {
	ret, errs := BatchMap(req.Ids, func(in string, i int) (*webhooksv1.Message, *webhooksv1.Error) {
		res, err := h.Message.Get(ctx, req.TenantId, in)
		if err != nil {
			return nil, Err(err, i)
		}
		return MessageToProto(res), nil
	})
	return &webhooksv1.GetMessagesResponse{Data: ret, Errors: errs}, nil
}

// ListApps implements webhooksv1.WebHookServiceServer.
func (h *Handler) ListApps(ctx context.Context, req *webhooksv1.ListAppsRequest) (*webhooksv1.ListAppsResponse, error) {
	res, err := h.Application.List(ctx, &svix.ApplicationListOptions{
		Iterator: lo.EmptyableToPtr(req.Page.Cursor),
		Limit:    &req.Page.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &webhooksv1.ListAppsResponse{
		Data: lo.Map(res.Data, func(item svix.ApplicationOut, _ int) *webhooksv1.App {
			return AppToProto(&item)
		}),
		Page: &webhooksv1.PageResponse{
			Prev: res.GetPrevIterator(),
			Next: res.GetIterator(),
			Done: res.Done,
		},
	}, nil
}

// ListEndpoints implements webhooksv1.WebHookServiceServer.
func (h *Handler) ListEndpoints(ctx context.Context, req *webhooksv1.ListEndpointsRequest) (*webhooksv1.ListEndpointsResponse, error) {
	res, err := h.Endpoint.List(ctx, req.TenantId, &svix.EndpointListOptions{
		Iterator: lo.EmptyableToPtr(req.Page.Cursor),
		Limit:    &req.Page.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &webhooksv1.ListEndpointsResponse{
		Data: lo.Map(res.Data, func(item svix.EndpointOut, _ int) *webhooksv1.Endpoint {
			return EndpointToProto(&item)
		}),
		Page: &webhooksv1.PageResponse{
			Prev: res.GetPrevIterator(),
			Next: res.GetIterator(),
			Done: res.Done,
		},
	}, nil
}

// ListEventTypes implements webhooksv1.WebHookServiceServer.
func (h *Handler) ListEventTypes(ctx context.Context, req *webhooksv1.ListEventTypesRequest) (*webhooksv1.ListEventTypesResponse, error) {
	res, err := h.EventType.List(ctx, &svix.EventTypeListOptions{
		Iterator: lo.EmptyableToPtr(req.Page.Cursor),
		Limit:    &req.Page.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &webhooksv1.ListEventTypesResponse{
		Data: lo.Map(res.Data, func(item svix.EventTypeOut, _ int) *webhooksv1.EventType {
			return EventTypeToProto(&item)
		}),
		Page: &webhooksv1.PageResponse{
			Prev: res.GetPrevIterator(),
			Next: res.GetIterator(),
			Done: res.Done,
		},
	}, nil
}

// ListMessages implements webhooksv1.WebHookServiceServer.
func (h *Handler) ListMessages(ctx context.Context, req *webhooksv1.ListMessagesRequest) (*webhooksv1.ListMessagesResponse, error) {
	res, err := h.Message.List(ctx, req.TenantId, &svix.MessageListOptions{
		Iterator: lo.EmptyableToPtr(req.Page.Cursor),
		Limit:    &req.Page.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &webhooksv1.ListMessagesResponse{
		Data: lo.Map(res.Data, func(item svix.MessageOut, _ int) *webhooksv1.Message {
			return MessageToProto(&item)
		}),
		Page: &webhooksv1.PageResponse{
			Prev: res.GetPrevIterator(),
			Next: res.GetIterator(),
			Done: res.Done,
		},
	}, nil
}

type BatchFn[F, T, E any] func(F, int) (T, *E)

func BatchMap[F comparable, T, E any](in []F, fn BatchFn[F, T, E]) (ret map[F]T, errs []*E) {
	ret = map[F]T{}
	for i, v := range in {
		if res, err := fn(v, i); err != nil {
			errs = append(errs, err)
		} else {
			ret[v] = res
		}
	}
	return ret, errs
}
func Batch[F, T, E any](in []F, fn BatchFn[F, T, E]) (ret []T, errs []*E) {
	for i, v := range in {
		if res, err := fn(v, i); err != nil {
			errs = append(errs, err)
		} else {
			ret = append(ret, res)
		}
	}
	return ret, errs
}
func Err(err error, i int) *webhooksv1.Error {
	code := 500
	var serr svix.Error
	if errors.As(err, &serr) {
		code = serr.Status()
	}
	return &webhooksv1.Error{
		Code:   int32(code),
		Index:  strconv.Itoa(i),
		Reason: err.Error(),
	}
}
