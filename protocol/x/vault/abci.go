package vault

import (
	"runtime/debug"

	"github.com/bitoro-network/chain/protocol/lib/abci"
	"github.com/bitoro-network/chain/protocol/lib/log"
	"github.com/bitoro-network/chain/protocol/x/vault/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(
	ctx sdk.Context,
	keeper *keeper.Keeper,
) {
	// Panic is not expected in BeginBlocker, but we should recover instead of
	// halting the chain.
	if err := abci.RunCached(ctx, func(ctx sdk.Context) error {
		keeper.DecommissionNonPositiveEquityVaults(ctx)
		return nil
	}); err != nil {
		log.ErrorLog(
			ctx,
			"panic in vault BeginBlocker",
			"stack",
			string(debug.Stack()),
		)
	}
}

func EndBlocker(
	ctx sdk.Context,
	keeper *keeper.Keeper,
) {
	// Panic is not expected in EndBlocker, but we should recover instead of
	// halting the chain.
	if err := abci.RunCached(ctx, func(ctx sdk.Context) error {
		keeper.RefreshAllVaultOrders(ctx)
		keeper.SweepMainVaultBankBalance(ctx)
		return nil
	}); err != nil {
		log.ErrorLog(
			ctx,
			"panic in vault EndBlocker",
			"stack",
			string(debug.Stack()),
		)
	}
}
