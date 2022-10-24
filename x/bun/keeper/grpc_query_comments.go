package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	"bun/x/bun/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Comments(goCtx context.Context, req *types.QueryCommentsRequest) (*types.QueryCommentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var comments []types.Comment
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	commentStore := prefix.NewStore(store, types.KeyPrefix(types.CommentKey))

	post, _ := k.GetPost(ctx, req.Id)

	postId := post.Id

	pageRes, err := query.Paginate(commentStore, req.Pagination, func(key []byte, value []byte) error {
		var comment types.Comment
		if err := k.cdc.Unmarshal(value, &comment); err != nil {
			return err
		}
		if comment.PostID == postId {
			comments = append(comments, comment)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryCommentsResponse{Post: post, Comment: comments, Pagination: pageRes}, nil
}
