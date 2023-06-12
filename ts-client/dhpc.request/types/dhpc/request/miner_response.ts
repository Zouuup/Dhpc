/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dhpc.request";

export interface MinerResponse {
  uUID: string;
  requestUUID: string;
  hash: string;
  answer: number;
  creator: string;
  dataUsed: string;
  salt: number;
}

function createBaseMinerResponse(): MinerResponse {
  return { uUID: "", requestUUID: "", hash: "", answer: 0, creator: "", dataUsed: "", salt: 0 };
}

export const MinerResponse = {
  encode(message: MinerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uUID !== "") {
      writer.uint32(10).string(message.uUID);
    }
    if (message.requestUUID !== "") {
      writer.uint32(18).string(message.requestUUID);
    }
    if (message.hash !== "") {
      writer.uint32(26).string(message.hash);
    }
    if (message.answer !== 0) {
      writer.uint32(32).int32(message.answer);
    }
    if (message.creator !== "") {
      writer.uint32(42).string(message.creator);
    }
    if (message.dataUsed !== "") {
      writer.uint32(50).string(message.dataUsed);
    }
    if (message.salt !== 0) {
      writer.uint32(56).int32(message.salt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MinerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMinerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uUID = reader.string();
          break;
        case 2:
          message.requestUUID = reader.string();
          break;
        case 3:
          message.hash = reader.string();
          break;
        case 4:
          message.answer = reader.int32();
          break;
        case 5:
          message.creator = reader.string();
          break;
        case 6:
          message.dataUsed = reader.string();
          break;
        case 7:
          message.salt = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MinerResponse {
    return {
      uUID: isSet(object.uUID) ? String(object.uUID) : "",
      requestUUID: isSet(object.requestUUID) ? String(object.requestUUID) : "",
      hash: isSet(object.hash) ? String(object.hash) : "",
      answer: isSet(object.answer) ? Number(object.answer) : 0,
      creator: isSet(object.creator) ? String(object.creator) : "",
      dataUsed: isSet(object.dataUsed) ? String(object.dataUsed) : "",
      salt: isSet(object.salt) ? Number(object.salt) : 0,
    };
  },

  toJSON(message: MinerResponse): unknown {
    const obj: any = {};
    message.uUID !== undefined && (obj.uUID = message.uUID);
    message.requestUUID !== undefined && (obj.requestUUID = message.requestUUID);
    message.hash !== undefined && (obj.hash = message.hash);
    message.answer !== undefined && (obj.answer = Math.round(message.answer));
    message.creator !== undefined && (obj.creator = message.creator);
    message.dataUsed !== undefined && (obj.dataUsed = message.dataUsed);
    message.salt !== undefined && (obj.salt = Math.round(message.salt));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MinerResponse>, I>>(object: I): MinerResponse {
    const message = createBaseMinerResponse();
    message.uUID = object.uUID ?? "";
    message.requestUUID = object.requestUUID ?? "";
    message.hash = object.hash ?? "";
    message.answer = object.answer ?? 0;
    message.creator = object.creator ?? "";
    message.dataUsed = object.dataUsed ?? "";
    message.salt = object.salt ?? 0;
    return message;
  },
};

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
