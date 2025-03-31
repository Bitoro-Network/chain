package govplus

import (
	"github.com/Bitoro-Network/chain/protocol/x/govplus/keeper"
	"github.com/Bitoro-Network/chain/protocol/x/govplus/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the govplus module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {}

// ExportGenesis returns the govplus module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{}
}
