package keeper_test

import (
	"context"
	"testing"

	keepertest "bun/testutil/keeper"
	"bun/x/dex/keeper"
	"bun/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DexKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
