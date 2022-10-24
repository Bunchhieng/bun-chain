import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreatePost } from "./types/bun/bun/tx";
import { MsgCreateComment } from "./types/bun/bun/tx";
import { MsgDeleteComment } from "./types/bun/bun/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/bun.bun.MsgCreatePost", MsgCreatePost],
    ["/bun.bun.MsgCreateComment", MsgCreateComment],
    ["/bun.bun.MsgDeleteComment", MsgDeleteComment],
    
];

export { msgTypes }