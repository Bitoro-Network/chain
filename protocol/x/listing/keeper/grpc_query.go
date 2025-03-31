package keeper

import (
	"github.com/bitoro-network/chain/protocol/x/listing/types"
)

var _ types.QueryServer = Keeper{}
