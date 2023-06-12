import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgWithdrawToken } from "./types/dhpc/user/tx";
import { MsgAddLinkedRequester } from "./types/dhpc/user/tx";
import { MsgRemoveLinkedRequester } from "./types/dhpc/user/tx";
import { MsgDepositToken } from "./types/dhpc/user/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/dhpc.user.MsgWithdrawToken", MsgWithdrawToken],
    ["/dhpc.user.MsgAddLinkedRequester", MsgAddLinkedRequester],
    ["/dhpc.user.MsgRemoveLinkedRequester", MsgRemoveLinkedRequester],
    ["/dhpc.user.MsgDepositToken", MsgDepositToken],
    
];

export { msgTypes }