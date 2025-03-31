package prices

import (
	pricestypes "github.com/Bitoro-Network/chain/protocol/x/prices/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PriceUpdateGenerator is an interface to abstract the logic of retrieving a
// `MsgUpdateMarketPrices` for any block.
type PriceUpdateGenerator interface {
	GetValidMarketPriceUpdates(ctx sdk.Context, extCommitBz []byte) (*pricestypes.MsgUpdateMarketPrices, error)
}
