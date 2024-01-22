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

function serialize_events_v1_CreateNotificationTypesRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.CreateNotificationTypesRequest)) {
    throw new Error('Expected argument of type events.v1.CreateNotificationTypesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_CreateNotificationTypesRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.CreateNotificationTypesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_CreateNotificationTypesResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.CreateNotificationTypesResponse)) {
    throw new Error('Expected argument of type events.v1.CreateNotificationTypesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_CreateNotificationTypesResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.CreateNotificationTypesResponse.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_events_v1_DeleteEventTypesRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.DeleteEventTypesRequest)) {
    throw new Error('Expected argument of type events.v1.DeleteEventTypesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_DeleteEventTypesRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.DeleteEventTypesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_DeleteEventTypesResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.DeleteEventTypesResponse)) {
    throw new Error('Expected argument of type events.v1.DeleteEventTypesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_DeleteEventTypesResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.DeleteEventTypesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_DeleteNotificationTypesRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.DeleteNotificationTypesRequest)) {
    throw new Error('Expected argument of type events.v1.DeleteNotificationTypesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_DeleteNotificationTypesRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.DeleteNotificationTypesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_DeleteNotificationTypesResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.DeleteNotificationTypesResponse)) {
    throw new Error('Expected argument of type events.v1.DeleteNotificationTypesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_DeleteNotificationTypesResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.DeleteNotificationTypesResponse.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_events_v1_GetNotificationTypesRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.GetNotificationTypesRequest)) {
    throw new Error('Expected argument of type events.v1.GetNotificationTypesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_GetNotificationTypesRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.GetNotificationTypesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_GetNotificationTypesResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.GetNotificationTypesResponse)) {
    throw new Error('Expected argument of type events.v1.GetNotificationTypesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_GetNotificationTypesResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.GetNotificationTypesResponse.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_events_v1_ListNotificationTypesRequest(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.ListNotificationTypesRequest)) {
    throw new Error('Expected argument of type events.v1.ListNotificationTypesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_ListNotificationTypesRequest(buffer_arg) {
  return app_router_grpc_v1_event_pb.ListNotificationTypesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_events_v1_ListNotificationTypesResponse(arg) {
  if (!(arg instanceof app_router_grpc_v1_event_pb.ListNotificationTypesResponse)) {
    throw new Error('Expected argument of type events.v1.ListNotificationTypesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_events_v1_ListNotificationTypesResponse(buffer_arg) {
  return app_router_grpc_v1_event_pb.ListNotificationTypesResponse.deserializeBinary(new Uint8Array(buffer_arg));
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
  deleteEventTypes: {
    path: '/events.v1.EventService/DeleteEventTypes',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.DeleteEventTypesRequest,
    responseType: app_router_grpc_v1_event_pb.DeleteEventTypesResponse,
    requestSerialize: serialize_events_v1_DeleteEventTypesRequest,
    requestDeserialize: deserialize_events_v1_DeleteEventTypesRequest,
    responseSerialize: serialize_events_v1_DeleteEventTypesResponse,
    responseDeserialize: deserialize_events_v1_DeleteEventTypesResponse,
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
  getNotificationTypes: {
    path: '/events.v1.EventService/GetNotificationTypes',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.GetNotificationTypesRequest,
    responseType: app_router_grpc_v1_event_pb.GetNotificationTypesResponse,
    requestSerialize: serialize_events_v1_GetNotificationTypesRequest,
    requestDeserialize: deserialize_events_v1_GetNotificationTypesRequest,
    responseSerialize: serialize_events_v1_GetNotificationTypesResponse,
    responseDeserialize: deserialize_events_v1_GetNotificationTypesResponse,
  },
  listNotificationTypes: {
    path: '/events.v1.EventService/ListNotificationTypes',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.ListNotificationTypesRequest,
    responseType: app_router_grpc_v1_event_pb.ListNotificationTypesResponse,
    requestSerialize: serialize_events_v1_ListNotificationTypesRequest,
    requestDeserialize: deserialize_events_v1_ListNotificationTypesRequest,
    responseSerialize: serialize_events_v1_ListNotificationTypesResponse,
    responseDeserialize: deserialize_events_v1_ListNotificationTypesResponse,
  },
  createNotificationTypes: {
    path: '/events.v1.EventService/CreateNotificationTypes',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.CreateNotificationTypesRequest,
    responseType: app_router_grpc_v1_event_pb.CreateNotificationTypesResponse,
    requestSerialize: serialize_events_v1_CreateNotificationTypesRequest,
    requestDeserialize: deserialize_events_v1_CreateNotificationTypesRequest,
    responseSerialize: serialize_events_v1_CreateNotificationTypesResponse,
    responseDeserialize: deserialize_events_v1_CreateNotificationTypesResponse,
  },
  deleteNotificationTypes: {
    path: '/events.v1.EventService/DeleteNotificationTypes',
    requestStream: false,
    responseStream: false,
    requestType: app_router_grpc_v1_event_pb.DeleteNotificationTypesRequest,
    responseType: app_router_grpc_v1_event_pb.DeleteNotificationTypesResponse,
    requestSerialize: serialize_events_v1_DeleteNotificationTypesRequest,
    requestDeserialize: deserialize_events_v1_DeleteNotificationTypesRequest,
    responseSerialize: serialize_events_v1_DeleteNotificationTypesResponse,
    responseDeserialize: deserialize_events_v1_DeleteNotificationTypesResponse,
  },
};

exports.EventServiceClient = grpc.makeGenericClientConstructor(EventServiceService);
