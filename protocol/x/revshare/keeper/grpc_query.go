package keeper

import (
	"github.com/Bitoro-Network/chain/protocol/x/revshare/types"
)

var _ types.QueryServer = Keeper{}
