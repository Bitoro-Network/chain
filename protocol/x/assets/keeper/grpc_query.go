package keeper

import (
	"github.com/Bitoro-Network/chain/protocol/x/assets/types"
)

var _ types.QueryServer = Keeper{}
