package keeper

import (
	"bun/x/dex/types"
)

var _ types.QueryServer = Keeper{}
