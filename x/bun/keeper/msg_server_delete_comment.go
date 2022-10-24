package keeper

import (
	"context"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	"bun/x/bun/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteComment(goCtx context.Context, msg *types.MsgDeleteComment) (*types.MsgDeleteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	comment, exist := k.GetComment(ctx, msg.CommentID)
	if !exist {
		return nil, types.ErrCommentNotFound
	}

	if msg.PostID != comment.PostID {
		return nil, types.ErrID
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommentKey))
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, comment.Id)
	store.Delete(bz)

	return &types.MsgDeleteCommentResponse{}, nil
}
