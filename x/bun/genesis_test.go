package bun_test

import (
	"testing"

	keepertest "bun/testutil/keeper"
	"bun/testutil/nullify"
	"bun/x/bun"
	"bun/x/bun/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BunKeeper(t)
	bun.InitGenesis(ctx, *k, genesisState)
	got := bun.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
