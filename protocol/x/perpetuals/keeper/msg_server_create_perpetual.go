package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	"github.com/bitoro-network/chain/protocol/lib"
	"github.com/bitoro-network/chain/protocol/x/perpetuals/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (k msgServer) CreatePerpetual(
	goCtx context.Context,
	msg *types.MsgCreatePerpetual,
) (*types.MsgCreatePerpetualResponse, error) {
	if !k.Keeper.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)

	_, err := k.Keeper.CreatePerpetual(
		ctx,
		msg.Params.Id,
		msg.Params.Ticker,
		msg.Params.MarketId,
		msg.Params.AtomicResolution,
		msg.Params.DefaultFundingPpm,
		msg.Params.LiquidityTier,
		msg.Params.MarketType,
	)
	if err != nil {
		return &types.MsgCreatePerpetualResponse{}, err
	}

	return &types.MsgCreatePerpetualResponse{}, nil
}
