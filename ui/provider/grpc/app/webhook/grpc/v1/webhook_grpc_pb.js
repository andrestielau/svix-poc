// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var app_webhook_grpc_v1_webhook_pb = require('../../../../app/webhook/grpc/v1/webhook_pb.js');

function serialize_webhooks_v1_CreateAppsRequest(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.CreateAppsRequest)) {
    throw new Error('Expected argument of type webhooks.v1.CreateAppsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_CreateAppsRequest(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.CreateAppsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_CreateAppsResponse(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.CreateAppsResponse)) {
    throw new Error('Expected argument of type webhooks.v1.CreateAppsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_CreateAppsResponse(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.CreateAppsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_CreateEndpointsRequest(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.CreateEndpointsRequest)) {
    throw new Error('Expected argument of type webhooks.v1.CreateEndpointsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_CreateEndpointsRequest(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.CreateEndpointsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_CreateEndpointsResponse(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.CreateEndpointsResponse)) {
    throw new Error('Expected argument of type webhooks.v1.CreateEndpointsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_CreateEndpointsResponse(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.CreateEndpointsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_CreateEventTypesRequest(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.CreateEventTypesRequest)) {
    throw new Error('Expected argument of type webhooks.v1.CreateEventTypesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_CreateEventTypesRequest(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.CreateEventTypesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_CreateEventTypesResponse(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.CreateEventTypesResponse)) {
    throw new Error('Expected argument of type webhooks.v1.CreateEventTypesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_CreateEventTypesResponse(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.CreateEventTypesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_CreateMessagesRequest(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.CreateMessagesRequest)) {
    throw new Error('Expected argument of type webhooks.v1.CreateMessagesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_CreateMessagesRequest(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.CreateMessagesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_CreateMessagesResponse(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.CreateMessagesResponse)) {
    throw new Error('Expected argument of type webhooks.v1.CreateMessagesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_CreateMessagesResponse(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.CreateMessagesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_GetAppsRequest(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.GetAppsRequest)) {
    throw new Error('Expected argument of type webhooks.v1.GetAppsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_GetAppsRequest(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.GetAppsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_GetAppsResponse(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.GetAppsResponse)) {
    throw new Error('Expected argument of type webhooks.v1.GetAppsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_GetAppsResponse(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.GetAppsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_GetEndpointsRequest(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.GetEndpointsRequest)) {
    throw new Error('Expected argument of type webhooks.v1.GetEndpointsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_GetEndpointsRequest(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.GetEndpointsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_GetEndpointsResponse(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.GetEndpointsResponse)) {
    throw new Error('Expected argument of type webhooks.v1.GetEndpointsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_GetEndpointsResponse(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.GetEndpointsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_GetEventTypesRequest(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.GetEventTypesRequest)) {
    throw new Error('Expected argument of type webhooks.v1.GetEventTypesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_GetEventTypesRequest(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.GetEventTypesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_GetEventTypesResponse(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.GetEventTypesResponse)) {
    throw new Error('Expected argument of type webhooks.v1.GetEventTypesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_GetEventTypesResponse(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.GetEventTypesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_GetMessagesRequest(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.GetMessagesRequest)) {
    throw new Error('Expected argument of type webhooks.v1.GetMessagesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_GetMessagesRequest(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.GetMessagesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_GetMessagesResponse(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.GetMessagesResponse)) {
    throw new Error('Expected argument of type webhooks.v1.GetMessagesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_GetMessagesResponse(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.GetMessagesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_ListAppsRequest(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.ListAppsRequest)) {
    throw new Error('Expected argument of type webhooks.v1.ListAppsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_ListAppsRequest(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.ListAppsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_ListAppsResponse(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.ListAppsResponse)) {
    throw new Error('Expected argument of type webhooks.v1.ListAppsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_ListAppsResponse(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.ListAppsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_ListEndpointsRequest(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.ListEndpointsRequest)) {
    throw new Error('Expected argument of type webhooks.v1.ListEndpointsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_ListEndpointsRequest(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.ListEndpointsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_ListEndpointsResponse(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.ListEndpointsResponse)) {
    throw new Error('Expected argument of type webhooks.v1.ListEndpointsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_ListEndpointsResponse(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.ListEndpointsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_ListEventTypesRequest(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.ListEventTypesRequest)) {
    throw new Error('Expected argument of type webhooks.v1.ListEventTypesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_ListEventTypesRequest(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.ListEventTypesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_ListEventTypesResponse(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.ListEventTypesResponse)) {
    throw new Error('Expected argument of type webhooks.v1.ListEventTypesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_ListEventTypesResponse(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.ListEventTypesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_ListMessagesRequest(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.ListMessagesRequest)) {
    throw new Error('Expected argument of type webhooks.v1.ListMessagesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_ListMessagesRequest(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.ListMessagesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_webhooks_v1_ListMessagesResponse(arg) {
  if (!(arg instanceof app_webhook_grpc_v1_webhook_pb.ListMessagesResponse)) {
    throw new Error('Expected argument of type webhooks.v1.ListMessagesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_webhooks_v1_ListMessagesResponse(buffer_arg) {
  return app_webhook_grpc_v1_webhook_pb.ListMessagesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var WebHookServiceService = exports.WebHookServiceService = {
  getApps: {
    path: '/webhooks.v1.WebHookService/GetApps',
    requestStream: false,
    responseStream: false,
    requestType: app_webhook_grpc_v1_webhook_pb.GetAppsRequest,
    responseType: app_webhook_grpc_v1_webhook_pb.GetAppsResponse,
    requestSerialize: serialize_webhooks_v1_GetAppsRequest,
    requestDeserialize: deserialize_webhooks_v1_GetAppsRequest,
    responseSerialize: serialize_webhooks_v1_GetAppsResponse,
    responseDeserialize: deserialize_webhooks_v1_GetAppsResponse,
  },
  listApps: {
    path: '/webhooks.v1.WebHookService/ListApps',
    requestStream: false,
    responseStream: false,
    requestType: app_webhook_grpc_v1_webhook_pb.ListAppsRequest,
    responseType: app_webhook_grpc_v1_webhook_pb.ListAppsResponse,
    requestSerialize: serialize_webhooks_v1_ListAppsRequest,
    requestDeserialize: deserialize_webhooks_v1_ListAppsRequest,
    responseSerialize: serialize_webhooks_v1_ListAppsResponse,
    responseDeserialize: deserialize_webhooks_v1_ListAppsResponse,
  },
  createApps: {
    path: '/webhooks.v1.WebHookService/CreateApps',
    requestStream: false,
    responseStream: false,
    requestType: app_webhook_grpc_v1_webhook_pb.CreateAppsRequest,
    responseType: app_webhook_grpc_v1_webhook_pb.CreateAppsResponse,
    requestSerialize: serialize_webhooks_v1_CreateAppsRequest,
    requestDeserialize: deserialize_webhooks_v1_CreateAppsRequest,
    responseSerialize: serialize_webhooks_v1_CreateAppsResponse,
    responseDeserialize: deserialize_webhooks_v1_CreateAppsResponse,
  },
  getEndpoints: {
    path: '/webhooks.v1.WebHookService/GetEndpoints',
    requestStream: false,
    responseStream: false,
    requestType: app_webhook_grpc_v1_webhook_pb.GetEndpointsRequest,
    responseType: app_webhook_grpc_v1_webhook_pb.GetEndpointsResponse,
    requestSerialize: serialize_webhooks_v1_GetEndpointsRequest,
    requestDeserialize: deserialize_webhooks_v1_GetEndpointsRequest,
    responseSerialize: serialize_webhooks_v1_GetEndpointsResponse,
    responseDeserialize: deserialize_webhooks_v1_GetEndpointsResponse,
  },
  listEndpoints: {
    path: '/webhooks.v1.WebHookService/ListEndpoints',
    requestStream: false,
    responseStream: false,
    requestType: app_webhook_grpc_v1_webhook_pb.ListEndpointsRequest,
    responseType: app_webhook_grpc_v1_webhook_pb.ListEndpointsResponse,
    requestSerialize: serialize_webhooks_v1_ListEndpointsRequest,
    requestDeserialize: deserialize_webhooks_v1_ListEndpointsRequest,
    responseSerialize: serialize_webhooks_v1_ListEndpointsResponse,
    responseDeserialize: deserialize_webhooks_v1_ListEndpointsResponse,
  },
  createEndpoints: {
    path: '/webhooks.v1.WebHookService/CreateEndpoints',
    requestStream: false,
    responseStream: false,
    requestType: app_webhook_grpc_v1_webhook_pb.CreateEndpointsRequest,
    responseType: app_webhook_grpc_v1_webhook_pb.CreateEndpointsResponse,
    requestSerialize: serialize_webhooks_v1_CreateEndpointsRequest,
    requestDeserialize: deserialize_webhooks_v1_CreateEndpointsRequest,
    responseSerialize: serialize_webhooks_v1_CreateEndpointsResponse,
    responseDeserialize: deserialize_webhooks_v1_CreateEndpointsResponse,
  },
  getMessages: {
    path: '/webhooks.v1.WebHookService/GetMessages',
    requestStream: false,
    responseStream: false,
    requestType: app_webhook_grpc_v1_webhook_pb.GetMessagesRequest,
    responseType: app_webhook_grpc_v1_webhook_pb.GetMessagesResponse,
    requestSerialize: serialize_webhooks_v1_GetMessagesRequest,
    requestDeserialize: deserialize_webhooks_v1_GetMessagesRequest,
    responseSerialize: serialize_webhooks_v1_GetMessagesResponse,
    responseDeserialize: deserialize_webhooks_v1_GetMessagesResponse,
  },
  listMessages: {
    path: '/webhooks.v1.WebHookService/ListMessages',
    requestStream: false,
    responseStream: false,
    requestType: app_webhook_grpc_v1_webhook_pb.ListMessagesRequest,
    responseType: app_webhook_grpc_v1_webhook_pb.ListMessagesResponse,
    requestSerialize: serialize_webhooks_v1_ListMessagesRequest,
    requestDeserialize: deserialize_webhooks_v1_ListMessagesRequest,
    responseSerialize: serialize_webhooks_v1_ListMessagesResponse,
    responseDeserialize: deserialize_webhooks_v1_ListMessagesResponse,
  },
  createMessages: {
    path: '/webhooks.v1.WebHookService/CreateMessages',
    requestStream: false,
    responseStream: false,
    requestType: app_webhook_grpc_v1_webhook_pb.CreateMessagesRequest,
    responseType: app_webhook_grpc_v1_webhook_pb.CreateMessagesResponse,
    requestSerialize: serialize_webhooks_v1_CreateMessagesRequest,
    requestDeserialize: deserialize_webhooks_v1_CreateMessagesRequest,
    responseSerialize: serialize_webhooks_v1_CreateMessagesResponse,
    responseDeserialize: deserialize_webhooks_v1_CreateMessagesResponse,
  },
  getEventTypes: {
    path: '/webhooks.v1.WebHookService/GetEventTypes',
    requestStream: false,
    responseStream: false,
    requestType: app_webhook_grpc_v1_webhook_pb.GetEventTypesRequest,
    responseType: app_webhook_grpc_v1_webhook_pb.GetEventTypesResponse,
    requestSerialize: serialize_webhooks_v1_GetEventTypesRequest,
    requestDeserialize: deserialize_webhooks_v1_GetEventTypesRequest,
    responseSerialize: serialize_webhooks_v1_GetEventTypesResponse,
    responseDeserialize: deserialize_webhooks_v1_GetEventTypesResponse,
  },
  listEventTypes: {
    path: '/webhooks.v1.WebHookService/ListEventTypes',
    requestStream: false,
    responseStream: false,
    requestType: app_webhook_grpc_v1_webhook_pb.ListEventTypesRequest,
    responseType: app_webhook_grpc_v1_webhook_pb.ListEventTypesResponse,
    requestSerialize: serialize_webhooks_v1_ListEventTypesRequest,
    requestDeserialize: deserialize_webhooks_v1_ListEventTypesRequest,
    responseSerialize: serialize_webhooks_v1_ListEventTypesResponse,
    responseDeserialize: deserialize_webhooks_v1_ListEventTypesResponse,
  },
  createEventTypes: {
    path: '/webhooks.v1.WebHookService/CreateEventTypes',
    requestStream: false,
    responseStream: false,
    requestType: app_webhook_grpc_v1_webhook_pb.CreateEventTypesRequest,
    responseType: app_webhook_grpc_v1_webhook_pb.CreateEventTypesResponse,
    requestSerialize: serialize_webhooks_v1_CreateEventTypesRequest,
    requestDeserialize: deserialize_webhooks_v1_CreateEventTypesRequest,
    responseSerialize: serialize_webhooks_v1_CreateEventTypesResponse,
    responseDeserialize: deserialize_webhooks_v1_CreateEventTypesResponse,
  },
};

exports.WebHookServiceClient = grpc.makeGenericClientConstructor(WebHookServiceService);
