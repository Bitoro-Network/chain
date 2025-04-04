package affiliates

import (
	"github.com/bitoro-network/chain/protocol/x/affiliates/keeper"
	"github.com/bitoro-network/chain/protocol/x/affiliates/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	err := k.UpdateAffiliateTiers(ctx, genState.AffiliateTiers)
	if err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	affiliateTiers, err := k.GetAllAffiliateTiers(ctx)
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		AffiliateTiers: affiliateTiers,
	}
}
