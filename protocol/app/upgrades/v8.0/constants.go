package v_8_0

import (
	store "cosmossdk.io/store/types"
	"github.com/bitoro-network/chain/protocol/app/upgrades"
)

const (
	UpgradeName = "v8.0"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:   UpgradeName,
	StoreUpgrades: store.StoreUpgrades{},
}
