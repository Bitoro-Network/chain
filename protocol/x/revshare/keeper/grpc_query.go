package keeper

import (
	"github.com/bitoro-network/chain/protocol/x/revshare/types"
)

var _ types.QueryServer = Keeper{}
