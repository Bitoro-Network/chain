package vest

import (
	"math/rand"

	"github.com/bitoro-network/chain/protocol/testutil/sample"
	vestsimulation "github.com/bitoro-network/chain/protocol/x/vest/simulation"
	"github.com/bitoro-network/chain/protocol/x/vest/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_                            = sample.AccAddress
	_                            = vestsimulation.FindAccount
	_                            = simulation.MsgEntryKind
	_                            = baseapp.Paramspace
	_                            = rand.Rand{}
	_ module.AppModuleSimulation = AppModule{}
	_ module.HasProposalMsgs     = AppModule{}
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	vestGenesis := types.GenesisState{
		VestEntries: types.DefaultGenesis().VestEntries,
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&vestGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	return operations
}

// TODO(DEC-906): implement simulated gov proposal.
// ProposalMsgs doesn't return any content functions for governance proposals
func (AppModule) ProposalMsgs(_ module.SimulationState) []simtypes.WeightedProposalMsg {
	return nil
}
