package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/bitoro-network/chain/protocol/lib"
	"github.com/bitoro-network/chain/protocol/x/perpetuals/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (k msgServer) UpdatePerpetualParams(
	goCtx context.Context,
	msg *types.MsgUpdatePerpetualParams,
) (*types.MsgUpdatePerpetualParamsResponse, error) {
	if !k.Keeper.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)

	_, err := k.Keeper.ModifyPerpetual(
		ctx,
		msg.PerpetualParams.Id,
		msg.PerpetualParams.Ticker,
		msg.PerpetualParams.MarketId,
		msg.PerpetualParams.DefaultFundingPpm,
		msg.PerpetualParams.LiquidityTier,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdatePerpetualParamsResponse{}, nil
}
