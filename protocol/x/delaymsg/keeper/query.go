package keeper

import (
	"github.com/bitoro-network/chain/protocol/x/delaymsg/types"
)

var _ types.QueryServer = Keeper{}
