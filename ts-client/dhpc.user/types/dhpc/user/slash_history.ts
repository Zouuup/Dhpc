/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dhpc.user";

export interface SlashHistory {
  index: string;
  datetime: number;
  amount: number;
  level: number;
}

function createBaseSlashHistory(): SlashHistory {
  return { index: "", datetime: 0, amount: 0, level: 0 };
}

export const SlashHistory = {
  encode(message: SlashHistory, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.datetime !== 0) {
      writer.uint32(16).int32(message.datetime);
    }
    if (message.amount !== 0) {
      writer.uint32(24).int32(message.amount);
    }
    if (message.level !== 0) {
      writer.uint32(32).int32(message.level);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SlashHistory {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSlashHistory();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.datetime = reader.int32();
          break;
        case 3:
          message.amount = reader.int32();
          break;
        case 4:
          message.level = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SlashHistory {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      datetime: isSet(object.datetime) ? Number(object.datetime) : 0,
      amount: isSet(object.amount) ? Number(object.amount) : 0,
      level: isSet(object.level) ? Number(object.level) : 0,
    };
  },

  toJSON(message: SlashHistory): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.datetime !== undefined && (obj.datetime = Math.round(message.datetime));
    message.amount !== undefined && (obj.amount = Math.round(message.amount));
    message.level !== undefined && (obj.level = Math.round(message.level));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SlashHistory>, I>>(object: I): SlashHistory {
    const message = createBaseSlashHistory();
    message.index = object.index ?? "";
    message.datetime = object.datetime ?? 0;
    message.amount = object.amount ?? 0;
    message.level = object.level ?? 0;
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
