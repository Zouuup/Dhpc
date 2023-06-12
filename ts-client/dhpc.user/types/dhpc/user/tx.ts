/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dhpc.user";

export interface MsgDepositToken {
  creator: string;
  deposit: string;
}

export interface MsgDepositTokenResponse {
}

export interface MsgWithdrawToken {
  creator: string;
  withdraw: string;
}

export interface MsgWithdrawTokenResponse {
}

export interface MsgAddLinkedRequester {
  creator: string;
  network: string;
  address: string;
}

export interface MsgAddLinkedRequesterResponse {
}

export interface MsgRemoveLinkedRequester {
  creator: string;
  network: string;
  address: string;
}

export interface MsgRemoveLinkedRequesterResponse {
}

function createBaseMsgDepositToken(): MsgDepositToken {
  return { creator: "", deposit: "" };
}

export const MsgDepositToken = {
  encode(message: MsgDepositToken, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.deposit !== "") {
      writer.uint32(18).string(message.deposit);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDepositToken {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDepositToken();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.deposit = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDepositToken {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      deposit: isSet(object.deposit) ? String(object.deposit) : "",
    };
  },

  toJSON(message: MsgDepositToken): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.deposit !== undefined && (obj.deposit = message.deposit);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDepositToken>, I>>(object: I): MsgDepositToken {
    const message = createBaseMsgDepositToken();
    message.creator = object.creator ?? "";
    message.deposit = object.deposit ?? "";
    return message;
  },
};

function createBaseMsgDepositTokenResponse(): MsgDepositTokenResponse {
  return {};
}

export const MsgDepositTokenResponse = {
  encode(_: MsgDepositTokenResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDepositTokenResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDepositTokenResponse();
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

  fromJSON(_: any): MsgDepositTokenResponse {
    return {};
  },

  toJSON(_: MsgDepositTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDepositTokenResponse>, I>>(_: I): MsgDepositTokenResponse {
    const message = createBaseMsgDepositTokenResponse();
    return message;
  },
};

function createBaseMsgWithdrawToken(): MsgWithdrawToken {
  return { creator: "", withdraw: "" };
}

export const MsgWithdrawToken = {
  encode(message: MsgWithdrawToken, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.withdraw !== "") {
      writer.uint32(18).string(message.withdraw);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgWithdrawToken {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawToken();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.withdraw = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdrawToken {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      withdraw: isSet(object.withdraw) ? String(object.withdraw) : "",
    };
  },

  toJSON(message: MsgWithdrawToken): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.withdraw !== undefined && (obj.withdraw = message.withdraw);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgWithdrawToken>, I>>(object: I): MsgWithdrawToken {
    const message = createBaseMsgWithdrawToken();
    message.creator = object.creator ?? "";
    message.withdraw = object.withdraw ?? "";
    return message;
  },
};

function createBaseMsgWithdrawTokenResponse(): MsgWithdrawTokenResponse {
  return {};
}

export const MsgWithdrawTokenResponse = {
  encode(_: MsgWithdrawTokenResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgWithdrawTokenResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawTokenResponse();
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

  fromJSON(_: any): MsgWithdrawTokenResponse {
    return {};
  },

  toJSON(_: MsgWithdrawTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgWithdrawTokenResponse>, I>>(_: I): MsgWithdrawTokenResponse {
    const message = createBaseMsgWithdrawTokenResponse();
    return message;
  },
};

function createBaseMsgAddLinkedRequester(): MsgAddLinkedRequester {
  return { creator: "", network: "", address: "" };
}

export const MsgAddLinkedRequester = {
  encode(message: MsgAddLinkedRequester, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.network !== "") {
      writer.uint32(18).string(message.network);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddLinkedRequester {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddLinkedRequester();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
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

  fromJSON(object: any): MsgAddLinkedRequester {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      network: isSet(object.network) ? String(object.network) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgAddLinkedRequester): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.network !== undefined && (obj.network = message.network);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddLinkedRequester>, I>>(object: I): MsgAddLinkedRequester {
    const message = createBaseMsgAddLinkedRequester();
    message.creator = object.creator ?? "";
    message.network = object.network ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgAddLinkedRequesterResponse(): MsgAddLinkedRequesterResponse {
  return {};
}

export const MsgAddLinkedRequesterResponse = {
  encode(_: MsgAddLinkedRequesterResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddLinkedRequesterResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddLinkedRequesterResponse();
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

  fromJSON(_: any): MsgAddLinkedRequesterResponse {
    return {};
  },

  toJSON(_: MsgAddLinkedRequesterResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddLinkedRequesterResponse>, I>>(_: I): MsgAddLinkedRequesterResponse {
    const message = createBaseMsgAddLinkedRequesterResponse();
    return message;
  },
};

function createBaseMsgRemoveLinkedRequester(): MsgRemoveLinkedRequester {
  return { creator: "", network: "", address: "" };
}

export const MsgRemoveLinkedRequester = {
  encode(message: MsgRemoveLinkedRequester, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.network !== "") {
      writer.uint32(18).string(message.network);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemoveLinkedRequester {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveLinkedRequester();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
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

  fromJSON(object: any): MsgRemoveLinkedRequester {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      network: isSet(object.network) ? String(object.network) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgRemoveLinkedRequester): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.network !== undefined && (obj.network = message.network);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveLinkedRequester>, I>>(object: I): MsgRemoveLinkedRequester {
    const message = createBaseMsgRemoveLinkedRequester();
    message.creator = object.creator ?? "";
    message.network = object.network ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgRemoveLinkedRequesterResponse(): MsgRemoveLinkedRequesterResponse {
  return {};
}

export const MsgRemoveLinkedRequesterResponse = {
  encode(_: MsgRemoveLinkedRequesterResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemoveLinkedRequesterResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveLinkedRequesterResponse();
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

  fromJSON(_: any): MsgRemoveLinkedRequesterResponse {
    return {};
  },

  toJSON(_: MsgRemoveLinkedRequesterResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveLinkedRequesterResponse>, I>>(
    _: I,
  ): MsgRemoveLinkedRequesterResponse {
    const message = createBaseMsgRemoveLinkedRequesterResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  DepositToken(request: MsgDepositToken): Promise<MsgDepositTokenResponse>;
  WithdrawToken(request: MsgWithdrawToken): Promise<MsgWithdrawTokenResponse>;
  AddLinkedRequester(request: MsgAddLinkedRequester): Promise<MsgAddLinkedRequesterResponse>;
  RemoveLinkedRequester(request: MsgRemoveLinkedRequester): Promise<MsgRemoveLinkedRequesterResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.DepositToken = this.DepositToken.bind(this);
    this.WithdrawToken = this.WithdrawToken.bind(this);
    this.AddLinkedRequester = this.AddLinkedRequester.bind(this);
    this.RemoveLinkedRequester = this.RemoveLinkedRequester.bind(this);
  }
  DepositToken(request: MsgDepositToken): Promise<MsgDepositTokenResponse> {
    const data = MsgDepositToken.encode(request).finish();
    const promise = this.rpc.request("dhpc.user.Msg", "DepositToken", data);
    return promise.then((data) => MsgDepositTokenResponse.decode(new _m0.Reader(data)));
  }

  WithdrawToken(request: MsgWithdrawToken): Promise<MsgWithdrawTokenResponse> {
    const data = MsgWithdrawToken.encode(request).finish();
    const promise = this.rpc.request("dhpc.user.Msg", "WithdrawToken", data);
    return promise.then((data) => MsgWithdrawTokenResponse.decode(new _m0.Reader(data)));
  }

  AddLinkedRequester(request: MsgAddLinkedRequester): Promise<MsgAddLinkedRequesterResponse> {
    const data = MsgAddLinkedRequester.encode(request).finish();
    const promise = this.rpc.request("dhpc.user.Msg", "AddLinkedRequester", data);
    return promise.then((data) => MsgAddLinkedRequesterResponse.decode(new _m0.Reader(data)));
  }

  RemoveLinkedRequester(request: MsgRemoveLinkedRequester): Promise<MsgRemoveLinkedRequesterResponse> {
    const data = MsgRemoveLinkedRequester.encode(request).finish();
    const promise = this.rpc.request("dhpc.user.Msg", "RemoveLinkedRequester", data);
    return promise.then((data) => MsgRemoveLinkedRequesterResponse.decode(new _m0.Reader(data)));
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
