// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgDeleteTreasury } from "./types/dhpc/request/tx";
import { MsgCreateAllowedOracles } from "./types/dhpc/request/tx";
import { MsgDeleteRequestRecord } from "./types/dhpc/request/tx";
import { MsgUpdateAllowedOracles } from "./types/dhpc/request/tx";
import { MsgCreateMinerResponse } from "./types/dhpc/request/tx";
import { MsgCreateRequestRecord } from "./types/dhpc/request/tx";
import { MsgDeleteAllowedOracles } from "./types/dhpc/request/tx";
import { MsgUpdateMinerResponse } from "./types/dhpc/request/tx";
import { MsgCreateTreasury } from "./types/dhpc/request/tx";
import { MsgUpdateRequestRecord } from "./types/dhpc/request/tx";
import { MsgDeleteMinerResponse } from "./types/dhpc/request/tx";
import { MsgUpdateTreasury } from "./types/dhpc/request/tx";

import { AddTreasury as typeAddTreasury} from "./types"
import { AllowedOracles as typeAllowedOracles} from "./types"
import { MinerResponse as typeMinerResponse} from "./types"
import { Params as typeParams} from "./types"
import { RequestRecord as typeRequestRecord} from "./types"
import { Treasury as typeTreasury} from "./types"

export { MsgDeleteTreasury, MsgCreateAllowedOracles, MsgDeleteRequestRecord, MsgUpdateAllowedOracles, MsgCreateMinerResponse, MsgCreateRequestRecord, MsgDeleteAllowedOracles, MsgUpdateMinerResponse, MsgCreateTreasury, MsgUpdateRequestRecord, MsgDeleteMinerResponse, MsgUpdateTreasury };

type sendMsgDeleteTreasuryParams = {
  value: MsgDeleteTreasury,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateAllowedOraclesParams = {
  value: MsgCreateAllowedOracles,
  fee?: StdFee,
  memo?: string
};

type sendMsgDeleteRequestRecordParams = {
  value: MsgDeleteRequestRecord,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateAllowedOraclesParams = {
  value: MsgUpdateAllowedOracles,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateMinerResponseParams = {
  value: MsgCreateMinerResponse,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateRequestRecordParams = {
  value: MsgCreateRequestRecord,
  fee?: StdFee,
  memo?: string
};

type sendMsgDeleteAllowedOraclesParams = {
  value: MsgDeleteAllowedOracles,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateMinerResponseParams = {
  value: MsgUpdateMinerResponse,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateTreasuryParams = {
  value: MsgCreateTreasury,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateRequestRecordParams = {
  value: MsgUpdateRequestRecord,
  fee?: StdFee,
  memo?: string
};

type sendMsgDeleteMinerResponseParams = {
  value: MsgDeleteMinerResponse,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateTreasuryParams = {
  value: MsgUpdateTreasury,
  fee?: StdFee,
  memo?: string
};


type msgDeleteTreasuryParams = {
  value: MsgDeleteTreasury,
};

type msgCreateAllowedOraclesParams = {
  value: MsgCreateAllowedOracles,
};

type msgDeleteRequestRecordParams = {
  value: MsgDeleteRequestRecord,
};

type msgUpdateAllowedOraclesParams = {
  value: MsgUpdateAllowedOracles,
};

type msgCreateMinerResponseParams = {
  value: MsgCreateMinerResponse,
};

type msgCreateRequestRecordParams = {
  value: MsgCreateRequestRecord,
};

type msgDeleteAllowedOraclesParams = {
  value: MsgDeleteAllowedOracles,
};

type msgUpdateMinerResponseParams = {
  value: MsgUpdateMinerResponse,
};

type msgCreateTreasuryParams = {
  value: MsgCreateTreasury,
};

type msgUpdateRequestRecordParams = {
  value: MsgUpdateRequestRecord,
};

type msgDeleteMinerResponseParams = {
  value: MsgDeleteMinerResponse,
};

type msgUpdateTreasuryParams = {
  value: MsgUpdateTreasury,
};


export const registry = new Registry(msgTypes);

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	const structure: {fields: Field[]} = { fields: [] }
	for (let [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgDeleteTreasury({ value, fee, memo }: sendMsgDeleteTreasuryParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDeleteTreasury: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDeleteTreasury({ value: MsgDeleteTreasury.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDeleteTreasury: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateAllowedOracles({ value, fee, memo }: sendMsgCreateAllowedOraclesParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateAllowedOracles: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateAllowedOracles({ value: MsgCreateAllowedOracles.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateAllowedOracles: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDeleteRequestRecord({ value, fee, memo }: sendMsgDeleteRequestRecordParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDeleteRequestRecord: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDeleteRequestRecord({ value: MsgDeleteRequestRecord.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDeleteRequestRecord: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateAllowedOracles({ value, fee, memo }: sendMsgUpdateAllowedOraclesParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateAllowedOracles: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateAllowedOracles({ value: MsgUpdateAllowedOracles.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateAllowedOracles: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateMinerResponse({ value, fee, memo }: sendMsgCreateMinerResponseParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateMinerResponse: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateMinerResponse({ value: MsgCreateMinerResponse.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateMinerResponse: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateRequestRecord({ value, fee, memo }: sendMsgCreateRequestRecordParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateRequestRecord: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateRequestRecord({ value: MsgCreateRequestRecord.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateRequestRecord: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDeleteAllowedOracles({ value, fee, memo }: sendMsgDeleteAllowedOraclesParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDeleteAllowedOracles: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDeleteAllowedOracles({ value: MsgDeleteAllowedOracles.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDeleteAllowedOracles: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateMinerResponse({ value, fee, memo }: sendMsgUpdateMinerResponseParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateMinerResponse: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateMinerResponse({ value: MsgUpdateMinerResponse.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateMinerResponse: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateTreasury({ value, fee, memo }: sendMsgCreateTreasuryParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateTreasury: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateTreasury({ value: MsgCreateTreasury.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateTreasury: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateRequestRecord({ value, fee, memo }: sendMsgUpdateRequestRecordParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateRequestRecord: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateRequestRecord({ value: MsgUpdateRequestRecord.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateRequestRecord: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDeleteMinerResponse({ value, fee, memo }: sendMsgDeleteMinerResponseParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDeleteMinerResponse: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDeleteMinerResponse({ value: MsgDeleteMinerResponse.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDeleteMinerResponse: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateTreasury({ value, fee, memo }: sendMsgUpdateTreasuryParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateTreasury: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateTreasury({ value: MsgUpdateTreasury.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateTreasury: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgDeleteTreasury({ value }: msgDeleteTreasuryParams): EncodeObject {
			try {
				return { typeUrl: "/dhpc.request.MsgDeleteTreasury", value: MsgDeleteTreasury.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeleteTreasury: Could not create message: ' + e.message)
			}
		},
		
		msgCreateAllowedOracles({ value }: msgCreateAllowedOraclesParams): EncodeObject {
			try {
				return { typeUrl: "/dhpc.request.MsgCreateAllowedOracles", value: MsgCreateAllowedOracles.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateAllowedOracles: Could not create message: ' + e.message)
			}
		},
		
		msgDeleteRequestRecord({ value }: msgDeleteRequestRecordParams): EncodeObject {
			try {
				return { typeUrl: "/dhpc.request.MsgDeleteRequestRecord", value: MsgDeleteRequestRecord.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeleteRequestRecord: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateAllowedOracles({ value }: msgUpdateAllowedOraclesParams): EncodeObject {
			try {
				return { typeUrl: "/dhpc.request.MsgUpdateAllowedOracles", value: MsgUpdateAllowedOracles.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateAllowedOracles: Could not create message: ' + e.message)
			}
		},
		
		msgCreateMinerResponse({ value }: msgCreateMinerResponseParams): EncodeObject {
			try {
				return { typeUrl: "/dhpc.request.MsgCreateMinerResponse", value: MsgCreateMinerResponse.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateMinerResponse: Could not create message: ' + e.message)
			}
		},
		
		msgCreateRequestRecord({ value }: msgCreateRequestRecordParams): EncodeObject {
			try {
				return { typeUrl: "/dhpc.request.MsgCreateRequestRecord", value: MsgCreateRequestRecord.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateRequestRecord: Could not create message: ' + e.message)
			}
		},
		
		msgDeleteAllowedOracles({ value }: msgDeleteAllowedOraclesParams): EncodeObject {
			try {
				return { typeUrl: "/dhpc.request.MsgDeleteAllowedOracles", value: MsgDeleteAllowedOracles.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeleteAllowedOracles: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateMinerResponse({ value }: msgUpdateMinerResponseParams): EncodeObject {
			try {
				return { typeUrl: "/dhpc.request.MsgUpdateMinerResponse", value: MsgUpdateMinerResponse.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateMinerResponse: Could not create message: ' + e.message)
			}
		},
		
		msgCreateTreasury({ value }: msgCreateTreasuryParams): EncodeObject {
			try {
				return { typeUrl: "/dhpc.request.MsgCreateTreasury", value: MsgCreateTreasury.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateTreasury: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateRequestRecord({ value }: msgUpdateRequestRecordParams): EncodeObject {
			try {
				return { typeUrl: "/dhpc.request.MsgUpdateRequestRecord", value: MsgUpdateRequestRecord.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateRequestRecord: Could not create message: ' + e.message)
			}
		},
		
		msgDeleteMinerResponse({ value }: msgDeleteMinerResponseParams): EncodeObject {
			try {
				return { typeUrl: "/dhpc.request.MsgDeleteMinerResponse", value: MsgDeleteMinerResponse.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeleteMinerResponse: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateTreasury({ value }: msgUpdateTreasuryParams): EncodeObject {
			try {
				return { typeUrl: "/dhpc.request.MsgUpdateTreasury", value: MsgUpdateTreasury.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateTreasury: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	public structure: Record<string,unknown>;
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		this.structure =  {
						AddTreasury: getStructure(typeAddTreasury.fromPartial({})),
						AllowedOracles: getStructure(typeAllowedOracles.fromPartial({})),
						MinerResponse: getStructure(typeMinerResponse.fromPartial({})),
						Params: getStructure(typeParams.fromPartial({})),
						RequestRecord: getStructure(typeRequestRecord.fromPartial({})),
						Treasury: getStructure(typeTreasury.fromPartial({})),
						
		};
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			DhpcRequest: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;