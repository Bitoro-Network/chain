package keeper

import (
	"context"

	"github.com/bitoro-network/chain/protocol/x/listing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) MarketsHardCap(
	ctx context.Context,
	req *types.QueryMarketsHardCap,
) (*types.QueryMarketsHardCapResponse, error) {
	hardCap := k.GetMarketsHardCap(sdk.UnwrapSDKContext(ctx))
	return &types.QueryMarketsHardCapResponse{
		HardCap: hardCap,
	}, nil
}
