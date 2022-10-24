package keeper

import (
	"bun/x/bun/types"
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateComment(goCtx context.Context, msg *types.MsgCreateComment) (*types.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	post, found := k.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.ErrNotFound
	}

	var comment = types.Comment{
		Creator:   msg.Creator,
		PostID:    msg.PostID,
		Title:     msg.Title,
		Body:      msg.Body,
		Id:        msg.Id,
		CreatedAt: ctx.BlockHeight(),
	}

	if comment.CreatedAt > post.CreatedAt+100 {
		return nil, sdkerrors.Wrapf(types.ErrCommentOld, "Comment created at %d is older than post created at %d", comment.CreatedAt, post.CreatedAt)
	}

	id := k.AppendComment(ctx, comment)
	return &types.MsgCreateCommentResponse{Id: id}, nil
}
