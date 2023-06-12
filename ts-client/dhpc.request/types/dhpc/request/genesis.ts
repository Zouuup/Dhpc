/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { AllowedOracles } from "./allowed_oracles";
import { MinerResponse } from "./miner_response";
import { Params } from "./params";
import { RequestRecord } from "./request_record";
import { Treasury } from "./treasury";

export const protobufPackage = "dhpc.request";

/** GenesisState defines the request module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  allowedOraclesList: AllowedOracles[];
  allowedOraclesCount: number;
  minerResponseList: MinerResponse[];
  requestRecordList: RequestRecord[];
  treasury: Treasury | undefined;
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    allowedOraclesList: [],
    allowedOraclesCount: 0,
    minerResponseList: [],
    requestRecordList: [],
    treasury: undefined,
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.allowedOraclesList) {
      AllowedOracles.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.allowedOraclesCount !== 0) {
      writer.uint32(24).uint64(message.allowedOraclesCount);
    }
    for (const v of message.minerResponseList) {
      MinerResponse.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.requestRecordList) {
      RequestRecord.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.treasury !== undefined) {
      Treasury.encode(message.treasury, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.allowedOraclesList.push(AllowedOracles.decode(reader, reader.uint32()));
          break;
        case 3:
          message.allowedOraclesCount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.minerResponseList.push(MinerResponse.decode(reader, reader.uint32()));
          break;
        case 5:
          message.requestRecordList.push(RequestRecord.decode(reader, reader.uint32()));
          break;
        case 6:
          message.treasury = Treasury.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      allowedOraclesList: Array.isArray(object?.allowedOraclesList)
        ? object.allowedOraclesList.map((e: any) => AllowedOracles.fromJSON(e))
        : [],
      allowedOraclesCount: isSet(object.allowedOraclesCount) ? Number(object.allowedOraclesCount) : 0,
      minerResponseList: Array.isArray(object?.minerResponseList)
        ? object.minerResponseList.map((e: any) => MinerResponse.fromJSON(e))
        : [],
      requestRecordList: Array.isArray(object?.requestRecordList)
        ? object.requestRecordList.map((e: any) => RequestRecord.fromJSON(e))
        : [],
      treasury: isSet(object.treasury) ? Treasury.fromJSON(object.treasury) : undefined,
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.allowedOraclesList) {
      obj.allowedOraclesList = message.allowedOraclesList.map((e) => e ? AllowedOracles.toJSON(e) : undefined);
    } else {
      obj.allowedOraclesList = [];
    }
    message.allowedOraclesCount !== undefined && (obj.allowedOraclesCount = Math.round(message.allowedOraclesCount));
    if (message.minerResponseList) {
      obj.minerResponseList = message.minerResponseList.map((e) => e ? MinerResponse.toJSON(e) : undefined);
    } else {
      obj.minerResponseList = [];
    }
    if (message.requestRecordList) {
      obj.requestRecordList = message.requestRecordList.map((e) => e ? RequestRecord.toJSON(e) : undefined);
    } else {
      obj.requestRecordList = [];
    }
    message.treasury !== undefined && (obj.treasury = message.treasury ? Treasury.toJSON(message.treasury) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.allowedOraclesList = object.allowedOraclesList?.map((e) => AllowedOracles.fromPartial(e)) || [];
    message.allowedOraclesCount = object.allowedOraclesCount ?? 0;
    message.minerResponseList = object.minerResponseList?.map((e) => MinerResponse.fromPartial(e)) || [];
    message.requestRecordList = object.requestRecordList?.map((e) => RequestRecord.fromPartial(e)) || [];
    message.treasury = (object.treasury !== undefined && object.treasury !== null)
      ? Treasury.fromPartial(object.treasury)
      : undefined;
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
