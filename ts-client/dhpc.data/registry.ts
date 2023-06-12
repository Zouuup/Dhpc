import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddData } from "./types/dhpc/data/tx";
import { MsgDeleteData } from "./types/dhpc/data/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/dhpc.data.MsgAddData", MsgAddData],
    ["/dhpc.data.MsgDeleteData", MsgDeleteData],
    
];

export { msgTypes }