package keeper

import (
	"github.com/bitoro-network/chain/protocol/lib"
	"github.com/bitoro-network/chain/protocol/mocks"
	dbm "github.com/cosmos/cosmos-db"

	storetypes "cosmossdk.io/store/types"
	affiliateskeeper "github.com/bitoro-network/chain/protocol/x/affiliates/keeper"
	delaymsgtypes "github.com/bitoro-network/chain/protocol/x/delaymsg/types"
	"github.com/bitoro-network/chain/protocol/x/feetiers/keeper"
	"github.com/bitoro-network/chain/protocol/x/feetiers/types"
	statskeeper "github.com/bitoro-network/chain/protocol/x/stats/keeper"
	vaultkeeper "github.com/bitoro-network/chain/protocol/x/vault/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
)

func createFeeTiersKeeper(
	stateStore storetypes.CommitMultiStore,
	statsKeeper *statskeeper.Keeper,
	vaultKeeper *vaultkeeper.Keeper,
	affiliatesKeeper *affiliateskeeper.Keeper,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	mockMsgSender := &mocks.IndexerMessageSender{}
	mockMsgSender.On("Enabled").Return(true)

	authorities := []string{
		delaymsgtypes.ModuleAddress.String(),
		lib.GovModuleAddress.String(),
	}
	k := keeper.NewKeeper(
		cdc,
		statsKeeper,
		affiliatesKeeper,
		storeKey,
		authorities,
	)
	k.SetVaultKeeper(vaultKeeper)

	return k, storeKey
}
