// @generated by protoc-gen-es v1.6.0 with parameter "target=ts"
// @generated from file app/router/grpc/v1/event.proto (package events.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message events.v1.EventType
 */
export class EventType extends Message<EventType> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: bytes schema = 2;
   */
  schema = new Uint8Array(0);

  constructor(data?: PartialMessage<EventType>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.EventType";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "schema", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EventType {
    return new EventType().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EventType {
    return new EventType().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EventType {
    return new EventType().fromJsonString(jsonString, options);
  }

  static equals(a: EventType | PlainMessage<EventType> | undefined, b: EventType | PlainMessage<EventType> | undefined): boolean {
    return proto3.util.equals(EventType, a, b);
  }
}

/**
 * @generated from message events.v1.Provider
 */
export class Provider extends Message<Provider> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<Provider>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.Provider";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Provider {
    return new Provider().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Provider {
    return new Provider().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Provider {
    return new Provider().fromJsonString(jsonString, options);
  }

  static equals(a: Provider | PlainMessage<Provider> | undefined, b: Provider | PlainMessage<Provider> | undefined): boolean {
    return proto3.util.equals(Provider, a, b);
  }
}

/**
 * @generated from message events.v1.Subscription
 */
export class Subscription extends Message<Subscription> {
  /**
   * @generated from field: string tenant_id = 1;
   */
  tenantId = "";

  /**
   * @generated from field: string id = 2;
   */
  id = "";

  /**
   * @generated from field: string notification = 3;
   */
  notification = "";

  /**
   * @generated from field: bytes metadata = 4;
   */
  metadata = new Uint8Array(0);

  constructor(data?: PartialMessage<Subscription>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.Subscription";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "tenant_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "notification", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "metadata", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Subscription {
    return new Subscription().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Subscription {
    return new Subscription().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Subscription {
    return new Subscription().fromJsonString(jsonString, options);
  }

  static equals(a: Subscription | PlainMessage<Subscription> | undefined, b: Subscription | PlainMessage<Subscription> | undefined): boolean {
    return proto3.util.equals(Subscription, a, b);
  }
}

/**
 * @generated from message events.v1.GetEventTypesRequest
 */
export class GetEventTypesRequest extends Message<GetEventTypesRequest> {
  /**
   * @generated from field: repeated string ids = 1;
   */
  ids: string[] = [];

  constructor(data?: PartialMessage<GetEventTypesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.GetEventTypesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetEventTypesRequest {
    return new GetEventTypesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetEventTypesRequest {
    return new GetEventTypesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetEventTypesRequest {
    return new GetEventTypesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetEventTypesRequest | PlainMessage<GetEventTypesRequest> | undefined, b: GetEventTypesRequest | PlainMessage<GetEventTypesRequest> | undefined): boolean {
    return proto3.util.equals(GetEventTypesRequest, a, b);
  }
}

/**
 * @generated from message events.v1.GetEventTypesResponse
 */
export class GetEventTypesResponse extends Message<GetEventTypesResponse> {
  /**
   * @generated from field: map<string, events.v1.EventType> data = 1;
   */
  data: { [key: string]: EventType } = {};

  /**
   * @generated from field: repeated events.v1.Error errors = 2;
   */
  errors: Error[] = [];

  constructor(data?: PartialMessage<GetEventTypesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.GetEventTypesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: EventType} },
    { no: 2, name: "errors", kind: "message", T: Error, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetEventTypesResponse {
    return new GetEventTypesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetEventTypesResponse {
    return new GetEventTypesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetEventTypesResponse {
    return new GetEventTypesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetEventTypesResponse | PlainMessage<GetEventTypesResponse> | undefined, b: GetEventTypesResponse | PlainMessage<GetEventTypesResponse> | undefined): boolean {
    return proto3.util.equals(GetEventTypesResponse, a, b);
  }
}

/**
 * @generated from message events.v1.GetProvidersRequest
 */
export class GetProvidersRequest extends Message<GetProvidersRequest> {
  /**
   * @generated from field: string tenant_id = 1;
   */
  tenantId = "";

  /**
   * @generated from field: repeated string ids = 2;
   */
  ids: string[] = [];

  constructor(data?: PartialMessage<GetProvidersRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.GetProvidersRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "tenant_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetProvidersRequest {
    return new GetProvidersRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetProvidersRequest {
    return new GetProvidersRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetProvidersRequest {
    return new GetProvidersRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetProvidersRequest | PlainMessage<GetProvidersRequest> | undefined, b: GetProvidersRequest | PlainMessage<GetProvidersRequest> | undefined): boolean {
    return proto3.util.equals(GetProvidersRequest, a, b);
  }
}

/**
 * @generated from message events.v1.GetProvidersResponse
 */
export class GetProvidersResponse extends Message<GetProvidersResponse> {
  /**
   * @generated from field: map<string, events.v1.Provider> data = 1;
   */
  data: { [key: string]: Provider } = {};

  /**
   * @generated from field: repeated events.v1.Error errors = 2;
   */
  errors: Error[] = [];

  constructor(data?: PartialMessage<GetProvidersResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.GetProvidersResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Provider} },
    { no: 2, name: "errors", kind: "message", T: Error, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetProvidersResponse {
    return new GetProvidersResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetProvidersResponse {
    return new GetProvidersResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetProvidersResponse {
    return new GetProvidersResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetProvidersResponse | PlainMessage<GetProvidersResponse> | undefined, b: GetProvidersResponse | PlainMessage<GetProvidersResponse> | undefined): boolean {
    return proto3.util.equals(GetProvidersResponse, a, b);
  }
}

/**
 * @generated from message events.v1.GetSubscriptionsRequest
 */
export class GetSubscriptionsRequest extends Message<GetSubscriptionsRequest> {
  /**
   * @generated from field: string tenant_id = 1;
   */
  tenantId = "";

  /**
   * @generated from field: repeated string ids = 2;
   */
  ids: string[] = [];

  constructor(data?: PartialMessage<GetSubscriptionsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.GetSubscriptionsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "tenant_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSubscriptionsRequest {
    return new GetSubscriptionsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSubscriptionsRequest {
    return new GetSubscriptionsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSubscriptionsRequest {
    return new GetSubscriptionsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetSubscriptionsRequest | PlainMessage<GetSubscriptionsRequest> | undefined, b: GetSubscriptionsRequest | PlainMessage<GetSubscriptionsRequest> | undefined): boolean {
    return proto3.util.equals(GetSubscriptionsRequest, a, b);
  }
}

/**
 * @generated from message events.v1.GetSubscriptionsResponse
 */
export class GetSubscriptionsResponse extends Message<GetSubscriptionsResponse> {
  /**
   * @generated from field: map<string, events.v1.Subscription> data = 1;
   */
  data: { [key: string]: Subscription } = {};

  /**
   * @generated from field: repeated events.v1.Error errors = 2;
   */
  errors: Error[] = [];

  constructor(data?: PartialMessage<GetSubscriptionsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.GetSubscriptionsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Subscription} },
    { no: 2, name: "errors", kind: "message", T: Error, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSubscriptionsResponse {
    return new GetSubscriptionsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSubscriptionsResponse {
    return new GetSubscriptionsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSubscriptionsResponse {
    return new GetSubscriptionsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetSubscriptionsResponse | PlainMessage<GetSubscriptionsResponse> | undefined, b: GetSubscriptionsResponse | PlainMessage<GetSubscriptionsResponse> | undefined): boolean {
    return proto3.util.equals(GetSubscriptionsResponse, a, b);
  }
}

/**
 * @generated from message events.v1.ListEventTypesRequest
 */
export class ListEventTypesRequest extends Message<ListEventTypesRequest> {
  /**
   * @generated from field: events.v1.PageRequest page = 1;
   */
  page?: PageRequest;

  constructor(data?: PartialMessage<ListEventTypesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.ListEventTypesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "page", kind: "message", T: PageRequest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEventTypesRequest {
    return new ListEventTypesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEventTypesRequest {
    return new ListEventTypesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEventTypesRequest {
    return new ListEventTypesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListEventTypesRequest | PlainMessage<ListEventTypesRequest> | undefined, b: ListEventTypesRequest | PlainMessage<ListEventTypesRequest> | undefined): boolean {
    return proto3.util.equals(ListEventTypesRequest, a, b);
  }
}

/**
 * @generated from message events.v1.ListEventTypesResponse
 */
export class ListEventTypesResponse extends Message<ListEventTypesResponse> {
  /**
   * @generated from field: repeated events.v1.EventType data = 1;
   */
  data: EventType[] = [];

  /**
   * @generated from field: events.v1.PageResponse page = 2;
   */
  page?: PageResponse;

  constructor(data?: PartialMessage<ListEventTypesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.ListEventTypesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: EventType, repeated: true },
    { no: 2, name: "page", kind: "message", T: PageResponse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEventTypesResponse {
    return new ListEventTypesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEventTypesResponse {
    return new ListEventTypesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEventTypesResponse {
    return new ListEventTypesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListEventTypesResponse | PlainMessage<ListEventTypesResponse> | undefined, b: ListEventTypesResponse | PlainMessage<ListEventTypesResponse> | undefined): boolean {
    return proto3.util.equals(ListEventTypesResponse, a, b);
  }
}

/**
 * @generated from message events.v1.ListSubscriptionsRequest
 */
export class ListSubscriptionsRequest extends Message<ListSubscriptionsRequest> {
  /**
   * @generated from field: events.v1.PageRequest page = 1;
   */
  page?: PageRequest;

  constructor(data?: PartialMessage<ListSubscriptionsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.ListSubscriptionsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "page", kind: "message", T: PageRequest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSubscriptionsRequest {
    return new ListSubscriptionsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSubscriptionsRequest {
    return new ListSubscriptionsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSubscriptionsRequest {
    return new ListSubscriptionsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListSubscriptionsRequest | PlainMessage<ListSubscriptionsRequest> | undefined, b: ListSubscriptionsRequest | PlainMessage<ListSubscriptionsRequest> | undefined): boolean {
    return proto3.util.equals(ListSubscriptionsRequest, a, b);
  }
}

/**
 * @generated from message events.v1.ListSubscriptionsResponse
 */
export class ListSubscriptionsResponse extends Message<ListSubscriptionsResponse> {
  /**
   * @generated from field: repeated events.v1.Subscription data = 1;
   */
  data: Subscription[] = [];

  /**
   * @generated from field: events.v1.PageResponse page = 2;
   */
  page?: PageResponse;

  constructor(data?: PartialMessage<ListSubscriptionsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.ListSubscriptionsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: Subscription, repeated: true },
    { no: 2, name: "page", kind: "message", T: PageResponse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSubscriptionsResponse {
    return new ListSubscriptionsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSubscriptionsResponse {
    return new ListSubscriptionsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSubscriptionsResponse {
    return new ListSubscriptionsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListSubscriptionsResponse | PlainMessage<ListSubscriptionsResponse> | undefined, b: ListSubscriptionsResponse | PlainMessage<ListSubscriptionsResponse> | undefined): boolean {
    return proto3.util.equals(ListSubscriptionsResponse, a, b);
  }
}

/**
 * @generated from message events.v1.ListProvidersRequest
 */
export class ListProvidersRequest extends Message<ListProvidersRequest> {
  /**
   * @generated from field: events.v1.PageRequest page = 1;
   */
  page?: PageRequest;

  constructor(data?: PartialMessage<ListProvidersRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.ListProvidersRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "page", kind: "message", T: PageRequest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListProvidersRequest {
    return new ListProvidersRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListProvidersRequest {
    return new ListProvidersRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListProvidersRequest {
    return new ListProvidersRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListProvidersRequest | PlainMessage<ListProvidersRequest> | undefined, b: ListProvidersRequest | PlainMessage<ListProvidersRequest> | undefined): boolean {
    return proto3.util.equals(ListProvidersRequest, a, b);
  }
}

/**
 * @generated from message events.v1.ListProvidersResponse
 */
export class ListProvidersResponse extends Message<ListProvidersResponse> {
  /**
   * @generated from field: repeated events.v1.Provider data = 1;
   */
  data: Provider[] = [];

  /**
   * @generated from field: events.v1.PageResponse page = 2;
   */
  page?: PageResponse;

  constructor(data?: PartialMessage<ListProvidersResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.ListProvidersResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: Provider, repeated: true },
    { no: 2, name: "page", kind: "message", T: PageResponse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListProvidersResponse {
    return new ListProvidersResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListProvidersResponse {
    return new ListProvidersResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListProvidersResponse {
    return new ListProvidersResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListProvidersResponse | PlainMessage<ListProvidersResponse> | undefined, b: ListProvidersResponse | PlainMessage<ListProvidersResponse> | undefined): boolean {
    return proto3.util.equals(ListProvidersResponse, a, b);
  }
}

/**
 * @generated from message events.v1.CreateEventTypesRequest
 */
export class CreateEventTypesRequest extends Message<CreateEventTypesRequest> {
  /**
   * @generated from field: repeated events.v1.EventType data = 1;
   */
  data: EventType[] = [];

  constructor(data?: PartialMessage<CreateEventTypesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.CreateEventTypesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: EventType, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateEventTypesRequest {
    return new CreateEventTypesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateEventTypesRequest {
    return new CreateEventTypesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateEventTypesRequest {
    return new CreateEventTypesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateEventTypesRequest | PlainMessage<CreateEventTypesRequest> | undefined, b: CreateEventTypesRequest | PlainMessage<CreateEventTypesRequest> | undefined): boolean {
    return proto3.util.equals(CreateEventTypesRequest, a, b);
  }
}

/**
 * @generated from message events.v1.CreateEventTypesResponse
 */
export class CreateEventTypesResponse extends Message<CreateEventTypesResponse> {
  /**
   * @generated from field: repeated events.v1.EventType data = 1;
   */
  data: EventType[] = [];

  /**
   * @generated from field: repeated events.v1.Error errors = 2;
   */
  errors: Error[] = [];

  constructor(data?: PartialMessage<CreateEventTypesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.CreateEventTypesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: EventType, repeated: true },
    { no: 2, name: "errors", kind: "message", T: Error, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateEventTypesResponse {
    return new CreateEventTypesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateEventTypesResponse {
    return new CreateEventTypesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateEventTypesResponse {
    return new CreateEventTypesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateEventTypesResponse | PlainMessage<CreateEventTypesResponse> | undefined, b: CreateEventTypesResponse | PlainMessage<CreateEventTypesResponse> | undefined): boolean {
    return proto3.util.equals(CreateEventTypesResponse, a, b);
  }
}

/**
 * @generated from message events.v1.CreateSubscriptionsRequest
 */
export class CreateSubscriptionsRequest extends Message<CreateSubscriptionsRequest> {
  /**
   * @generated from field: repeated events.v1.Subscription data = 1;
   */
  data: Subscription[] = [];

  constructor(data?: PartialMessage<CreateSubscriptionsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.CreateSubscriptionsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: Subscription, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSubscriptionsRequest {
    return new CreateSubscriptionsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSubscriptionsRequest {
    return new CreateSubscriptionsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSubscriptionsRequest {
    return new CreateSubscriptionsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSubscriptionsRequest | PlainMessage<CreateSubscriptionsRequest> | undefined, b: CreateSubscriptionsRequest | PlainMessage<CreateSubscriptionsRequest> | undefined): boolean {
    return proto3.util.equals(CreateSubscriptionsRequest, a, b);
  }
}

/**
 * @generated from message events.v1.CreateSubscriptionsResponse
 */
export class CreateSubscriptionsResponse extends Message<CreateSubscriptionsResponse> {
  /**
   * @generated from field: repeated events.v1.Subscription data = 1;
   */
  data: Subscription[] = [];

  /**
   * @generated from field: repeated events.v1.Error errors = 2;
   */
  errors: Error[] = [];

  constructor(data?: PartialMessage<CreateSubscriptionsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.CreateSubscriptionsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: Subscription, repeated: true },
    { no: 2, name: "errors", kind: "message", T: Error, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSubscriptionsResponse {
    return new CreateSubscriptionsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSubscriptionsResponse {
    return new CreateSubscriptionsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSubscriptionsResponse {
    return new CreateSubscriptionsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSubscriptionsResponse | PlainMessage<CreateSubscriptionsResponse> | undefined, b: CreateSubscriptionsResponse | PlainMessage<CreateSubscriptionsResponse> | undefined): boolean {
    return proto3.util.equals(CreateSubscriptionsResponse, a, b);
  }
}

/**
 * @generated from message events.v1.PageRequest
 */
export class PageRequest extends Message<PageRequest> {
  /**
   * @generated from field: int32 limit = 1;
   */
  limit = 0;

  /**
   * @generated from field: string cursor = 2;
   */
  cursor = "";

  /**
   * @generated from field: bool desc = 3;
   */
  desc = false;

  constructor(data?: PartialMessage<PageRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.PageRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "cursor", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "desc", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PageRequest {
    return new PageRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PageRequest {
    return new PageRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PageRequest {
    return new PageRequest().fromJsonString(jsonString, options);
  }

  static equals(a: PageRequest | PlainMessage<PageRequest> | undefined, b: PageRequest | PlainMessage<PageRequest> | undefined): boolean {
    return proto3.util.equals(PageRequest, a, b);
  }
}

/**
 * @generated from message events.v1.PageResponse
 */
export class PageResponse extends Message<PageResponse> {
  /**
   * @generated from field: string next = 1;
   */
  next = "";

  /**
   * @generated from field: string prev = 2;
   */
  prev = "";

  /**
   * @generated from field: bool done = 3;
   */
  done = false;

  constructor(data?: PartialMessage<PageResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.PageResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "next", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "prev", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "done", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PageResponse {
    return new PageResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PageResponse {
    return new PageResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PageResponse {
    return new PageResponse().fromJsonString(jsonString, options);
  }

  static equals(a: PageResponse | PlainMessage<PageResponse> | undefined, b: PageResponse | PlainMessage<PageResponse> | undefined): boolean {
    return proto3.util.equals(PageResponse, a, b);
  }
}

/**
 * @generated from message events.v1.Error
 */
export class Error extends Message<Error> {
  /**
   * @generated from field: int32 code = 1;
   */
  code = 0;

  /**
   * @generated from field: string index = 2;
   */
  index = "";

  /**
   * @generated from field: string reason = 3;
   */
  reason = "";

  constructor(data?: PartialMessage<Error>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.Error";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "code", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "index", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "reason", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Error {
    return new Error().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Error {
    return new Error().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Error {
    return new Error().fromJsonString(jsonString, options);
  }

  static equals(a: Error | PlainMessage<Error> | undefined, b: Error | PlainMessage<Error> | undefined): boolean {
    return proto3.util.equals(Error, a, b);
  }
}

