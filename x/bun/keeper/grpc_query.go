package keeper

import (
	"bun/x/bun/types"
)

var _ types.QueryServer = Keeper{}
