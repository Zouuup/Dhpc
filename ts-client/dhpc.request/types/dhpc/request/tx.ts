/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { MinerResponse } from "./miner_response";

export const protobufPackage = "dhpc.request";

export interface MsgCreateAllowedOracles {
  creator: string;
  name: string;
  address: string;
}

export interface MsgCreateAllowedOraclesResponse {
  id: number;
}

export interface MsgUpdateAllowedOracles {
  creator: string;
  id: number;
  name: string;
  address: string;
}

export interface MsgUpdateAllowedOraclesResponse {
}

export interface MsgDeleteAllowedOracles {
  creator: string;
  id: number;
}

export interface MsgDeleteAllowedOraclesResponse {
}

export interface MsgCreateMinerResponse {
  creator: string;
  uUID: string;
  requestUUID: string;
  hash: string;
  answer: number;
  dataUsed: string;
}

export interface MsgCreateMinerResponseResponse {
}

export interface MsgUpdateMinerResponse {
  creator: string;
  uUID: string;
  requestUUID: string;
  answer: number;
  salt: number;
}

export interface MsgUpdateMinerResponseResponse {
}

export interface MsgDeleteMinerResponse {
  creator: string;
  uUID: string;
}

export interface MsgDeleteMinerResponseResponse {
}

export interface MsgCreateRequestRecord {
  creator: string;
  uUID: string;
  tXhash: string;
  network: string;
  from: string;
  address: string;
  score: number;
  oracle: string;
  block: number;
  stage: number;
  miners: MinerResponse | undefined;
}

export interface MsgCreateRequestRecordResponse {
}

export interface MsgUpdateRequestRecord {
  creator: string;
  uUID: string;
  tXhash: string;
  network: string;
  from: string;
  address: string;
  score: number;
  oracle: string;
  block: number;
  stage: number;
  miners: MinerResponse | undefined;
}

export interface MsgUpdateRequestRecordResponse {
}

export interface MsgDeleteRequestRecord {
  creator: string;
  uUID: string;
}

export interface MsgDeleteRequestRecordResponse {
}

export interface MsgCreateTreasury {
  creator: string;
  address: string;
}

export interface MsgCreateTreasuryResponse {
}

export interface MsgUpdateTreasury {
  creator: string;
  address: string;
}

export interface MsgUpdateTreasuryResponse {
}

export interface MsgDeleteTreasury {
  creator: string;
}

export interface MsgDeleteTreasuryResponse {
}

function createBaseMsgCreateAllowedOracles(): MsgCreateAllowedOracles {
  return { creator: "", name: "", address: "" };
}

export const MsgCreateAllowedOracles = {
  encode(message: MsgCreateAllowedOracles, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAllowedOracles {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAllowedOracles();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
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

  fromJSON(object: any): MsgCreateAllowedOracles {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgCreateAllowedOracles): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAllowedOracles>, I>>(object: I): MsgCreateAllowedOracles {
    const message = createBaseMsgCreateAllowedOracles();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgCreateAllowedOraclesResponse(): MsgCreateAllowedOraclesResponse {
  return { id: 0 };
}

export const MsgCreateAllowedOraclesResponse = {
  encode(message: MsgCreateAllowedOraclesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAllowedOraclesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAllowedOraclesResponse();
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

  fromJSON(object: any): MsgCreateAllowedOraclesResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgCreateAllowedOraclesResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAllowedOraclesResponse>, I>>(
    object: I,
  ): MsgCreateAllowedOraclesResponse {
    const message = createBaseMsgCreateAllowedOraclesResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgUpdateAllowedOracles(): MsgUpdateAllowedOracles {
  return { creator: "", id: 0, name: "", address: "" };
}

export const MsgUpdateAllowedOracles = {
  encode(message: MsgUpdateAllowedOracles, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.address !== "") {
      writer.uint32(34).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateAllowedOracles {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateAllowedOracles();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.name = reader.string();
          break;
        case 4:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateAllowedOracles {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgUpdateAllowedOracles): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.name !== undefined && (obj.name = message.name);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateAllowedOracles>, I>>(object: I): MsgUpdateAllowedOracles {
    const message = createBaseMsgUpdateAllowedOracles();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgUpdateAllowedOraclesResponse(): MsgUpdateAllowedOraclesResponse {
  return {};
}

export const MsgUpdateAllowedOraclesResponse = {
  encode(_: MsgUpdateAllowedOraclesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateAllowedOraclesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateAllowedOraclesResponse();
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

  fromJSON(_: any): MsgUpdateAllowedOraclesResponse {
    return {};
  },

  toJSON(_: MsgUpdateAllowedOraclesResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateAllowedOraclesResponse>, I>>(_: I): MsgUpdateAllowedOraclesResponse {
    const message = createBaseMsgUpdateAllowedOraclesResponse();
    return message;
  },
};

function createBaseMsgDeleteAllowedOracles(): MsgDeleteAllowedOracles {
  return { creator: "", id: 0 };
}

export const MsgDeleteAllowedOracles = {
  encode(message: MsgDeleteAllowedOracles, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteAllowedOracles {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteAllowedOracles();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteAllowedOracles {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgDeleteAllowedOracles): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteAllowedOracles>, I>>(object: I): MsgDeleteAllowedOracles {
    const message = createBaseMsgDeleteAllowedOracles();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgDeleteAllowedOraclesResponse(): MsgDeleteAllowedOraclesResponse {
  return {};
}

export const MsgDeleteAllowedOraclesResponse = {
  encode(_: MsgDeleteAllowedOraclesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteAllowedOraclesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteAllowedOraclesResponse();
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

  fromJSON(_: any): MsgDeleteAllowedOraclesResponse {
    return {};
  },

  toJSON(_: MsgDeleteAllowedOraclesResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteAllowedOraclesResponse>, I>>(_: I): MsgDeleteAllowedOraclesResponse {
    const message = createBaseMsgDeleteAllowedOraclesResponse();
    return message;
  },
};

function createBaseMsgCreateMinerResponse(): MsgCreateMinerResponse {
  return { creator: "", uUID: "", requestUUID: "", hash: "", answer: 0, dataUsed: "" };
}

export const MsgCreateMinerResponse = {
  encode(message: MsgCreateMinerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uUID !== "") {
      writer.uint32(18).string(message.uUID);
    }
    if (message.requestUUID !== "") {
      writer.uint32(26).string(message.requestUUID);
    }
    if (message.hash !== "") {
      writer.uint32(34).string(message.hash);
    }
    if (message.answer !== 0) {
      writer.uint32(40).int32(message.answer);
    }
    if (message.dataUsed !== "") {
      writer.uint32(50).string(message.dataUsed);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateMinerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateMinerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uUID = reader.string();
          break;
        case 3:
          message.requestUUID = reader.string();
          break;
        case 4:
          message.hash = reader.string();
          break;
        case 5:
          message.answer = reader.int32();
          break;
        case 6:
          message.dataUsed = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateMinerResponse {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      uUID: isSet(object.uUID) ? String(object.uUID) : "",
      requestUUID: isSet(object.requestUUID) ? String(object.requestUUID) : "",
      hash: isSet(object.hash) ? String(object.hash) : "",
      answer: isSet(object.answer) ? Number(object.answer) : 0,
      dataUsed: isSet(object.dataUsed) ? String(object.dataUsed) : "",
    };
  },

  toJSON(message: MsgCreateMinerResponse): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uUID !== undefined && (obj.uUID = message.uUID);
    message.requestUUID !== undefined && (obj.requestUUID = message.requestUUID);
    message.hash !== undefined && (obj.hash = message.hash);
    message.answer !== undefined && (obj.answer = Math.round(message.answer));
    message.dataUsed !== undefined && (obj.dataUsed = message.dataUsed);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateMinerResponse>, I>>(object: I): MsgCreateMinerResponse {
    const message = createBaseMsgCreateMinerResponse();
    message.creator = object.creator ?? "";
    message.uUID = object.uUID ?? "";
    message.requestUUID = object.requestUUID ?? "";
    message.hash = object.hash ?? "";
    message.answer = object.answer ?? 0;
    message.dataUsed = object.dataUsed ?? "";
    return message;
  },
};

function createBaseMsgCreateMinerResponseResponse(): MsgCreateMinerResponseResponse {
  return {};
}

export const MsgCreateMinerResponseResponse = {
  encode(_: MsgCreateMinerResponseResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateMinerResponseResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateMinerResponseResponse();
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

  fromJSON(_: any): MsgCreateMinerResponseResponse {
    return {};
  },

  toJSON(_: MsgCreateMinerResponseResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateMinerResponseResponse>, I>>(_: I): MsgCreateMinerResponseResponse {
    const message = createBaseMsgCreateMinerResponseResponse();
    return message;
  },
};

function createBaseMsgUpdateMinerResponse(): MsgUpdateMinerResponse {
  return { creator: "", uUID: "", requestUUID: "", answer: 0, salt: 0 };
}

export const MsgUpdateMinerResponse = {
  encode(message: MsgUpdateMinerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uUID !== "") {
      writer.uint32(18).string(message.uUID);
    }
    if (message.requestUUID !== "") {
      writer.uint32(26).string(message.requestUUID);
    }
    if (message.answer !== 0) {
      writer.uint32(40).int32(message.answer);
    }
    if (message.salt !== 0) {
      writer.uint32(56).int32(message.salt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateMinerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateMinerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uUID = reader.string();
          break;
        case 3:
          message.requestUUID = reader.string();
          break;
        case 5:
          message.answer = reader.int32();
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

  fromJSON(object: any): MsgUpdateMinerResponse {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      uUID: isSet(object.uUID) ? String(object.uUID) : "",
      requestUUID: isSet(object.requestUUID) ? String(object.requestUUID) : "",
      answer: isSet(object.answer) ? Number(object.answer) : 0,
      salt: isSet(object.salt) ? Number(object.salt) : 0,
    };
  },

  toJSON(message: MsgUpdateMinerResponse): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uUID !== undefined && (obj.uUID = message.uUID);
    message.requestUUID !== undefined && (obj.requestUUID = message.requestUUID);
    message.answer !== undefined && (obj.answer = Math.round(message.answer));
    message.salt !== undefined && (obj.salt = Math.round(message.salt));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateMinerResponse>, I>>(object: I): MsgUpdateMinerResponse {
    const message = createBaseMsgUpdateMinerResponse();
    message.creator = object.creator ?? "";
    message.uUID = object.uUID ?? "";
    message.requestUUID = object.requestUUID ?? "";
    message.answer = object.answer ?? 0;
    message.salt = object.salt ?? 0;
    return message;
  },
};

function createBaseMsgUpdateMinerResponseResponse(): MsgUpdateMinerResponseResponse {
  return {};
}

export const MsgUpdateMinerResponseResponse = {
  encode(_: MsgUpdateMinerResponseResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateMinerResponseResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateMinerResponseResponse();
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

  fromJSON(_: any): MsgUpdateMinerResponseResponse {
    return {};
  },

  toJSON(_: MsgUpdateMinerResponseResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateMinerResponseResponse>, I>>(_: I): MsgUpdateMinerResponseResponse {
    const message = createBaseMsgUpdateMinerResponseResponse();
    return message;
  },
};

function createBaseMsgDeleteMinerResponse(): MsgDeleteMinerResponse {
  return { creator: "", uUID: "" };
}

export const MsgDeleteMinerResponse = {
  encode(message: MsgDeleteMinerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uUID !== "") {
      writer.uint32(18).string(message.uUID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteMinerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteMinerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uUID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteMinerResponse {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      uUID: isSet(object.uUID) ? String(object.uUID) : "",
    };
  },

  toJSON(message: MsgDeleteMinerResponse): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uUID !== undefined && (obj.uUID = message.uUID);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteMinerResponse>, I>>(object: I): MsgDeleteMinerResponse {
    const message = createBaseMsgDeleteMinerResponse();
    message.creator = object.creator ?? "";
    message.uUID = object.uUID ?? "";
    return message;
  },
};

function createBaseMsgDeleteMinerResponseResponse(): MsgDeleteMinerResponseResponse {
  return {};
}

export const MsgDeleteMinerResponseResponse = {
  encode(_: MsgDeleteMinerResponseResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteMinerResponseResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteMinerResponseResponse();
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

  fromJSON(_: any): MsgDeleteMinerResponseResponse {
    return {};
  },

  toJSON(_: MsgDeleteMinerResponseResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteMinerResponseResponse>, I>>(_: I): MsgDeleteMinerResponseResponse {
    const message = createBaseMsgDeleteMinerResponseResponse();
    return message;
  },
};

function createBaseMsgCreateRequestRecord(): MsgCreateRequestRecord {
  return {
    creator: "",
    uUID: "",
    tXhash: "",
    network: "",
    from: "",
    address: "",
    score: 0,
    oracle: "",
    block: 0,
    stage: 0,
    miners: undefined,
  };
}

export const MsgCreateRequestRecord = {
  encode(message: MsgCreateRequestRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uUID !== "") {
      writer.uint32(18).string(message.uUID);
    }
    if (message.tXhash !== "") {
      writer.uint32(26).string(message.tXhash);
    }
    if (message.network !== "") {
      writer.uint32(34).string(message.network);
    }
    if (message.from !== "") {
      writer.uint32(42).string(message.from);
    }
    if (message.address !== "") {
      writer.uint32(50).string(message.address);
    }
    if (message.score !== 0) {
      writer.uint32(56).int32(message.score);
    }
    if (message.oracle !== "") {
      writer.uint32(66).string(message.oracle);
    }
    if (message.block !== 0) {
      writer.uint32(72).uint64(message.block);
    }
    if (message.stage !== 0) {
      writer.uint32(80).int32(message.stage);
    }
    if (message.miners !== undefined) {
      MinerResponse.encode(message.miners, writer.uint32(90).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateRequestRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateRequestRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uUID = reader.string();
          break;
        case 3:
          message.tXhash = reader.string();
          break;
        case 4:
          message.network = reader.string();
          break;
        case 5:
          message.from = reader.string();
          break;
        case 6:
          message.address = reader.string();
          break;
        case 7:
          message.score = reader.int32();
          break;
        case 8:
          message.oracle = reader.string();
          break;
        case 9:
          message.block = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.stage = reader.int32();
          break;
        case 11:
          message.miners = MinerResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateRequestRecord {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      uUID: isSet(object.uUID) ? String(object.uUID) : "",
      tXhash: isSet(object.tXhash) ? String(object.tXhash) : "",
      network: isSet(object.network) ? String(object.network) : "",
      from: isSet(object.from) ? String(object.from) : "",
      address: isSet(object.address) ? String(object.address) : "",
      score: isSet(object.score) ? Number(object.score) : 0,
      oracle: isSet(object.oracle) ? String(object.oracle) : "",
      block: isSet(object.block) ? Number(object.block) : 0,
      stage: isSet(object.stage) ? Number(object.stage) : 0,
      miners: isSet(object.miners) ? MinerResponse.fromJSON(object.miners) : undefined,
    };
  },

  toJSON(message: MsgCreateRequestRecord): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uUID !== undefined && (obj.uUID = message.uUID);
    message.tXhash !== undefined && (obj.tXhash = message.tXhash);
    message.network !== undefined && (obj.network = message.network);
    message.from !== undefined && (obj.from = message.from);
    message.address !== undefined && (obj.address = message.address);
    message.score !== undefined && (obj.score = Math.round(message.score));
    message.oracle !== undefined && (obj.oracle = message.oracle);
    message.block !== undefined && (obj.block = Math.round(message.block));
    message.stage !== undefined && (obj.stage = Math.round(message.stage));
    message.miners !== undefined && (obj.miners = message.miners ? MinerResponse.toJSON(message.miners) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateRequestRecord>, I>>(object: I): MsgCreateRequestRecord {
    const message = createBaseMsgCreateRequestRecord();
    message.creator = object.creator ?? "";
    message.uUID = object.uUID ?? "";
    message.tXhash = object.tXhash ?? "";
    message.network = object.network ?? "";
    message.from = object.from ?? "";
    message.address = object.address ?? "";
    message.score = object.score ?? 0;
    message.oracle = object.oracle ?? "";
    message.block = object.block ?? 0;
    message.stage = object.stage ?? 0;
    message.miners = (object.miners !== undefined && object.miners !== null)
      ? MinerResponse.fromPartial(object.miners)
      : undefined;
    return message;
  },
};

function createBaseMsgCreateRequestRecordResponse(): MsgCreateRequestRecordResponse {
  return {};
}

export const MsgCreateRequestRecordResponse = {
  encode(_: MsgCreateRequestRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateRequestRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateRequestRecordResponse();
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

  fromJSON(_: any): MsgCreateRequestRecordResponse {
    return {};
  },

  toJSON(_: MsgCreateRequestRecordResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateRequestRecordResponse>, I>>(_: I): MsgCreateRequestRecordResponse {
    const message = createBaseMsgCreateRequestRecordResponse();
    return message;
  },
};

function createBaseMsgUpdateRequestRecord(): MsgUpdateRequestRecord {
  return {
    creator: "",
    uUID: "",
    tXhash: "",
    network: "",
    from: "",
    address: "",
    score: 0,
    oracle: "",
    block: 0,
    stage: 0,
    miners: undefined,
  };
}

export const MsgUpdateRequestRecord = {
  encode(message: MsgUpdateRequestRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uUID !== "") {
      writer.uint32(18).string(message.uUID);
    }
    if (message.tXhash !== "") {
      writer.uint32(26).string(message.tXhash);
    }
    if (message.network !== "") {
      writer.uint32(34).string(message.network);
    }
    if (message.from !== "") {
      writer.uint32(42).string(message.from);
    }
    if (message.address !== "") {
      writer.uint32(50).string(message.address);
    }
    if (message.score !== 0) {
      writer.uint32(56).int32(message.score);
    }
    if (message.oracle !== "") {
      writer.uint32(66).string(message.oracle);
    }
    if (message.block !== 0) {
      writer.uint32(72).uint64(message.block);
    }
    if (message.stage !== 0) {
      writer.uint32(80).int32(message.stage);
    }
    if (message.miners !== undefined) {
      MinerResponse.encode(message.miners, writer.uint32(90).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateRequestRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateRequestRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uUID = reader.string();
          break;
        case 3:
          message.tXhash = reader.string();
          break;
        case 4:
          message.network = reader.string();
          break;
        case 5:
          message.from = reader.string();
          break;
        case 6:
          message.address = reader.string();
          break;
        case 7:
          message.score = reader.int32();
          break;
        case 8:
          message.oracle = reader.string();
          break;
        case 9:
          message.block = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.stage = reader.int32();
          break;
        case 11:
          message.miners = MinerResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateRequestRecord {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      uUID: isSet(object.uUID) ? String(object.uUID) : "",
      tXhash: isSet(object.tXhash) ? String(object.tXhash) : "",
      network: isSet(object.network) ? String(object.network) : "",
      from: isSet(object.from) ? String(object.from) : "",
      address: isSet(object.address) ? String(object.address) : "",
      score: isSet(object.score) ? Number(object.score) : 0,
      oracle: isSet(object.oracle) ? String(object.oracle) : "",
      block: isSet(object.block) ? Number(object.block) : 0,
      stage: isSet(object.stage) ? Number(object.stage) : 0,
      miners: isSet(object.miners) ? MinerResponse.fromJSON(object.miners) : undefined,
    };
  },

  toJSON(message: MsgUpdateRequestRecord): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uUID !== undefined && (obj.uUID = message.uUID);
    message.tXhash !== undefined && (obj.tXhash = message.tXhash);
    message.network !== undefined && (obj.network = message.network);
    message.from !== undefined && (obj.from = message.from);
    message.address !== undefined && (obj.address = message.address);
    message.score !== undefined && (obj.score = Math.round(message.score));
    message.oracle !== undefined && (obj.oracle = message.oracle);
    message.block !== undefined && (obj.block = Math.round(message.block));
    message.stage !== undefined && (obj.stage = Math.round(message.stage));
    message.miners !== undefined && (obj.miners = message.miners ? MinerResponse.toJSON(message.miners) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateRequestRecord>, I>>(object: I): MsgUpdateRequestRecord {
    const message = createBaseMsgUpdateRequestRecord();
    message.creator = object.creator ?? "";
    message.uUID = object.uUID ?? "";
    message.tXhash = object.tXhash ?? "";
    message.network = object.network ?? "";
    message.from = object.from ?? "";
    message.address = object.address ?? "";
    message.score = object.score ?? 0;
    message.oracle = object.oracle ?? "";
    message.block = object.block ?? 0;
    message.stage = object.stage ?? 0;
    message.miners = (object.miners !== undefined && object.miners !== null)
      ? MinerResponse.fromPartial(object.miners)
      : undefined;
    return message;
  },
};

function createBaseMsgUpdateRequestRecordResponse(): MsgUpdateRequestRecordResponse {
  return {};
}

export const MsgUpdateRequestRecordResponse = {
  encode(_: MsgUpdateRequestRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateRequestRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateRequestRecordResponse();
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

  fromJSON(_: any): MsgUpdateRequestRecordResponse {
    return {};
  },

  toJSON(_: MsgUpdateRequestRecordResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateRequestRecordResponse>, I>>(_: I): MsgUpdateRequestRecordResponse {
    const message = createBaseMsgUpdateRequestRecordResponse();
    return message;
  },
};

function createBaseMsgDeleteRequestRecord(): MsgDeleteRequestRecord {
  return { creator: "", uUID: "" };
}

export const MsgDeleteRequestRecord = {
  encode(message: MsgDeleteRequestRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uUID !== "") {
      writer.uint32(18).string(message.uUID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteRequestRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteRequestRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uUID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteRequestRecord {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      uUID: isSet(object.uUID) ? String(object.uUID) : "",
    };
  },

  toJSON(message: MsgDeleteRequestRecord): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uUID !== undefined && (obj.uUID = message.uUID);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteRequestRecord>, I>>(object: I): MsgDeleteRequestRecord {
    const message = createBaseMsgDeleteRequestRecord();
    message.creator = object.creator ?? "";
    message.uUID = object.uUID ?? "";
    return message;
  },
};

function createBaseMsgDeleteRequestRecordResponse(): MsgDeleteRequestRecordResponse {
  return {};
}

export const MsgDeleteRequestRecordResponse = {
  encode(_: MsgDeleteRequestRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteRequestRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteRequestRecordResponse();
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

  fromJSON(_: any): MsgDeleteRequestRecordResponse {
    return {};
  },

  toJSON(_: MsgDeleteRequestRecordResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteRequestRecordResponse>, I>>(_: I): MsgDeleteRequestRecordResponse {
    const message = createBaseMsgDeleteRequestRecordResponse();
    return message;
  },
};

function createBaseMsgCreateTreasury(): MsgCreateTreasury {
  return { creator: "", address: "" };
}

export const MsgCreateTreasury = {
  encode(message: MsgCreateTreasury, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateTreasury {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateTreasury();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
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

  fromJSON(object: any): MsgCreateTreasury {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgCreateTreasury): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateTreasury>, I>>(object: I): MsgCreateTreasury {
    const message = createBaseMsgCreateTreasury();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgCreateTreasuryResponse(): MsgCreateTreasuryResponse {
  return {};
}

export const MsgCreateTreasuryResponse = {
  encode(_: MsgCreateTreasuryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateTreasuryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateTreasuryResponse();
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

  fromJSON(_: any): MsgCreateTreasuryResponse {
    return {};
  },

  toJSON(_: MsgCreateTreasuryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateTreasuryResponse>, I>>(_: I): MsgCreateTreasuryResponse {
    const message = createBaseMsgCreateTreasuryResponse();
    return message;
  },
};

function createBaseMsgUpdateTreasury(): MsgUpdateTreasury {
  return { creator: "", address: "" };
}

export const MsgUpdateTreasury = {
  encode(message: MsgUpdateTreasury, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateTreasury {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateTreasury();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
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

  fromJSON(object: any): MsgUpdateTreasury {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgUpdateTreasury): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateTreasury>, I>>(object: I): MsgUpdateTreasury {
    const message = createBaseMsgUpdateTreasury();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgUpdateTreasuryResponse(): MsgUpdateTreasuryResponse {
  return {};
}

export const MsgUpdateTreasuryResponse = {
  encode(_: MsgUpdateTreasuryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateTreasuryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateTreasuryResponse();
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

  fromJSON(_: any): MsgUpdateTreasuryResponse {
    return {};
  },

  toJSON(_: MsgUpdateTreasuryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateTreasuryResponse>, I>>(_: I): MsgUpdateTreasuryResponse {
    const message = createBaseMsgUpdateTreasuryResponse();
    return message;
  },
};

function createBaseMsgDeleteTreasury(): MsgDeleteTreasury {
  return { creator: "" };
}

export const MsgDeleteTreasury = {
  encode(message: MsgDeleteTreasury, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteTreasury {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteTreasury();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteTreasury {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgDeleteTreasury): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteTreasury>, I>>(object: I): MsgDeleteTreasury {
    const message = createBaseMsgDeleteTreasury();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgDeleteTreasuryResponse(): MsgDeleteTreasuryResponse {
  return {};
}

export const MsgDeleteTreasuryResponse = {
  encode(_: MsgDeleteTreasuryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteTreasuryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteTreasuryResponse();
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

  fromJSON(_: any): MsgDeleteTreasuryResponse {
    return {};
  },

  toJSON(_: MsgDeleteTreasuryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteTreasuryResponse>, I>>(_: I): MsgDeleteTreasuryResponse {
    const message = createBaseMsgDeleteTreasuryResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateAllowedOracles(request: MsgCreateAllowedOracles): Promise<MsgCreateAllowedOraclesResponse>;
  UpdateAllowedOracles(request: MsgUpdateAllowedOracles): Promise<MsgUpdateAllowedOraclesResponse>;
  DeleteAllowedOracles(request: MsgDeleteAllowedOracles): Promise<MsgDeleteAllowedOraclesResponse>;
  CreateMinerResponse(request: MsgCreateMinerResponse): Promise<MsgCreateMinerResponseResponse>;
  UpdateMinerResponse(request: MsgUpdateMinerResponse): Promise<MsgUpdateMinerResponseResponse>;
  DeleteMinerResponse(request: MsgDeleteMinerResponse): Promise<MsgDeleteMinerResponseResponse>;
  CreateRequestRecord(request: MsgCreateRequestRecord): Promise<MsgCreateRequestRecordResponse>;
  UpdateRequestRecord(request: MsgUpdateRequestRecord): Promise<MsgUpdateRequestRecordResponse>;
  DeleteRequestRecord(request: MsgDeleteRequestRecord): Promise<MsgDeleteRequestRecordResponse>;
  CreateTreasury(request: MsgCreateTreasury): Promise<MsgCreateTreasuryResponse>;
  UpdateTreasury(request: MsgUpdateTreasury): Promise<MsgUpdateTreasuryResponse>;
  DeleteTreasury(request: MsgDeleteTreasury): Promise<MsgDeleteTreasuryResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateAllowedOracles = this.CreateAllowedOracles.bind(this);
    this.UpdateAllowedOracles = this.UpdateAllowedOracles.bind(this);
    this.DeleteAllowedOracles = this.DeleteAllowedOracles.bind(this);
    this.CreateMinerResponse = this.CreateMinerResponse.bind(this);
    this.UpdateMinerResponse = this.UpdateMinerResponse.bind(this);
    this.DeleteMinerResponse = this.DeleteMinerResponse.bind(this);
    this.CreateRequestRecord = this.CreateRequestRecord.bind(this);
    this.UpdateRequestRecord = this.UpdateRequestRecord.bind(this);
    this.DeleteRequestRecord = this.DeleteRequestRecord.bind(this);
    this.CreateTreasury = this.CreateTreasury.bind(this);
    this.UpdateTreasury = this.UpdateTreasury.bind(this);
    this.DeleteTreasury = this.DeleteTreasury.bind(this);
  }
  CreateAllowedOracles(request: MsgCreateAllowedOracles): Promise<MsgCreateAllowedOraclesResponse> {
    const data = MsgCreateAllowedOracles.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Msg", "CreateAllowedOracles", data);
    return promise.then((data) => MsgCreateAllowedOraclesResponse.decode(new _m0.Reader(data)));
  }

  UpdateAllowedOracles(request: MsgUpdateAllowedOracles): Promise<MsgUpdateAllowedOraclesResponse> {
    const data = MsgUpdateAllowedOracles.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Msg", "UpdateAllowedOracles", data);
    return promise.then((data) => MsgUpdateAllowedOraclesResponse.decode(new _m0.Reader(data)));
  }

  DeleteAllowedOracles(request: MsgDeleteAllowedOracles): Promise<MsgDeleteAllowedOraclesResponse> {
    const data = MsgDeleteAllowedOracles.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Msg", "DeleteAllowedOracles", data);
    return promise.then((data) => MsgDeleteAllowedOraclesResponse.decode(new _m0.Reader(data)));
  }

  CreateMinerResponse(request: MsgCreateMinerResponse): Promise<MsgCreateMinerResponseResponse> {
    const data = MsgCreateMinerResponse.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Msg", "CreateMinerResponse", data);
    return promise.then((data) => MsgCreateMinerResponseResponse.decode(new _m0.Reader(data)));
  }

  UpdateMinerResponse(request: MsgUpdateMinerResponse): Promise<MsgUpdateMinerResponseResponse> {
    const data = MsgUpdateMinerResponse.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Msg", "UpdateMinerResponse", data);
    return promise.then((data) => MsgUpdateMinerResponseResponse.decode(new _m0.Reader(data)));
  }

  DeleteMinerResponse(request: MsgDeleteMinerResponse): Promise<MsgDeleteMinerResponseResponse> {
    const data = MsgDeleteMinerResponse.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Msg", "DeleteMinerResponse", data);
    return promise.then((data) => MsgDeleteMinerResponseResponse.decode(new _m0.Reader(data)));
  }

  CreateRequestRecord(request: MsgCreateRequestRecord): Promise<MsgCreateRequestRecordResponse> {
    const data = MsgCreateRequestRecord.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Msg", "CreateRequestRecord", data);
    return promise.then((data) => MsgCreateRequestRecordResponse.decode(new _m0.Reader(data)));
  }

  UpdateRequestRecord(request: MsgUpdateRequestRecord): Promise<MsgUpdateRequestRecordResponse> {
    const data = MsgUpdateRequestRecord.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Msg", "UpdateRequestRecord", data);
    return promise.then((data) => MsgUpdateRequestRecordResponse.decode(new _m0.Reader(data)));
  }

  DeleteRequestRecord(request: MsgDeleteRequestRecord): Promise<MsgDeleteRequestRecordResponse> {
    const data = MsgDeleteRequestRecord.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Msg", "DeleteRequestRecord", data);
    return promise.then((data) => MsgDeleteRequestRecordResponse.decode(new _m0.Reader(data)));
  }

  CreateTreasury(request: MsgCreateTreasury): Promise<MsgCreateTreasuryResponse> {
    const data = MsgCreateTreasury.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Msg", "CreateTreasury", data);
    return promise.then((data) => MsgCreateTreasuryResponse.decode(new _m0.Reader(data)));
  }

  UpdateTreasury(request: MsgUpdateTreasury): Promise<MsgUpdateTreasuryResponse> {
    const data = MsgUpdateTreasury.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Msg", "UpdateTreasury", data);
    return promise.then((data) => MsgUpdateTreasuryResponse.decode(new _m0.Reader(data)));
  }

  DeleteTreasury(request: MsgDeleteTreasury): Promise<MsgDeleteTreasuryResponse> {
    const data = MsgDeleteTreasury.encode(request).finish();
    const promise = this.rpc.request("dhpc.request.Msg", "DeleteTreasury", data);
    return promise.then((data) => MsgDeleteTreasuryResponse.decode(new _m0.Reader(data)));
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
