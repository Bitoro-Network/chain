package keeper

import (
	"github.com/bitoro-network/chain/protocol/x/assets/types"
)

var _ types.QueryServer = Keeper{}
