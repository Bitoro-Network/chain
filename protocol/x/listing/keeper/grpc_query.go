package keeper

import (
	"github.com/Bitoro-Network/chain/protocol/x/listing/types"
)

var _ types.QueryServer = Keeper{}
