package keeper

import (
	"github.com/bitoro-network/chain/protocol/lib"
	"github.com/bitoro-network/chain/protocol/mocks"
	delaymsgtypes "github.com/bitoro-network/chain/protocol/x/delaymsg/types"
	dbm "github.com/cosmos/cosmos-db"

	storetypes "cosmossdk.io/store/types"
	epochskeeper "github.com/bitoro-network/chain/protocol/x/epochs/keeper"
	"github.com/bitoro-network/chain/protocol/x/stats/keeper"
	"github.com/bitoro-network/chain/protocol/x/stats/types"
	"github.com/cosmos/cosmos-sdk/codec"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

func createStatsKeeper(
	stateStore storetypes.CommitMultiStore,
	epochsKeeper *epochskeeper.Keeper,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
	stakingKeeper *stakingkeeper.Keeper,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	transientStoreKey := storetypes.NewTransientStoreKey(types.TransientStoreKey)

	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(transientStoreKey, storetypes.StoreTypeTransient, db)

	mockMsgSender := &mocks.IndexerMessageSender{}
	mockMsgSender.On("Enabled").Return(true)

	authorities := []string{
		delaymsgtypes.ModuleAddress.String(),
		lib.GovModuleAddress.String(),
	}
	k := keeper.NewKeeper(
		cdc,
		epochsKeeper,
		storeKey,
		transientStoreKey,
		authorities,
		stakingKeeper,
	)

	return k, storeKey
}
