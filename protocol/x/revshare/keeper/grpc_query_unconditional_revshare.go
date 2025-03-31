package keeper

import (
	"context"

	"github.com/bitoro-network/chain/protocol/x/revshare/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) UnconditionalRevShareConfig(
	ctx context.Context,
	req *types.QueryUnconditionalRevShareConfig,
) (*types.QueryUnconditionalRevShareConfigResponse, error) {
	config, err := k.GetUnconditionalRevShareConfigParams(sdk.UnwrapSDKContext(ctx))
	if err != nil {
		return nil, err
	}
	return &types.QueryUnconditionalRevShareConfigResponse{
		Config: config,
	}, nil
}
