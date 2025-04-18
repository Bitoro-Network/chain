package keeper

import (
	"context"

	"github.com/bitoro-network/chain/protocol/x/revshare/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) MarketMapperRevShareDetails(
	ctx context.Context,
	req *types.QueryMarketMapperRevShareDetails,
) (*types.QueryMarketMapperRevShareDetailsResponse, error) {
	revShareDetails, err := k.GetMarketMapperRevShareDetails(sdk.UnwrapSDKContext(ctx), req.MarketId)
	if err != nil {
		return nil, err
	}
	return &types.QueryMarketMapperRevShareDetailsResponse{
		Details: revShareDetails,
	}, nil
}
