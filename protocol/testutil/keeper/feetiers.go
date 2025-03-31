package keeper

import (
	"github.com/Bitoro-Network/chain/protocol/lib"
	"github.com/Bitoro-Network/chain/protocol/mocks"
	dbm "github.com/cosmos/cosmos-db"

	storetypes "cosmossdk.io/store/types"
	affiliateskeeper "github.com/Bitoro-Network/chain/protocol/x/affiliates/keeper"
	delaymsgtypes "github.com/Bitoro-Network/chain/protocol/x/delaymsg/types"
	"github.com/Bitoro-Network/chain/protocol/x/feetiers/keeper"
	"github.com/Bitoro-Network/chain/protocol/x/feetiers/types"
	statskeeper "github.com/Bitoro-Network/chain/protocol/x/stats/keeper"
	vaultkeeper "github.com/Bitoro-Network/chain/protocol/x/vault/keeper"
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
