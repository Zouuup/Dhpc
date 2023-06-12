/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { MinerResponse } from "./miner_response";

export const protobufPackage = "dhpc.request";

export interface RequestRecord {
  uUID: string;
  tXhash: string;
  network: string;
  from: string;
  address: string;
  score: number;
  oracle: string;
  block: number;
  stage: number;
  miners: MinerResponse[];
  creator: string;
  createdBlock: number;
  updatedBlock: number;
}

function createBaseRequestRecord(): RequestRecord {
  return {
    uUID: "",
    tXhash: "",
    network: "",
    from: "",
    address: "",
    score: 0,
    oracle: "",
    block: 0,
    stage: 0,
    miners: [],
    creator: "",
    createdBlock: 0,
    updatedBlock: 0,
  };
}

export const RequestRecord = {
  encode(message: RequestRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uUID !== "") {
      writer.uint32(10).string(message.uUID);
    }
    if (message.tXhash !== "") {
      writer.uint32(18).string(message.tXhash);
    }
    if (message.network !== "") {
      writer.uint32(26).string(message.network);
    }
    if (message.from !== "") {
      writer.uint32(34).string(message.from);
    }
    if (message.address !== "") {
      writer.uint32(42).string(message.address);
    }
    if (message.score !== 0) {
      writer.uint32(48).int32(message.score);
    }
    if (message.oracle !== "") {
      writer.uint32(58).string(message.oracle);
    }
    if (message.block !== 0) {
      writer.uint32(64).uint64(message.block);
    }
    if (message.stage !== 0) {
      writer.uint32(72).int32(message.stage);
    }
    for (const v of message.miners) {
      MinerResponse.encode(v!, writer.uint32(82).fork()).ldelim();
    }
    if (message.creator !== "") {
      writer.uint32(90).string(message.creator);
    }
    if (message.createdBlock !== 0) {
      writer.uint32(96).int64(message.createdBlock);
    }
    if (message.updatedBlock !== 0) {
      writer.uint32(104).int64(message.updatedBlock);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RequestRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRequestRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uUID = reader.string();
          break;
        case 2:
          message.tXhash = reader.string();
          break;
        case 3:
          message.network = reader.string();
          break;
        case 4:
          message.from = reader.string();
          break;
        case 5:
          message.address = reader.string();
          break;
        case 6:
          message.score = reader.int32();
          break;
        case 7:
          message.oracle = reader.string();
          break;
        case 8:
          message.block = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.stage = reader.int32();
          break;
        case 10:
          message.miners.push(MinerResponse.decode(reader, reader.uint32()));
          break;
        case 11:
          message.creator = reader.string();
          break;
        case 12:
          message.createdBlock = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.updatedBlock = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RequestRecord {
    return {
      uUID: isSet(object.uUID) ? String(object.uUID) : "",
      tXhash: isSet(object.tXhash) ? String(object.tXhash) : "",
      network: isSet(object.network) ? String(object.network) : "",
      from: isSet(object.from) ? String(object.from) : "",
      address: isSet(object.address) ? String(object.address) : "",
      score: isSet(object.score) ? Number(object.score) : 0,
      oracle: isSet(object.oracle) ? String(object.oracle) : "",
      block: isSet(object.block) ? Number(object.block) : 0,
      stage: isSet(object.stage) ? Number(object.stage) : 0,
      miners: Array.isArray(object?.miners) ? object.miners.map((e: any) => MinerResponse.fromJSON(e)) : [],
      creator: isSet(object.creator) ? String(object.creator) : "",
      createdBlock: isSet(object.createdBlock) ? Number(object.createdBlock) : 0,
      updatedBlock: isSet(object.updatedBlock) ? Number(object.updatedBlock) : 0,
    };
  },

  toJSON(message: RequestRecord): unknown {
    const obj: any = {};
    message.uUID !== undefined && (obj.uUID = message.uUID);
    message.tXhash !== undefined && (obj.tXhash = message.tXhash);
    message.network !== undefined && (obj.network = message.network);
    message.from !== undefined && (obj.from = message.from);
    message.address !== undefined && (obj.address = message.address);
    message.score !== undefined && (obj.score = Math.round(message.score));
    message.oracle !== undefined && (obj.oracle = message.oracle);
    message.block !== undefined && (obj.block = Math.round(message.block));
    message.stage !== undefined && (obj.stage = Math.round(message.stage));
    if (message.miners) {
      obj.miners = message.miners.map((e) => e ? MinerResponse.toJSON(e) : undefined);
    } else {
      obj.miners = [];
    }
    message.creator !== undefined && (obj.creator = message.creator);
    message.createdBlock !== undefined && (obj.createdBlock = Math.round(message.createdBlock));
    message.updatedBlock !== undefined && (obj.updatedBlock = Math.round(message.updatedBlock));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RequestRecord>, I>>(object: I): RequestRecord {
    const message = createBaseRequestRecord();
    message.uUID = object.uUID ?? "";
    message.tXhash = object.tXhash ?? "";
    message.network = object.network ?? "";
    message.from = object.from ?? "";
    message.address = object.address ?? "";
    message.score = object.score ?? 0;
    message.oracle = object.oracle ?? "";
    message.block = object.block ?? 0;
    message.stage = object.stage ?? 0;
    message.miners = object.miners?.map((e) => MinerResponse.fromPartial(e)) || [];
    message.creator = object.creator ?? "";
    message.createdBlock = object.createdBlock ?? 0;
    message.updatedBlock = object.updatedBlock ?? 0;
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
