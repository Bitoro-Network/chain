package keeper

import (
	"github.com/bitoro-network/chain/protocol/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `GetBridgeEventFromServer` returns the bridge event with the given id from the server. `found` is false
// if the event is not found.
func (k Keeper) GetBridgeEventFromServer(ctx sdk.Context, id uint32) (event types.BridgeEvent, found bool) {
	event, _, found = k.bridgeEventManager.GetBridgeEventById(id)
	return event, found
}
