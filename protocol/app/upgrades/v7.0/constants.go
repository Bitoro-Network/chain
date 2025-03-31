package v_7_0

import (
	store "cosmossdk.io/store/types"
	"github.com/Bitoro-Network/chain/protocol/app/upgrades"
	affiliatetypes "github.com/Bitoro-Network/chain/protocol/x/affiliates/types"
)

const (
	UpgradeName = "v7.0"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			affiliatetypes.StoreKey,
		},
	},
}
