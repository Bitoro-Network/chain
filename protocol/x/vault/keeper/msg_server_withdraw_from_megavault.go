package keeper

import (
	"context"

	"github.com/bitoro-network/chain/protocol/dtypes"
	"github.com/bitoro-network/chain/protocol/lib"
	"github.com/bitoro-network/chain/protocol/x/vault/types"
)

// WithdrawFromMegavault withdraws from megavault to a subaccount.
func (k msgServer) WithdrawFromMegavault(
	goCtx context.Context,
	msg *types.MsgWithdrawFromMegavault,
) (*types.MsgWithdrawFromMegavaultResponse, error) {
	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)

	redeemedQuoteQuantums, err := k.Keeper.WithdrawFromMegavault(
		ctx,
		msg.SubaccountId,
		msg.Shares.NumShares.BigInt(),
		msg.MinQuoteQuantums.BigInt(),
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawFromMegavaultResponse{
		QuoteQuantums: dtypes.NewIntFromBigInt(redeemedQuoteQuantums),
	}, nil
}
