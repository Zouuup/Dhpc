/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { AllowedOracles } from "./allowed_oracles";
import { MinerResponse } from "./miner_response";
import { Params } from "./params";
import { RequestRecord } from "./request_record";
import { Treasury } from "./treasury";

export const protobufPackage = "dhpc.request";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetAllowedOraclesRequest {
  id: number;
}

export interface QueryGetAllowedOraclesResponse {
  AllowedOracles: AllowedOracles | undefined;
}

export interface QueryAllAllowedOraclesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllAllowedOraclesResponse {
  AllowedOracles: AllowedOracles[];
  pagination: PageResponse | undefined;
}

export interface QueryGetMinerResponseRequest {
  uUID: string;
}

export interface QueryGetMinerResponseResponse {
  minerResponse: MinerResponse | undefined;
}

export interface QueryAllMinerResponseRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllMinerResponseResponse {
  minerResponse: MinerResponse[];
  pagination: PageResponse | undefined;
}

export interface QueryGetRequestRecordRequest {
  uUID: string;
}

export interface QueryGetRequestRecordResponse {
  requestRecord: RequestRecord | undefined;
}

export interface QueryAllRequestRecordRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRequestRecordResponse {
  requestRecord: RequestRecord[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTreasuryRequest {
}

export interface QueryGetTreasuryResponse {
  Treasury: Treasury | undefined;
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

function createBaseQueryGetAllowedOraclesRequest(): QueryGetAllowedOraclesRequest {
  return { id: 0 };
}

export const QueryGetAllowedOraclesRequest = {
  encode(message: QueryGetAllowedOraclesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAllowedOraclesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAllowedOraclesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAllowedOraclesRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetAllowedOraclesRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetAllowedOraclesRequest>, I>>(
    object: I,
  ): QueryGetAllowedOraclesRequest {
    const message = createBaseQueryGetAllowedOraclesRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryGetAllowedOraclesResponse(): QueryGetAllowedOraclesResponse {
  return { AllowedOracles: undefined };
}

export const QueryGetAllowedOraclesResponse = {
  encode(message: QueryGetAllowedOraclesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.AllowedOracles !== undefined) {
      AllowedOracles.encode(message.AllowedOracles, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAllowedOraclesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAllowedOraclesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.AllowedOracles = AllowedOracles.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAllowedOraclesResponse {
    return {
      AllowedOracles: isSet(object.AllowedOracles) ? AllowedOracles.fromJSON(object.AllowedOracles) : undefined,
    };
  },

  toJSON(message: QueryGetAllowedOraclesResponse): unknown {
    const obj: any = {};
    message.AllowedOracles !== undefined
      && (obj.AllowedOracles = message.AllowedOracles ? AllowedOracles.toJSON(message.AllowedOracles) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetAllowedOraclesResponse>, I>>(
    object: I,
  ): QueryGetAllowedOraclesResponse {
    const message = createBaseQueryGetAllowedOraclesResponse();
    message.AllowedOracles = (object.AllowedOracles !== undefined && object.AllowedOracles !== null)
      ? AllowedOracles.fromPartial(object.AllowedOracles)
      : undefined;
    return message;
  },
};

function createBaseQueryAllAllowedOraclesRequest(): QueryAllAllowedOraclesRequest {
  return { pagination: undefined };
}

export const QueryAllAllowedOraclesRequest = {
  encode(message: QueryAllAllowedOraclesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllAllowedOraclesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllAllowedOraclesRequest();
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

  fromJSON(object: any): QueryAllAllowedOraclesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllAllowedOraclesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllAllowedOraclesRequest>, I>>(
    object: I,
  ): QueryAllAllowedOraclesRequest {
    const message = createBaseQueryAllAllowedOraclesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllAllowedOraclesResponse(): QueryAllAllowedOraclesResponse {
  return { AllowedOracles: [], pagination: undefined };
}

export const QueryAllAllowedOraclesResponse = {
  encode(message: QueryAllAllowedOraclesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.AllowedOracles) {
      AllowedOracles.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllAllowedOraclesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllAllowedOraclesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.AllowedOracles.push(AllowedOracles.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllAllowedOraclesResponse {
    return {
      AllowedOracles: Array.isArray(object?.AllowedOracles)
        ? object.AllowedOracles.map((e: any) => AllowedOracles.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllAllowedOraclesResponse): unknown {
    const obj: any = {};
    if (message.AllowedOracles) {
      obj.AllowedOracles = message.AllowedOracles.map((e) => e ? AllowedOracles.toJSON(e) : undefined);
    } else {
      obj.AllowedOracles = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllAllowedOraclesResponse>, I>>(
    object: I,
  ): QueryAllAllowedOraclesResponse {
    const message = createBaseQueryAllAllowedOraclesResponse();
    message.AllowedOracles = object.AllowedOracles?.map((e) => AllowedOracles.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetMinerResponseRequest(): QueryGetMinerResponseRequest {
  return { uUID: "" };
}

export const QueryGetMinerResponseRequest = {
  encode(message: QueryGetMinerResponseRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uUID !== "") {
      writer.uint32(10).string(message.uUID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetMinerResponseRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetMinerResponseRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uUID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMinerResponseRequest {
    return { uUID: isSet(object.uUID) ? String(object.uUID) : "" };
  },

  toJSON(message: QueryGetMinerResponseRequest): unknown {
    const obj: any = {};
    message.uUID !== undefined && (obj.uUID = message.uUID);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetMinerResponseRequest>, I>>(object: I): QueryGetMinerResponseRequest {
    const message = createBaseQueryGetMinerResponseRequest();
    message.uUID = object.uUID ?? "";
    return message;
  },
};

function createBaseQueryGetMinerResponseResponse(): QueryGetMinerResponseResponse {
  return { minerResponse: undefined };
}

export const QueryGetMinerResponseResponse = {
  encode(message: QueryGetMinerResponseResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.minerResponse !== undefined) {
      MinerResponse.encode(message.minerResponse, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetMinerResponseResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetMinerResponseResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.minerResponse = MinerResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMinerResponseResponse {
    return { minerResponse: isSet(object.minerResponse) ? MinerResponse.fromJSON(object.minerResponse) : undefined };
  },

  toJSON(message: QueryGetMinerResponseResponse): unknown {
    const obj: any = {};
    message.minerResponse !== undefined
      && (obj.minerResponse = message.minerResponse ? MinerResponse.toJSON(message.minerResponse) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetMinerResponseResponse>, I>>(
    object: I,
  ): QueryGetMinerResponseResponse {
    const message = createBaseQueryGetMinerResponseResponse();
    message.minerResponse = (object.minerResponse !== undefined && object.minerResponse !== null)
      ? MinerResponse.fromPartial(object.minerResponse)
      : undefined;
    return message;
  },
};

function createBaseQueryAllMinerResponseRequest(): QueryAllMinerResponseRequest {
  return { pagination: undefined };
}

export const QueryAllMinerResponseRequest = {
  encode(message: QueryAllMinerResponseRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllMinerResponseRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllMinerResponseRequest();
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

  fromJSON(object: any): QueryAllMinerResponseRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllMinerResponseRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllMinerResponseRequest>, I>>(object: I): QueryAllMinerResponseRequest {
    const message = createBaseQueryAllMinerResponseRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllMinerResponseResponse(): QueryAllMinerResponseResponse {
  return { minerResponse: [], pagination: undefined };
}

export const QueryAllMinerResponseResponse = {
  encode(message: QueryAllMinerResponseResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.minerResponse) {
      MinerResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllMinerResponseResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllMinerResponseResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.minerResponse.push(MinerResponse.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllMinerResponseResponse {
    return {
      minerResponse: Array.isArray(object?.minerResponse)
        ? object.minerResponse.map((e: any) => MinerResponse.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllMinerResponseResponse): unknown {
    const obj: any = {};
    if (message.minerResponse) {
      obj.minerResponse = message.minerResponse.map((e) => e ? MinerResponse.toJSON(e) : undefined);
    } else {
      obj.minerResponse = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllMinerResponseResponse>, I>>(
    object: I,
  ): QueryAllMinerResponseResponse {
    const message = createBaseQueryAllMinerResponseResponse();
    message.minerResponse = object.minerResponse?.map((e) => MinerResponse.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetRequestRecordRequest(): QueryGetRequestRecordRequest {
  return { uUID: "" };
}

export const QueryGetRequestRecordRequest = {
  encode(message: QueryGetRequestRecordRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uUID !== "") {
      writer.uint32(10).string(message.uUID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRequestRecordRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRequestRecordRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uUID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRequestRecordRequest {
    return { uUID: isSet(object.uUID) ? String(object.uUID) : "" };
  },

  toJSON(message: QueryGetRequestRecordRequest): unknown {
    const obj: any = {};
    message.uUID !== undefined && (obj.uUID = message.uUID);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRequestRecordRequest>, I>>(object: I): QueryGetRequestRecordRequest {
    const message = createBaseQueryGetRequestRecordRequest();
    message.uUID = object.uUID ?? "";
    return message;
  },
};

function createBaseQueryGetRequestRecordResponse(): QueryGetRequestRecordResponse {
  return { requestRecord: undefined };
}

export const QueryGetRequestRecordResponse = {
  encode(message: QueryGetRequestRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.requestRecord !== undefined) {
      RequestRecord.encode(message.requestRecord, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRequestRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRequestRecordResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.requestRecord = RequestRecord.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRequestRecordResponse {
    return { requestRecord: isSet(object.requestRecord) ? RequestRecord.fromJSON(object.requestRecord) : undefined };
  },

  toJSON(message: QueryGetRequestRecordResponse): unknown {
    const obj: any = {};
    message.requestRecord !== undefined
      && (obj.requestRecord = message.requestRecord ? RequestRecord.toJSON(message.requestRecord) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRequestRecordResponse>, I>>(
    object: I,
  ): QueryGetRequestRecordResponse {
    const message = createBaseQueryGetRequestRecordResponse();
    message.requestRecord = (object.requestRecord !== undefined && object.requestRecord !== null)
      ? RequestRecord.fromPartial(object.requestRecord)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRequestRecordRequest(): QueryAllRequestRecordRequest {
  return { pagination: undefined };
}

export const QueryAllRequestRecordRequest = {
  encode(message: QueryAllRequestRecordRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRequestRecordRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRequestRecordRequest();
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

  fromJSON(object: any): QueryAllRequestRecordRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllRequestRecordRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRequestRecordRequest>, I>>(object: I): QueryAllRequestRecordRequest {
    const message = createBaseQueryAllRequestRecordRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRequestRecordResponse(): QueryAllRequestRecordResponse {
  return { requestRecord: [], pagination: undefined };
}

export const QueryAllRequestRecordResponse = {
  encode(message: QueryAllRequestRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.requestRecord) {
      RequestRecord.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRequestRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRequestRecordResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.requestRecord.push(RequestRecord.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllRequestRecordResponse {
    return {
      requestRecord: Array.isArray(object?.requestRecord)
        ? object.requestRecord.map((e: any) => RequestRecord.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllRequestRecordResponse): unknown {
    const obj: any = {};
    if (message.requestRecord) {
      obj.requestRecord = message.requestRecord.map((e) => e ? RequestRecord.toJSON(e) : undefined);
    } else {
      obj.requestRecord = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRequestRecordResponse>, I>>(
    object: I,
  ): QueryAllRequestRecordResponse {
    const message = createBaseQueryAllRequestRecordResponse();
    message.requestRecord = object.requestRecord?.map((e) => RequestRecord.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetTreasuryRequest(): QueryGetTreasuryRequest {
  return {};
}

export const QueryGetTreasuryRequest = {
  encode(_: QueryGetTreasuryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTreasuryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTreasuryRequest();
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

  fromJSON(_: any): QueryGetTreasuryRequest {
    return {};
  },

  toJSON(_: QueryGetTreasuryRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetTreasuryRequest>, I>>(_: I): QueryGetTreasuryRequest {
    const message = createBaseQueryGetTreasuryRequest();
    return message;
  },
};

function createBaseQueryGetTreasuryResponse(): QueryGetTreasuryResponse {
  return { Treasury: undefined };
}

export const QueryGetTreasuryResponse = {
  encode(message: QueryGetTreasuryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.Treasury !== undefined) {
      Treasury.encode(message.Treasury, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTreasuryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTreasuryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Treasury = Treasury.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTreasuryResponse {
    return { Treasury: isSet(object.Treasury) ? Treasury.fromJSON(object.Treasury) : undefined };
  },

  toJSON(message: QueryGetTreasuryResponse): unknown {
    const obj: any = {};
    message.Treasury !== undefined && (obj.Treasury = message.Treasury ? Treasury.toJSON(message.Treasury) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetTreasuryResponse>, I>>(object: I): QueryGetTreasuryResponse {
    const message = createBaseQueryGetTreasuryResponse();
    message.Treasury = (object.Treasury !== undefined && object.Treasury !== null)
      ? Treasury.fromPartial(object.Treasury)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of AllowedOracles items. */
  AllowedOracles(request: QueryGetAllowedOraclesRequest): Promise<QueryGetAllowedOraclesResponse>;
  AllowedOraclesAll(request: QueryAllAllowedOraclesRequest): Promise<QueryAllAllowedOraclesResponse>;
  /** Queries a list of MinerResponse items. */
  MinerResponse(request: QueryGetMinerResponseRequest): Promise<QueryGetMinerResponseResponse>;
  MinerResponseAll(request: QueryAllMinerResponseRequest): Promise<QueryAllMinerResponseResponse>;
  /** Queries a list of RequestRecord items. */
  RequestRecord(request: QueryGetRequestRecordRequest): Promise<QueryGetRequestRecordResponse>;
  RequestRecordAll(request: QueryAllRequestRecordRequest): Promise<QueryAllRequestRecordResponse>;
  RequestRecordAllMinerPending(request: QueryAllRequestRecordRequest): Promise<QueryAllRequestRecordResponse>;
  RequestRecordAllAnswerPending(request: QueryAllRequestRecordRequest): Promise<QueryAllRequestRecordResponse>;
  /** Queries a Treasury by index. */
  Treasury(request: QueryGetTreasuryRequest): Promise<QueryGetTreasuryResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.AllowedOracles = this.AllowedOracles.bind(this);
    this.AllowedOraclesAll = this.AllowedOraclesAll.bind(this);
    this.MinerResponse = this.MinerResponse.bind(this);
    this.MinerResponseAll = this.MinerResponseAll.bind(this);
    this.RequestRecord = this.RequestRecord.bind(this);
    this.RequestRecordAll = this.RequestRecordAll.bind(this);
    this.RequestRecordAllMinerPending = this.RequestRecordAllMinerPending.bind(this);
    this.RequestRecordAllAnswerPending = this.RequestRecordAllAnswerPending.bind(this);
    this.Treasury = this.Treasury.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  AllowedOracles(request: QueryGetAllowedOraclesRequest): Promise<QueryGetAllowedOraclesResponse> {
    const data = QueryGetAllowedOraclesRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Query", "AllowedOracles", data);
    return promise.then((data) => QueryGetAllowedOraclesResponse.decode(new _m0.Reader(data)));
  }

  AllowedOraclesAll(request: QueryAllAllowedOraclesRequest): Promise<QueryAllAllowedOraclesResponse> {
    const data = QueryAllAllowedOraclesRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Query", "AllowedOraclesAll", data);
    return promise.then((data) => QueryAllAllowedOraclesResponse.decode(new _m0.Reader(data)));
  }

  MinerResponse(request: QueryGetMinerResponseRequest): Promise<QueryGetMinerResponseResponse> {
    const data = QueryGetMinerResponseRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Query", "MinerResponse", data);
    return promise.then((data) => QueryGetMinerResponseResponse.decode(new _m0.Reader(data)));
  }

  MinerResponseAll(request: QueryAllMinerResponseRequest): Promise<QueryAllMinerResponseResponse> {
    const data = QueryAllMinerResponseRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Query", "MinerResponseAll", data);
    return promise.then((data) => QueryAllMinerResponseResponse.decode(new _m0.Reader(data)));
  }

  RequestRecord(request: QueryGetRequestRecordRequest): Promise<QueryGetRequestRecordResponse> {
    const data = QueryGetRequestRecordRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Query", "RequestRecord", data);
    return promise.then((data) => QueryGetRequestRecordResponse.decode(new _m0.Reader(data)));
  }

  RequestRecordAll(request: QueryAllRequestRecordRequest): Promise<QueryAllRequestRecordResponse> {
    const data = QueryAllRequestRecordRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Query", "RequestRecordAll", data);
    return promise.then((data) => QueryAllRequestRecordResponse.decode(new _m0.Reader(data)));
  }

  RequestRecordAllMinerPending(request: QueryAllRequestRecordRequest): Promise<QueryAllRequestRecordResponse> {
    const data = QueryAllRequestRecordRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Query", "RequestRecordAllMinerPending", data);
    return promise.then((data) => QueryAllRequestRecordResponse.decode(new _m0.Reader(data)));
  }

  RequestRecordAllAnswerPending(request: QueryAllRequestRecordRequest): Promise<QueryAllRequestRecordResponse> {
    const data = QueryAllRequestRecordRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Query", "RequestRecordAllAnswerPending", data);
    return promise.then((data) => QueryAllRequestRecordResponse.decode(new _m0.Reader(data)));
  }

  Treasury(request: QueryGetTreasuryRequest): Promise<QueryGetTreasuryResponse> {
    const data = QueryGetTreasuryRequest.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Query", "Treasury", data);
    return promise.then((data) => QueryGetTreasuryResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
