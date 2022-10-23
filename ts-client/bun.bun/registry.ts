import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateComment } from "./types/bun/bun/tx";
import { MsgCreatePost } from "./types/bun/bun/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/bun.bun.MsgCreateComment", MsgCreateComment],
    ["/bun.bun.MsgCreatePost", MsgCreatePost],
    
];

export { msgTypes }