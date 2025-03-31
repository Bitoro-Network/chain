package keeper

import (
	"github.com/Bitoro-Network/chain/protocol/x/vest/types"
)

var _ types.QueryServer = Keeper{}
