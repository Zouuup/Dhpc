/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dhpc.data";

export interface Data {
  address: string;
  owner: string;
  network: string;
  dateAdded: number;
  dateUpdated: number;
  blockValidity: number;
  event: string;
  score: number;
  hit: number;
  hash: string;
}

function createBaseData(): Data {
  return {
    address: "",
    owner: "",
    network: "",
    dateAdded: 0,
    dateUpdated: 0,
    blockValidity: 0,
    event: "",
    score: 0,
    hit: 0,
    hash: "",
  };
}

export const Data = {
  encode(message: Data, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.network !== "") {
      writer.uint32(26).string(message.network);
    }
    if (message.dateAdded !== 0) {
      writer.uint32(32).uint64(message.dateAdded);
    }
    if (message.dateUpdated !== 0) {
      writer.uint32(40).uint64(message.dateUpdated);
    }
    if (message.blockValidity !== 0) {
      writer.uint32(48).uint64(message.blockValidity);
    }
    if (message.event !== "") {
      writer.uint32(58).string(message.event);
    }
    if (message.score !== 0) {
      writer.uint32(64).int32(message.score);
    }
    if (message.hit !== 0) {
      writer.uint32(72).uint64(message.hit);
    }
    if (message.hash !== "") {
      writer.uint32(82).string(message.hash);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Data {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseData();
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
        case 4:
          message.dateAdded = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.dateUpdated = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.blockValidity = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.event = reader.string();
          break;
        case 8:
          message.score = reader.int32();
          break;
        case 9:
          message.hit = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.hash = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Data {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
      network: isSet(object.network) ? String(object.network) : "",
      dateAdded: isSet(object.dateAdded) ? Number(object.dateAdded) : 0,
      dateUpdated: isSet(object.dateUpdated) ? Number(object.dateUpdated) : 0,
      blockValidity: isSet(object.blockValidity) ? Number(object.blockValidity) : 0,
      event: isSet(object.event) ? String(object.event) : "",
      score: isSet(object.score) ? Number(object.score) : 0,
      hit: isSet(object.hit) ? Number(object.hit) : 0,
      hash: isSet(object.hash) ? String(object.hash) : "",
    };
  },

  toJSON(message: Data): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.owner !== undefined && (obj.owner = message.owner);
    message.network !== undefined && (obj.network = message.network);
    message.dateAdded !== undefined && (obj.dateAdded = Math.round(message.dateAdded));
    message.dateUpdated !== undefined && (obj.dateUpdated = Math.round(message.dateUpdated));
    message.blockValidity !== undefined && (obj.blockValidity = Math.round(message.blockValidity));
    message.event !== undefined && (obj.event = message.event);
    message.score !== undefined && (obj.score = Math.round(message.score));
    message.hit !== undefined && (obj.hit = Math.round(message.hit));
    message.hash !== undefined && (obj.hash = message.hash);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Data>, I>>(object: I): Data {
    const message = createBaseData();
    message.address = object.address ?? "";
    message.owner = object.owner ?? "";
    message.network = object.network ?? "";
    message.dateAdded = object.dateAdded ?? 0;
    message.dateUpdated = object.dateUpdated ?? 0;
    message.blockValidity = object.blockValidity ?? 0;
    message.event = object.event ?? "";
    message.score = object.score ?? 0;
    message.hit = object.hit ?? 0;
    message.hash = object.hash ?? "";
    return message;
  },
};

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
