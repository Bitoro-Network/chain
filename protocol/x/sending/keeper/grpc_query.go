package keeper

import (
	"github.com/Bitoro-Network/chain/protocol/x/sending/types"
)

var _ types.QueryServer = Keeper{}
