package keeper

import (
	"bun/x/bun/types"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetPost(ctx sdk.Context, id uint64) (val types.Post, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))

	byteKey := make([]byte, 8)

	binary.BigEndian.PutUint64(byteKey, id)

	k.cdc.MustUnmarshal(store.Get(byteKey), &val)

	return val, true
}

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	count := k.GetPostCount(ctx)

	post.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, post.Id)

	appendedValue := k.cdc.MustMarshal(&post)

	store.Set(byteKey, appendedValue)

	k.SetPostCount(ctx, count+1)
	return count
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostCountKey))

	// Convert the PostCountKey to bytes
	byteKey := types.KeyPrefix(types.PostCountKey)

	// Get the value of the count
	bz := store.Get(byteKey)

	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}

	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostCountKey))

	byteKey := types.KeyPrefix(types.PostCountKey)

	bz := make([]byte, 8)

	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}
