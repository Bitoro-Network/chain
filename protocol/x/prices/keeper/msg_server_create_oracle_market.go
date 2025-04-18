package keeper

import (
	"context"

	pricefeedmetrics "github.com/bitoro-network/chain/protocol/daemons/pricefeed/metrics"
	"github.com/bitoro-network/chain/protocol/lib"
	"github.com/bitoro-network/chain/protocol/lib/metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"
	gometrics "github.com/hashicorp/go-metrics"

	errorsmod "cosmossdk.io/errors"
	"github.com/bitoro-network/chain/protocol/x/prices/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (k msgServer) CreateOracleMarket(
	goCtx context.Context,
	msg *types.MsgCreateOracleMarket,
) (
	response *types.MsgCreateOracleMarketResponse,
	err error,
) {
	// Increment the appropriate success/error counter when the function finishes.
	defer func() {
		success := metrics.Success
		if err != nil {
			success = metrics.Error
		}
		telemetry.IncrCounterWithLabels(
			[]string{types.ModuleName, metrics.CreateOracleMarket, success},
			1,
			[]gometrics.Label{pricefeedmetrics.GetLabelForMarketId(msg.Params.Id)},
		)
	}()

	if !k.Keeper.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)

	exponent, err := k.Keeper.GetExponent(ctx, msg.Params.Pair)
	if err != nil {
		return nil, err
	}

	// Use zero oracle price to create the new market.
	// Note that valid oracle price updates cannot be zero (checked in MsgUpdateMarketPrices.ValidateBasic),
	// so a zero oracle price indicates that the oracle price has never been updated.
	zeroMarketPrice := types.MarketPrice{
		Id:       msg.Params.Id,
		Exponent: exponent,
		Price:    0,
	}
	if _, err = k.Keeper.CreateMarket(ctx, msg.Params, zeroMarketPrice); err != nil {
		return nil, err
	}

	return &types.MsgCreateOracleMarketResponse{}, nil
}
