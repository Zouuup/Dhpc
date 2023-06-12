import { GeneratedType } from "@cosmjs/proto-signing";
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

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/dhpc.request.MsgDeleteTreasury", MsgDeleteTreasury],
    ["/dhpc.request.MsgCreateAllowedOracles", MsgCreateAllowedOracles],
    ["/dhpc.request.MsgDeleteRequestRecord", MsgDeleteRequestRecord],
    ["/dhpc.request.MsgUpdateAllowedOracles", MsgUpdateAllowedOracles],
    ["/dhpc.request.MsgCreateMinerResponse", MsgCreateMinerResponse],
    ["/dhpc.request.MsgCreateRequestRecord", MsgCreateRequestRecord],
    ["/dhpc.request.MsgDeleteAllowedOracles", MsgDeleteAllowedOracles],
    ["/dhpc.request.MsgUpdateMinerResponse", MsgUpdateMinerResponse],
    ["/dhpc.request.MsgCreateTreasury", MsgCreateTreasury],
    ["/dhpc.request.MsgUpdateRequestRecord", MsgUpdateRequestRecord],
    ["/dhpc.request.MsgDeleteMinerResponse", MsgDeleteMinerResponse],
    ["/dhpc.request.MsgUpdateTreasury", MsgUpdateTreasury],
    
];

export { msgTypes }