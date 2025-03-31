package keeper

import (
	"fmt"

	"cosmossdk.io/store/prefix"
	"github.com/Bitoro-Network/chain/protocol/x/subaccounts/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetSafetyHeapStore returns the safety heap store.
func (k Keeper) GetSafetyHeapStore(
	ctx sdk.Context,
	perpetualId uint32,
	side types.SafetyHeapPositionSide,
) prefix.Store {
	return prefix.NewStore(
		ctx.KVStore(k.storeKey),
		k.GetSafetyHeapKeyPrefix(perpetualId, side),
	)
}

// GetSafetyHeapKeyPrefix returns the prefix for the safety heap store.
func (k Keeper) GetSafetyHeapKeyPrefix(
	perpetualId uint32,
	side types.SafetyHeapPositionSide,
) []byte {
	return []byte(
		fmt.Sprintf(
			"%s/%d/%d/",
			types.SafetyHeapStorePrefix,
			perpetualId,
			side,
		),
	)
}
