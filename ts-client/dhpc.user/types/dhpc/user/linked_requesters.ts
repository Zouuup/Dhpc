/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dhpc.user";

export interface LinkedRequesters {
  index: string;
  network: string;
  address: string;
}

function createBaseLinkedRequesters(): LinkedRequesters {
  return { index: "", network: "", address: "" };
}

export const LinkedRequesters = {
  encode(message: LinkedRequesters, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.network !== "") {
      writer.uint32(18).string(message.network);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LinkedRequesters {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLinkedRequesters();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.network = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LinkedRequesters {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      network: isSet(object.network) ? String(object.network) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: LinkedRequesters): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.network !== undefined && (obj.network = message.network);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LinkedRequesters>, I>>(object: I): LinkedRequesters {
    const message = createBaseLinkedRequesters();
    message.index = object.index ?? "";
    message.network = object.network ?? "";
    message.address = object.address ?? "";
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
