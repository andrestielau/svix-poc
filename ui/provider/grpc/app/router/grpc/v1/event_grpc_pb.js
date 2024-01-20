// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var app_router_grpc_v1_event_pb = require('../../../../app/router/grpc/v1/event_pb.js');

function serialize_events_v1_CreateEventTypesRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.CreateEventTypesRequest)) {
    throw new Error('Expected argument of type events.v1.CreateEventTypesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_CreateEventTypesRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.CreateEventTypesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_CreateEventTypesResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.CreateEventTypesResponse)) {
    throw new Error('Expected argument of type events.v1.CreateEventTypesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_CreateEventTypesResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.CreateEventTypesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_CreateSubscriptionsRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.CreateSubscriptionsRequest)) {
    throw new Error('Expected argument of type events.v1.CreateSubscriptionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_CreateSubscriptionsRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.CreateSubscriptionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_CreateSubscriptionsResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.CreateSubscriptionsResponse)) {
    throw new Error('Expected argument of type events.v1.CreateSubscriptionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_CreateSubscriptionsResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.CreateSubscriptionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_GetEventTypesRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.GetEventTypesRequest)) {
    throw new Error('Expected argument of type events.v1.GetEventTypesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_GetEventTypesRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.GetEventTypesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_GetEventTypesResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.GetEventTypesResponse)) {
    throw new Error('Expected argument of type events.v1.GetEventTypesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_GetEventTypesResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.GetEventTypesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_GetProvidersRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.GetProvidersRequest)) {
    throw new Error('Expected argument of type events.v1.GetProvidersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_GetProvidersRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.GetProvidersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_GetProvidersResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.GetProvidersResponse)) {
    throw new Error('Expected argument of type events.v1.GetProvidersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_GetProvidersResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.GetProvidersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_GetSubscriptionsRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.GetSubscriptionsRequest)) {
    throw new Error('Expected argument of type events.v1.GetSubscriptionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_GetSubscriptionsRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.GetSubscriptionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_GetSubscriptionsResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.GetSubscriptionsResponse)) {
    throw new Error('Expected argument of type events.v1.GetSubscriptionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_GetSubscriptionsResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.GetSubscriptionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_ListEventTypesRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.ListEventTypesRequest)) {
    throw new Error('Expected argument of type events.v1.ListEventTypesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_ListEventTypesRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.ListEventTypesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_ListEventTypesResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.ListEventTypesResponse)) {
    throw new Error('Expected argument of type events.v1.ListEventTypesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_ListEventTypesResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.ListEventTypesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_ListProvidersRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.ListProvidersRequest)) {
    throw new Error('Expected argument of type events.v1.ListProvidersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_ListProvidersRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.ListProvidersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_ListProvidersResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.ListProvidersResponse)) {
    throw new Error('Expected argument of type events.v1.ListProvidersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_ListProvidersResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.ListProvidersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_ListSubscriptionsRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.ListSubscriptionsRequest)) {
    throw new Error('Expected argument of type events.v1.ListSubscriptionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_ListSubscriptionsRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.ListSubscriptionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_ListSubscriptionsResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.ListSubscriptionsResponse)) {
    throw new Error('Expected argument of type events.v1.ListSubscriptionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_ListSubscriptionsResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.ListSubscriptionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var EventServiceService = exports.EventServiceService = {
  getProviders: {
    path: '/events.v1.EventService/GetProviders',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.GetProvidersRequest,
    responseType: app_router_grpc_v1_event_pb.GetProvidersResponse,
    requestSerialize: serialize_events_v1_GetProvidersRequest,
    requestDeserialize: deserialize_events_v1_GetProvidersRequest,
    responseSerialize: serialize_events_v1_GetProvidersResponse,
    responseDeserialize: deserialize_events_v1_GetProvidersResponse,
  },
  listProviders: {
    path: '/events.v1.EventService/ListProviders',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.ListProvidersRequest,
    responseType: app_router_grpc_v1_event_pb.ListProvidersResponse,
    requestSerialize: serialize_events_v1_ListProvidersRequest,
    requestDeserialize: deserialize_events_v1_ListProvidersRequest,
    responseSerialize: serialize_events_v1_ListProvidersResponse,
    responseDeserialize: deserialize_events_v1_ListProvidersResponse,
  },
  getEventTypes: {
    path: '/events.v1.EventService/GetEventTypes',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.GetEventTypesRequest,
    responseType: app_router_grpc_v1_event_pb.GetEventTypesResponse,
    requestSerialize: serialize_events_v1_GetEventTypesRequest,
    requestDeserialize: deserialize_events_v1_GetEventTypesRequest,
    responseSerialize: serialize_events_v1_GetEventTypesResponse,
    responseDeserialize: deserialize_events_v1_GetEventTypesResponse,
  },
  listEventTypes: {
    path: '/events.v1.EventService/ListEventTypes',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.ListEventTypesRequest,
    responseType: app_router_grpc_v1_event_pb.ListEventTypesResponse,
    requestSerialize: serialize_events_v1_ListEventTypesRequest,
    requestDeserialize: deserialize_events_v1_ListEventTypesRequest,
    responseSerialize: serialize_events_v1_ListEventTypesResponse,
    responseDeserialize: deserialize_events_v1_ListEventTypesResponse,
  },
  createEventTypes: {
    path: '/events.v1.EventService/CreateEventTypes',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.CreateEventTypesRequest,
    responseType: app_router_grpc_v1_event_pb.CreateEventTypesResponse,
    requestSerialize: serialize_events_v1_CreateEventTypesRequest,
    requestDeserialize: deserialize_events_v1_CreateEventTypesRequest,
    responseSerialize: serialize_events_v1_CreateEventTypesResponse,
    responseDeserialize: deserialize_events_v1_CreateEventTypesResponse,
  },
  getSubscriptions: {
    path: '/events.v1.EventService/GetSubscriptions',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.GetSubscriptionsRequest,
    responseType: app_router_grpc_v1_event_pb.GetSubscriptionsResponse,
    requestSerialize: serialize_events_v1_GetSubscriptionsRequest,
    requestDeserialize: deserialize_events_v1_GetSubscriptionsRequest,
    responseSerialize: serialize_events_v1_GetSubscriptionsResponse,
    responseDeserialize: deserialize_events_v1_GetSubscriptionsResponse,
  },
  listSubscriptions: {
    path: '/events.v1.EventService/ListSubscriptions',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.ListSubscriptionsRequest,
    responseType: app_router_grpc_v1_event_pb.ListSubscriptionsResponse,
    requestSerialize: serialize_events_v1_ListSubscriptionsRequest,
    requestDeserialize: deserialize_events_v1_ListSubscriptionsRequest,
    responseSerialize: serialize_events_v1_ListSubscriptionsResponse,
    responseDeserialize: deserialize_events_v1_ListSubscriptionsResponse,
  },
  createSubscriptions: {
    path: '/events.v1.EventService/CreateSubscriptions',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.CreateSubscriptionsRequest,
    responseType: app_router_grpc_v1_event_pb.CreateSubscriptionsResponse,
    requestSerialize: serialize_events_v1_CreateSubscriptionsRequest,
    requestDeserialize: deserialize_events_v1_CreateSubscriptionsRequest,
    responseSerialize: serialize_events_v1_CreateSubscriptionsResponse,
    responseDeserialize: deserialize_events_v1_CreateSubscriptionsResponse,
  },
};

exports.EventServiceClient = grpc.makeGenericClientConstructor(EventServiceService);
