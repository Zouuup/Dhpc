/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dhpc.data";

export interface MsgAddData {
  creator: string;
  address: string;
  owner: string;
  network: string;
  event: string;
  blockvalidity: number;
  score: number;
  hash: string;
}

export interface MsgAddDataResponse {
}

export interface MsgDeleteData {
  creator: string;
  address: string;
  owner: string;
  network: string;
}

export interface MsgDeleteDataResponse {
}

function createBaseMsgAddData(): MsgAddData {
  return { creator: "", address: "", owner: "", network: "", event: "", blockvalidity: 0, score: 0, hash: "" };
}

export const MsgAddData = {
  encode(message: MsgAddData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    if (message.network !== "") {
      writer.uint32(34).string(message.network);
    }
    if (message.event !== "") {
      writer.uint32(42).string(message.event);
    }
    if (message.blockvalidity !== 0) {
      writer.uint32(48).uint64(message.blockvalidity);
    }
    if (message.score !== 0) {
      writer.uint32(56).int32(message.score);
    }
    if (message.hash !== "") {
      writer.uint32(66).string(message.hash);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.owner = reader.string();
          break;
        case 4:
          message.network = reader.string();
          break;
        case 5:
          message.event = reader.string();
          break;
        case 6:
          message.blockvalidity = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.score = reader.int32();
          break;
        case 8:
          message.hash = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddData {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
      network: isSet(object.network) ? String(object.network) : "",
      event: isSet(object.event) ? String(object.event) : "",
      blockvalidity: isSet(object.blockvalidity) ? Number(object.blockvalidity) : 0,
      score: isSet(object.score) ? Number(object.score) : 0,
      hash: isSet(object.hash) ? String(object.hash) : "",
    };
  },

  toJSON(message: MsgAddData): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.owner !== undefined && (obj.owner = message.owner);
    message.network !== undefined && (obj.network = message.network);
    message.event !== undefined && (obj.event = message.event);
    message.blockvalidity !== undefined && (obj.blockvalidity = Math.round(message.blockvalidity));
    message.score !== undefined && (obj.score = Math.round(message.score));
    message.hash !== undefined && (obj.hash = message.hash);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddData>, I>>(object: I): MsgAddData {
    const message = createBaseMsgAddData();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    message.owner = object.owner ?? "";
    message.network = object.network ?? "";
    message.event = object.event ?? "";
    message.blockvalidity = object.blockvalidity ?? 0;
    message.score = object.score ?? 0;
    message.hash = object.hash ?? "";
    return message;
  },
};

function createBaseMsgAddDataResponse(): MsgAddDataResponse {
  return {};
}

export const MsgAddDataResponse = {
  encode(_: MsgAddDataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddDataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddDataResponse();
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

  fromJSON(_: any): MsgAddDataResponse {
    return {};
  },

  toJSON(_: MsgAddDataResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddDataResponse>, I>>(_: I): MsgAddDataResponse {
    const message = createBaseMsgAddDataResponse();
    return message;
  },
};

function createBaseMsgDeleteData(): MsgDeleteData {
  return { creator: "", address: "", owner: "", network: "" };
}

export const MsgDeleteData = {
  encode(message: MsgDeleteData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    if (message.network !== "") {
      writer.uint32(34).string(message.network);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.owner = reader.string();
          break;
        case 4:
          message.network = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteData {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
      network: isSet(object.network) ? String(object.network) : "",
    };
  },

  toJSON(message: MsgDeleteData): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.owner !== undefined && (obj.owner = message.owner);
    message.network !== undefined && (obj.network = message.network);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteData>, I>>(object: I): MsgDeleteData {
    const message = createBaseMsgDeleteData();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    message.owner = object.owner ?? "";
    message.network = object.network ?? "";
    return message;
  },
};

function createBaseMsgDeleteDataResponse(): MsgDeleteDataResponse {
  return {};
}

export const MsgDeleteDataResponse = {
  encode(_: MsgDeleteDataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDataResponse();
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

  fromJSON(_: any): MsgDeleteDataResponse {
    return {};
  },

  toJSON(_: MsgDeleteDataResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDataResponse>, I>>(_: I): MsgDeleteDataResponse {
    const message = createBaseMsgDeleteDataResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  AddData(request: MsgAddData): Promise<MsgAddDataResponse>;
  DeleteData(request: MsgDeleteData): Promise<MsgDeleteDataResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.AddData = this.AddData.bind(this);
    this.DeleteData = this.DeleteData.bind(this);
  }
  AddData(request: MsgAddData): Promise<MsgAddDataResponse> {
    const data = MsgAddData.encode(request).finish();
    const promise = this.rpc.request("dhpc.data.Msg", "AddData", data);
    return promise.then((data) => MsgAddDataResponse.decode(new _m0.Reader(data)));
  }

  DeleteData(request: MsgDeleteData): Promise<MsgDeleteDataResponse> {
    const data = MsgDeleteData.encode(request).finish();
    const promise = this.rpc.request("dhpc.data.Msg", "DeleteData", data);
    return promise.then((data) => MsgDeleteDataResponse.decode(new _m0.Reader(data)));
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
