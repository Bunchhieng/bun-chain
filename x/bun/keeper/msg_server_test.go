package keeper_test

import (
	"context"
	"testing"

	keepertest "bun/testutil/keeper"
	"bun/x/bun/keeper"
	"bun/x/bun/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BunKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
