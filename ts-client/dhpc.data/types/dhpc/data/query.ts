/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Data } from "./data";
import { Params } from "./params";

export const protobufPackage = "dhpc.data";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetDataRequest {
  address: string;
  owner: string;
  network: string;
}

export interface QueryAllDataRequestByAddr {
  address: string;
  pagination: PageRequest | undefined;
}

export interface QueryGetDataResponse {
  data: Data | undefined;
}

export interface QueryAllDataRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDataResponse {
  data: Data[];
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

function createBaseQueryGetDataRequest(): QueryGetDataRequest {
  return { address: "", owner: "", network: "" };
}

export const QueryGetDataRequest = {
  encode(message: QueryGetDataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.network !== "") {
      writer.uint32(26).string(message.network);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDataRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.network = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDataRequest {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
      network: isSet(object.network) ? String(object.network) : "",
    };
  },

  toJSON(message: QueryGetDataRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.owner !== undefined && (obj.owner = message.owner);
    message.network !== undefined && (obj.network = message.network);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDataRequest>, I>>(object: I): QueryGetDataRequest {
    const message = createBaseQueryGetDataRequest();
    message.address = object.address ?? "";
    message.owner = object.owner ?? "";
    message.network = object.network ?? "";
    return message;
  },
};

function createBaseQueryAllDataRequestByAddr(): QueryAllDataRequestByAddr {
  return { address: "", pagination: undefined };
}

export const QueryAllDataRequestByAddr = {
  encode(message: QueryAllDataRequestByAddr, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDataRequestByAddr {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDataRequestByAddr();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllDataRequestByAddr {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDataRequestByAddr): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDataRequestByAddr>, I>>(object: I): QueryAllDataRequestByAddr {
    const message = createBaseQueryAllDataRequestByAddr();
    message.address = object.address ?? "";
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetDataResponse(): QueryGetDataResponse {
  return { data: undefined };
}

export const QueryGetDataResponse = {
  encode(message: QueryGetDataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.data !== undefined) {
      Data.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDataResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = Data.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDataResponse {
    return { data: isSet(object.data) ? Data.fromJSON(object.data) : undefined };
  },

  toJSON(message: QueryGetDataResponse): unknown {
    const obj: any = {};
    message.data !== undefined && (obj.data = message.data ? Data.toJSON(message.data) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDataResponse>, I>>(object: I): QueryGetDataResponse {
    const message = createBaseQueryGetDataResponse();
    message.data = (object.data !== undefined && object.data !== null) ? Data.fromPartial(object.data) : undefined;
    return message;
  },
};

function createBaseQueryAllDataRequest(): QueryAllDataRequest {
  return { pagination: undefined };
}

export const QueryAllDataRequest = {
  encode(message: QueryAllDataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDataRequest();
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

  fromJSON(object: any): QueryAllDataRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllDataRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDataRequest>, I>>(object: I): QueryAllDataRequest {
    const message = createBaseQueryAllDataRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDataResponse(): QueryAllDataResponse {
  return { data: [], pagination: undefined };
}

export const QueryAllDataResponse = {
  encode(message: QueryAllDataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.data) {
      Data.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDataResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data.push(Data.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllDataResponse {
    return {
      data: Array.isArray(object?.data) ? object.data.map((e: any) => Data.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDataResponse): unknown {
    const obj: any = {};
    if (message.data) {
      obj.data = message.data.map((e) => e ? Data.toJSON(e) : undefined);
    } else {
      obj.data = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDataResponse>, I>>(object: I): QueryAllDataResponse {
    const message = createBaseQueryAllDataResponse();
    message.data = object.data?.map((e) => Data.fromPartial(e)) || [];
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
  /** Queries a list of Data items. */
  Data(request: QueryGetDataRequest): Promise<QueryGetDataResponse>;
  DataAll(request: QueryAllDataRequest): Promise<QueryAllDataResponse>;
  DataAllByAddr(request: QueryAllDataRequestByAddr): Promise<QueryAllDataResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Data = this.Data.bind(this);
    this.DataAll = this.DataAll.bind(this);
    this.DataAllByAddr = this.DataAllByAddr.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.data.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Data(request: QueryGetDataRequest): Promise<QueryGetDataResponse> {
    const data = QueryGetDataRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.data.Query", "Data", data);
    return promise.then((data) => QueryGetDataResponse.decode(new _m0.Reader(data)));
  }

  DataAll(request: QueryAllDataRequest): Promise<QueryAllDataResponse> {
    const data = QueryAllDataRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.data.Query", "DataAll", data);
    return promise.then((data) => QueryAllDataResponse.decode(new _m0.Reader(data)));
  }

  DataAllByAddr(request: QueryAllDataRequestByAddr): Promise<QueryAllDataResponse> {
    const data = QueryAllDataRequestByAddr.encode(request).finish();
    const promise = this.rpc.request("dhpc.data.Query", "DataAllByAddr", data);
    return promise.then((data) => QueryAllDataResponse.decode(new _m0.Reader(data)));
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
