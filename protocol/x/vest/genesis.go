package vest

import (
	"github.com/bitoro-network/chain/protocol/x/vest/keeper"
	"github.com/bitoro-network/chain/protocol/x/vest/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, entry := range genState.VestEntries {
		if err := k.SetVestEntry(ctx, entry); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	allEntries := k.GetAllVestEntries(ctx)
	return &types.GenesisState{
		VestEntries: allEntries,
	}
}
