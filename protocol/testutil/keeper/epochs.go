package keeper

import (
	"testing"

	dbm "github.com/cosmos/cosmos-db"

	storetypes "cosmossdk.io/store/types"
	"github.com/bitoro-network/chain/protocol/x/epochs/keeper"
	"github.com/bitoro-network/chain/protocol/x/epochs/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TestEpochInfoName           = "name"
	TestEpochDuration           = uint32(20)
	TestCreateEpochBlockTimeSec = 1656900000
)

func EpochsKeeper(
	t testing.TB,
) (
	ctx sdk.Context,
	epochsKeeper *keeper.Keeper,
	storeKey storetypes.StoreKey,
) {
	ctx = initKeepers(t, func(
		db *dbm.MemDB,
		registry codectypes.InterfaceRegistry,
		cdc *codec.ProtoCodec,
		stateStore storetypes.CommitMultiStore,
		transientStoreKey storetypes.StoreKey,
	) []GenesisInitializer {
		epochsKeeper, storeKey = createEpochsKeeper(
			stateStore,
			db,
			cdc,
		)

		return []GenesisInitializer{epochsKeeper}
	})

	return ctx, epochsKeeper, storeKey
}

func createEpochsKeeper(
	stateStore storetypes.CommitMultiStore,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
	)

	return k, storeKey
}
