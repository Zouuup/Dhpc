/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { LinkedRequesters } from "./linked_requesters";
import { SlashHistory } from "./slash_history";

export const protobufPackage = "dhpc.user";

export interface User {
  account: string;
  reputation: number;
  deposit: Coin[];
  linkedRequester: LinkedRequesters[];
  slashHistory: SlashHistory[];
}

function createBaseUser(): User {
  return { account: "", reputation: 0, deposit: [], linkedRequester: [], slashHistory: [] };
}

export const User = {
  encode(message: User, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== "") {
      writer.uint32(10).string(message.account);
    }
    if (message.reputation !== 0) {
      writer.uint32(16).int32(message.reputation);
    }
    for (const v of message.deposit) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.linkedRequester) {
      LinkedRequesters.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.slashHistory) {
      SlashHistory.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): User {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.account = reader.string();
          break;
        case 2:
          message.reputation = reader.int32();
          break;
        case 3:
          message.deposit.push(Coin.decode(reader, reader.uint32()));
          break;
        case 4:
          message.linkedRequester.push(LinkedRequesters.decode(reader, reader.uint32()));
          break;
        case 5:
          message.slashHistory.push(SlashHistory.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): User {
    return {
      account: isSet(object.account) ? String(object.account) : "",
      reputation: isSet(object.reputation) ? Number(object.reputation) : 0,
      deposit: Array.isArray(object?.deposit) ? object.deposit.map((e: any) => Coin.fromJSON(e)) : [],
      linkedRequester: Array.isArray(object?.linkedRequester)
        ? object.linkedRequester.map((e: any) => LinkedRequesters.fromJSON(e))
        : [],
      slashHistory: Array.isArray(object?.slashHistory)
        ? object.slashHistory.map((e: any) => SlashHistory.fromJSON(e))
        : [],
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    message.account !== undefined && (obj.account = message.account);
    message.reputation !== undefined && (obj.reputation = Math.round(message.reputation));
    if (message.deposit) {
      obj.deposit = message.deposit.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.deposit = [];
    }
    if (message.linkedRequester) {
      obj.linkedRequester = message.linkedRequester.map((e) => e ? LinkedRequesters.toJSON(e) : undefined);
    } else {
      obj.linkedRequester = [];
    }
    if (message.slashHistory) {
      obj.slashHistory = message.slashHistory.map((e) => e ? SlashHistory.toJSON(e) : undefined);
    } else {
      obj.slashHistory = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<User>, I>>(object: I): User {
    const message = createBaseUser();
    message.account = object.account ?? "";
    message.reputation = object.reputation ?? 0;
    message.deposit = object.deposit?.map((e) => Coin.fromPartial(e)) || [];
    message.linkedRequester = object.linkedRequester?.map((e) => LinkedRequesters.fromPartial(e)) || [];
    message.slashHistory = object.slashHistory?.map((e) => SlashHistory.fromPartial(e)) || [];
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
