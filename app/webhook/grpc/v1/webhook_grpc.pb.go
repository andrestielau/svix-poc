// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: app/webhook/grpc/v1/webhook.proto

package webhooksv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	WebHookService_GetApps_FullMethodName          = "/webhooks.v1.WebHookService/GetApps"
	WebHookService_ListApps_FullMethodName         = "/webhooks.v1.WebHookService/ListApps"
	WebHookService_CreateApps_FullMethodName       = "/webhooks.v1.WebHookService/CreateApps"
	WebHookService_GetEndpoints_FullMethodName     = "/webhooks.v1.WebHookService/GetEndpoints"
	WebHookService_ListEndpoints_FullMethodName    = "/webhooks.v1.WebHookService/ListEndpoints"
	WebHookService_CreateEndpoints_FullMethodName  = "/webhooks.v1.WebHookService/CreateEndpoints"
	WebHookService_GetMessages_FullMethodName      = "/webhooks.v1.WebHookService/GetMessages"
	WebHookService_ListMessages_FullMethodName     = "/webhooks.v1.WebHookService/ListMessages"
	WebHookService_CreateMessages_FullMethodName   = "/webhooks.v1.WebHookService/CreateMessages"
	WebHookService_GetEventTypes_FullMethodName    = "/webhooks.v1.WebHookService/GetEventTypes"
	WebHookService_ListEventTypes_FullMethodName   = "/webhooks.v1.WebHookService/ListEventTypes"
	WebHookService_CreateEventTypes_FullMethodName = "/webhooks.v1.WebHookService/CreateEventTypes"
)

// WebHookServiceClient is the client API for WebHookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebHookServiceClient interface {
	GetApps(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppsResponse, error)
	ListApps(ctx context.Context, in *ListAppsRequest, opts ...grpc.CallOption) (*ListAppsResponse, error)
	CreateApps(ctx context.Context, in *CreateAppsRequest, opts ...grpc.CallOption) (*CreateAppsResponse, error)
	GetEndpoints(ctx context.Context, in *GetEndpointsRequest, opts ...grpc.CallOption) (*GetEndpointsResponse, error)
	ListEndpoints(ctx context.Context, in *ListEndpointsRequest, opts ...grpc.CallOption) (*ListEndpointsResponse, error)
	CreateEndpoints(ctx context.Context, in *CreateEndpointsRequest, opts ...grpc.CallOption) (*CreateEndpointsResponse, error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error)
	CreateMessages(ctx context.Context, in *CreateMessagesRequest, opts ...grpc.CallOption) (*CreateMessagesResponse, error)
	GetEventTypes(ctx context.Context, in *GetEventTypesRequest, opts ...grpc.CallOption) (*GetEventTypesResponse, error)
	ListEventTypes(ctx context.Context, in *ListEventTypesRequest, opts ...grpc.CallOption) (*ListEventTypesResponse, error)
	CreateEventTypes(ctx context.Context, in *CreateEventTypesRequest, opts ...grpc.CallOption) (*CreateEventTypesResponse, error)
}

type webHookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebHookServiceClient(cc grpc.ClientConnInterface) WebHookServiceClient {
	return &webHookServiceClient{cc}
}

func (c *webHookServiceClient) GetApps(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppsResponse, error) {
	out := new(GetAppsResponse)
	err := c.cc.Invoke(ctx, WebHookService_GetApps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHookServiceClient) ListApps(ctx context.Context, in *ListAppsRequest, opts ...grpc.CallOption) (*ListAppsResponse, error) {
	out := new(ListAppsResponse)
	err := c.cc.Invoke(ctx, WebHookService_ListApps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHookServiceClient) CreateApps(ctx context.Context, in *CreateAppsRequest, opts ...grpc.CallOption) (*CreateAppsResponse, error) {
	out := new(CreateAppsResponse)
	err := c.cc.Invoke(ctx, WebHookService_CreateApps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHookServiceClient) GetEndpoints(ctx context.Context, in *GetEndpointsRequest, opts ...grpc.CallOption) (*GetEndpointsResponse, error) {
	out := new(GetEndpointsResponse)
	err := c.cc.Invoke(ctx, WebHookService_GetEndpoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHookServiceClient) ListEndpoints(ctx context.Context, in *ListEndpointsRequest, opts ...grpc.CallOption) (*ListEndpointsResponse, error) {
	out := new(ListEndpointsResponse)
	err := c.cc.Invoke(ctx, WebHookService_ListEndpoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHookServiceClient) CreateEndpoints(ctx context.Context, in *CreateEndpointsRequest, opts ...grpc.CallOption) (*CreateEndpointsResponse, error) {
	out := new(CreateEndpointsResponse)
	err := c.cc.Invoke(ctx, WebHookService_CreateEndpoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHookServiceClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, WebHookService_GetMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHookServiceClient) ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error) {
	out := new(ListMessagesResponse)
	err := c.cc.Invoke(ctx, WebHookService_ListMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHookServiceClient) CreateMessages(ctx context.Context, in *CreateMessagesRequest, opts ...grpc.CallOption) (*CreateMessagesResponse, error) {
	out := new(CreateMessagesResponse)
	err := c.cc.Invoke(ctx, WebHookService_CreateMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHookServiceClient) GetEventTypes(ctx context.Context, in *GetEventTypesRequest, opts ...grpc.CallOption) (*GetEventTypesResponse, error) {
	out := new(GetEventTypesResponse)
	err := c.cc.Invoke(ctx, WebHookService_GetEventTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHookServiceClient) ListEventTypes(ctx context.Context, in *ListEventTypesRequest, opts ...grpc.CallOption) (*ListEventTypesResponse, error) {
	out := new(ListEventTypesResponse)
	err := c.cc.Invoke(ctx, WebHookService_ListEventTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHookServiceClient) CreateEventTypes(ctx context.Context, in *CreateEventTypesRequest, opts ...grpc.CallOption) (*CreateEventTypesResponse, error) {
	out := new(CreateEventTypesResponse)
	err := c.cc.Invoke(ctx, WebHookService_CreateEventTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebHookServiceServer is the server API for WebHookService service.
// All implementations should embed UnimplementedWebHookServiceServer
// for forward compatibility
type WebHookServiceServer interface {
	GetApps(context.Context, *GetAppsRequest) (*GetAppsResponse, error)
	ListApps(context.Context, *ListAppsRequest) (*ListAppsResponse, error)
	CreateApps(context.Context, *CreateAppsRequest) (*CreateAppsResponse, error)
	GetEndpoints(context.Context, *GetEndpointsRequest) (*GetEndpointsResponse, error)
	ListEndpoints(context.Context, *ListEndpointsRequest) (*ListEndpointsResponse, error)
	CreateEndpoints(context.Context, *CreateEndpointsRequest) (*CreateEndpointsResponse, error)
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	ListMessages(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error)
	CreateMessages(context.Context, *CreateMessagesRequest) (*CreateMessagesResponse, error)
	GetEventTypes(context.Context, *GetEventTypesRequest) (*GetEventTypesResponse, error)
	ListEventTypes(context.Context, *ListEventTypesRequest) (*ListEventTypesResponse, error)
	CreateEventTypes(context.Context, *CreateEventTypesRequest) (*CreateEventTypesResponse, error)
}

// UnimplementedWebHookServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWebHookServiceServer struct {
}

func (UnimplementedWebHookServiceServer) GetApps(context.Context, *GetAppsRequest) (*GetAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApps not implemented")
}
func (UnimplementedWebHookServiceServer) ListApps(context.Context, *ListAppsRequest) (*ListAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApps not implemented")
}
func (UnimplementedWebHookServiceServer) CreateApps(context.Context, *CreateAppsRequest) (*CreateAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApps not implemented")
}
func (UnimplementedWebHookServiceServer) GetEndpoints(context.Context, *GetEndpointsRequest) (*GetEndpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEndpoints not implemented")
}
func (UnimplementedWebHookServiceServer) ListEndpoints(context.Context, *ListEndpointsRequest) (*ListEndpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEndpoints not implemented")
}
func (UnimplementedWebHookServiceServer) CreateEndpoints(context.Context, *CreateEndpointsRequest) (*CreateEndpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEndpoints not implemented")
}
func (UnimplementedWebHookServiceServer) GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedWebHookServiceServer) ListMessages(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessages not implemented")
}
func (UnimplementedWebHookServiceServer) CreateMessages(context.Context, *CreateMessagesRequest) (*CreateMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessages not implemented")
}
func (UnimplementedWebHookServiceServer) GetEventTypes(context.Context, *GetEventTypesRequest) (*GetEventTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventTypes not implemented")
}
func (UnimplementedWebHookServiceServer) ListEventTypes(context.Context, *ListEventTypesRequest) (*ListEventTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEventTypes not implemented")
}
func (UnimplementedWebHookServiceServer) CreateEventTypes(context.Context, *CreateEventTypesRequest) (*CreateEventTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEventTypes not implemented")
}

// UnsafeWebHookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebHookServiceServer will
// result in compilation errors.
type UnsafeWebHookServiceServer interface {
	mustEmbedUnimplementedWebHookServiceServer()
}

func RegisterWebHookServiceServer(s grpc.ServiceRegistrar, srv WebHookServiceServer) {
	s.RegisterService(&WebHookService_ServiceDesc, srv)
}

func _WebHookService_GetApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHookServiceServer).GetApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebHookService_GetApps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHookServiceServer).GetApps(ctx, req.(*GetAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHookService_ListApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHookServiceServer).ListApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebHookService_ListApps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHookServiceServer).ListApps(ctx, req.(*ListAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHookService_CreateApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHookServiceServer).CreateApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebHookService_CreateApps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHookServiceServer).CreateApps(ctx, req.(*CreateAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHookService_GetEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHookServiceServer).GetEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebHookService_GetEndpoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHookServiceServer).GetEndpoints(ctx, req.(*GetEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHookService_ListEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHookServiceServer).ListEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebHookService_ListEndpoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHookServiceServer).ListEndpoints(ctx, req.(*ListEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHookService_CreateEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHookServiceServer).CreateEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebHookService_CreateEndpoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHookServiceServer).CreateEndpoints(ctx, req.(*CreateEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHookService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHookServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebHookService_GetMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHookServiceServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHookService_ListMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHookServiceServer).ListMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebHookService_ListMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHookServiceServer).ListMessages(ctx, req.(*ListMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHookService_CreateMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHookServiceServer).CreateMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebHookService_CreateMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHookServiceServer).CreateMessages(ctx, req.(*CreateMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHookService_GetEventTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHookServiceServer).GetEventTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebHookService_GetEventTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHookServiceServer).GetEventTypes(ctx, req.(*GetEventTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHookService_ListEventTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHookServiceServer).ListEventTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebHookService_ListEventTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHookServiceServer).ListEventTypes(ctx, req.(*ListEventTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHookService_CreateEventTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHookServiceServer).CreateEventTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebHookService_CreateEventTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHookServiceServer).CreateEventTypes(ctx, req.(*CreateEventTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WebHookService_ServiceDesc is the grpc.ServiceDesc for WebHookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebHookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "webhooks.v1.WebHookService",
	HandlerType: (*WebHookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApps",
			Handler:    _WebHookService_GetApps_Handler,
		},
		{
			MethodName: "ListApps",
			Handler:    _WebHookService_ListApps_Handler,
		},
		{
			MethodName: "CreateApps",
			Handler:    _WebHookService_CreateApps_Handler,
		},
		{
			MethodName: "GetEndpoints",
			Handler:    _WebHookService_GetEndpoints_Handler,
		},
		{
			MethodName: "ListEndpoints",
			Handler:    _WebHookService_ListEndpoints_Handler,
		},
		{
			MethodName: "CreateEndpoints",
			Handler:    _WebHookService_CreateEndpoints_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _WebHookService_GetMessages_Handler,
		},
		{
			MethodName: "ListMessages",
			Handler:    _WebHookService_ListMessages_Handler,
		},
		{
			MethodName: "CreateMessages",
			Handler:    _WebHookService_CreateMessages_Handler,
		},
		{
			MethodName: "GetEventTypes",
			Handler:    _WebHookService_GetEventTypes_Handler,
		},
		{
			MethodName: "ListEventTypes",
			Handler:    _WebHookService_ListEventTypes_Handler,
		},
		{
			MethodName: "CreateEventTypes",
			Handler:    _WebHookService_CreateEventTypes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/webhook/grpc/v1/webhook.proto",
}
