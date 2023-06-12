/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { LinkedRequesters } from "./linked_requesters";
import { Params } from "./params";
import { SlashHistory } from "./slash_history";
import { User } from "./user";

export const protobufPackage = "dhpc.user";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetLinkedRequestersRequest {
  index: string;
}

export interface QueryGetLinkedRequestersResponse {
  linkedRequesters: LinkedRequesters | undefined;
}

export interface QueryAllLinkedRequestersRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLinkedRequestersResponse {
  linkedRequesters: LinkedRequesters[];
  pagination: PageResponse | undefined;
}

export interface QueryGetSlashHistoryRequest {
  index: string;
}

export interface QueryGetSlashHistoryResponse {
  slashHistory: SlashHistory | undefined;
}

export interface QueryAllSlashHistoryRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSlashHistoryResponse {
  slashHistory: SlashHistory[];
  pagination: PageResponse | undefined;
}

export interface QueryGetUserRequest {
  account: string;
}

export interface QueryGetUserResponse {
  user: User | undefined;
}

export interface QueryAllUserRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllUserResponse {
  user: User[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetLinkedRequestersRequest(): QueryGetLinkedRequestersRequest {
  return { index: "" };
}

export const QueryGetLinkedRequestersRequest = {
  encode(message: QueryGetLinkedRequestersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetLinkedRequestersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetLinkedRequestersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLinkedRequestersRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetLinkedRequestersRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetLinkedRequestersRequest>, I>>(
    object: I,
  ): QueryGetLinkedRequestersRequest {
    const message = createBaseQueryGetLinkedRequestersRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetLinkedRequestersResponse(): QueryGetLinkedRequestersResponse {
  return { linkedRequesters: undefined };
}

export const QueryGetLinkedRequestersResponse = {
  encode(message: QueryGetLinkedRequestersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.linkedRequesters !== undefined) {
      LinkedRequesters.encode(message.linkedRequesters, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetLinkedRequestersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetLinkedRequestersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.linkedRequesters = LinkedRequesters.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLinkedRequestersResponse {
    return {
      linkedRequesters: isSet(object.linkedRequesters) ? LinkedRequesters.fromJSON(object.linkedRequesters) : undefined,
    };
  },

  toJSON(message: QueryGetLinkedRequestersResponse): unknown {
    const obj: any = {};
    message.linkedRequesters !== undefined && (obj.linkedRequesters = message.linkedRequesters
      ? LinkedRequesters.toJSON(message.linkedRequesters)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetLinkedRequestersResponse>, I>>(
    object: I,
  ): QueryGetLinkedRequestersResponse {
    const message = createBaseQueryGetLinkedRequestersResponse();
    message.linkedRequesters = (object.linkedRequesters !== undefined && object.linkedRequesters !== null)
      ? LinkedRequesters.fromPartial(object.linkedRequesters)
      : undefined;
    return message;
  },
};

function createBaseQueryAllLinkedRequestersRequest(): QueryAllLinkedRequestersRequest {
  return { pagination: undefined };
}

export const QueryAllLinkedRequestersRequest = {
  encode(message: QueryAllLinkedRequestersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllLinkedRequestersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllLinkedRequestersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLinkedRequestersRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllLinkedRequestersRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllLinkedRequestersRequest>, I>>(
    object: I,
  ): QueryAllLinkedRequestersRequest {
    const message = createBaseQueryAllLinkedRequestersRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllLinkedRequestersResponse(): QueryAllLinkedRequestersResponse {
  return { linkedRequesters: [], pagination: undefined };
}

export const QueryAllLinkedRequestersResponse = {
  encode(message: QueryAllLinkedRequestersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.linkedRequesters) {
      LinkedRequesters.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllLinkedRequestersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllLinkedRequestersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.linkedRequesters.push(LinkedRequesters.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLinkedRequestersResponse {
    return {
      linkedRequesters: Array.isArray(object?.linkedRequesters)
        ? object.linkedRequesters.map((e: any) => LinkedRequesters.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllLinkedRequestersResponse): unknown {
    const obj: any = {};
    if (message.linkedRequesters) {
      obj.linkedRequesters = message.linkedRequesters.map((e) => e ? LinkedRequesters.toJSON(e) : undefined);
    } else {
      obj.linkedRequesters = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllLinkedRequestersResponse>, I>>(
    object: I,
  ): QueryAllLinkedRequestersResponse {
    const message = createBaseQueryAllLinkedRequestersResponse();
    message.linkedRequesters = object.linkedRequesters?.map((e) => LinkedRequesters.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetSlashHistoryRequest(): QueryGetSlashHistoryRequest {
  return { index: "" };
}

export const QueryGetSlashHistoryRequest = {
  encode(message: QueryGetSlashHistoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetSlashHistoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetSlashHistoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSlashHistoryRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetSlashHistoryRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetSlashHistoryRequest>, I>>(object: I): QueryGetSlashHistoryRequest {
    const message = createBaseQueryGetSlashHistoryRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetSlashHistoryResponse(): QueryGetSlashHistoryResponse {
  return { slashHistory: undefined };
}

export const QueryGetSlashHistoryResponse = {
  encode(message: QueryGetSlashHistoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.slashHistory !== undefined) {
      SlashHistory.encode(message.slashHistory, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetSlashHistoryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetSlashHistoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.slashHistory = SlashHistory.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSlashHistoryResponse {
    return { slashHistory: isSet(object.slashHistory) ? SlashHistory.fromJSON(object.slashHistory) : undefined };
  },

  toJSON(message: QueryGetSlashHistoryResponse): unknown {
    const obj: any = {};
    message.slashHistory !== undefined
      && (obj.slashHistory = message.slashHistory ? SlashHistory.toJSON(message.slashHistory) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetSlashHistoryResponse>, I>>(object: I): QueryGetSlashHistoryResponse {
    const message = createBaseQueryGetSlashHistoryResponse();
    message.slashHistory = (object.slashHistory !== undefined && object.slashHistory !== null)
      ? SlashHistory.fromPartial(object.slashHistory)
      : undefined;
    return message;
  },
};

function createBaseQueryAllSlashHistoryRequest(): QueryAllSlashHistoryRequest {
  return { pagination: undefined };
}

export const QueryAllSlashHistoryRequest = {
  encode(message: QueryAllSlashHistoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllSlashHistoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllSlashHistoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSlashHistoryRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllSlashHistoryRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllSlashHistoryRequest>, I>>(object: I): QueryAllSlashHistoryRequest {
    const message = createBaseQueryAllSlashHistoryRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllSlashHistoryResponse(): QueryAllSlashHistoryResponse {
  return { slashHistory: [], pagination: undefined };
}

export const QueryAllSlashHistoryResponse = {
  encode(message: QueryAllSlashHistoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.slashHistory) {
      SlashHistory.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllSlashHistoryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllSlashHistoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.slashHistory.push(SlashHistory.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSlashHistoryResponse {
    return {
      slashHistory: Array.isArray(object?.slashHistory)
        ? object.slashHistory.map((e: any) => SlashHistory.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllSlashHistoryResponse): unknown {
    const obj: any = {};
    if (message.slashHistory) {
      obj.slashHistory = message.slashHistory.map((e) => e ? SlashHistory.toJSON(e) : undefined);
    } else {
      obj.slashHistory = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllSlashHistoryResponse>, I>>(object: I): QueryAllSlashHistoryResponse {
    const message = createBaseQueryAllSlashHistoryResponse();
    message.slashHistory = object.slashHistory?.map((e) => SlashHistory.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetUserRequest(): QueryGetUserRequest {
  return { account: "" };
}

export const QueryGetUserRequest = {
  encode(message: QueryGetUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== "") {
      writer.uint32(10).string(message.account);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.account = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserRequest {
    return { account: isSet(object.account) ? String(object.account) : "" };
  },

  toJSON(message: QueryGetUserRequest): unknown {
    const obj: any = {};
    message.account !== undefined && (obj.account = message.account);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserRequest>, I>>(object: I): QueryGetUserRequest {
    const message = createBaseQueryGetUserRequest();
    message.account = object.account ?? "";
    return message;
  },
};

function createBaseQueryGetUserResponse(): QueryGetUserResponse {
  return { user: undefined };
}

export const QueryGetUserResponse = {
  encode(message: QueryGetUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserResponse {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: QueryGetUserResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserResponse>, I>>(object: I): QueryGetUserResponse {
    const message = createBaseQueryGetUserResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseQueryAllUserRequest(): QueryAllUserRequest {
  return { pagination: undefined };
}

export const QueryAllUserRequest = {
  encode(message: QueryAllUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllUserRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserRequest>, I>>(object: I): QueryAllUserRequest {
    const message = createBaseQueryAllUserRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserResponse(): QueryAllUserResponse {
  return { user: [], pagination: undefined };
}

export const QueryAllUserResponse = {
  encode(message: QueryAllUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.user) {
      User.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user.push(User.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserResponse {
    return {
      user: Array.isArray(object?.user) ? object.user.map((e: any) => User.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllUserResponse): unknown {
    const obj: any = {};
    if (message.user) {
      obj.user = message.user.map((e) => e ? User.toJSON(e) : undefined);
    } else {
      obj.user = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserResponse>, I>>(object: I): QueryAllUserResponse {
    const message = createBaseQueryAllUserResponse();
    message.user = object.user?.map((e) => User.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of LinkedRequesters items. */
  LinkedRequesters(request: QueryGetLinkedRequestersRequest): Promise<QueryGetLinkedRequestersResponse>;
  LinkedRequestersAll(request: QueryAllLinkedRequestersRequest): Promise<QueryAllLinkedRequestersResponse>;
  /** Queries a list of SlashHistory items. */
  SlashHistory(request: QueryGetSlashHistoryRequest): Promise<QueryGetSlashHistoryResponse>;
  SlashHistoryAll(request: QueryAllSlashHistoryRequest): Promise<QueryAllSlashHistoryResponse>;
  /** Queries a list of User items. */
  User(request: QueryGetUserRequest): Promise<QueryGetUserResponse>;
  UserAll(request: QueryAllUserRequest): Promise<QueryAllUserResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.LinkedRequesters = this.LinkedRequesters.bind(this);
    this.LinkedRequestersAll = this.LinkedRequestersAll.bind(this);
    this.SlashHistory = this.SlashHistory.bind(this);
    this.SlashHistoryAll = this.SlashHistoryAll.bind(this);
    this.User = this.User.bind(this);
    this.UserAll = this.UserAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.user.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  LinkedRequesters(request: QueryGetLinkedRequestersRequest): Promise<QueryGetLinkedRequestersResponse> {
    const data = QueryGetLinkedRequestersRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.user.Query", "LinkedRequesters", data);
    return promise.then((data) => QueryGetLinkedRequestersResponse.decode(new _m0.Reader(data)));
  }

  LinkedRequestersAll(request: QueryAllLinkedRequestersRequest): Promise<QueryAllLinkedRequestersResponse> {
    const data = QueryAllLinkedRequestersRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.user.Query", "LinkedRequestersAll", data);
    return promise.then((data) => QueryAllLinkedRequestersResponse.decode(new _m0.Reader(data)));
  }

  SlashHistory(request: QueryGetSlashHistoryRequest): Promise<QueryGetSlashHistoryResponse> {
    const data = QueryGetSlashHistoryRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.user.Query", "SlashHistory", data);
    return promise.then((data) => QueryGetSlashHistoryResponse.decode(new _m0.Reader(data)));
  }

  SlashHistoryAll(request: QueryAllSlashHistoryRequest): Promise<QueryAllSlashHistoryResponse> {
    const data = QueryAllSlashHistoryRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.user.Query", "SlashHistoryAll", data);
    return promise.then((data) => QueryAllSlashHistoryResponse.decode(new _m0.Reader(data)));
  }

  User(request: QueryGetUserRequest): Promise<QueryGetUserResponse> {
    const data = QueryGetUserRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.user.Query", "User", data);
    return promise.then((data) => QueryGetUserResponse.decode(new _m0.Reader(data)));
  }

  UserAll(request: QueryAllUserRequest): Promise<QueryAllUserResponse> {
    const data = QueryAllUserRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.user.Query", "UserAll", data);
    return promise.then((data) => QueryAllUserResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
