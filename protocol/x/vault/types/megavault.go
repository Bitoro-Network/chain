package types

import (
	satypes "github.com/bitoro-network/chain/protocol/x/subaccounts/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

var (
	MegavaultMainAddress = authtypes.NewModuleAddress(MegavaultAccountName)
	// MegavaultMainSubaccount is subaccount 0 of the module account derived from string "megavault".
	MegavaultMainSubaccount = satypes.SubaccountId{
		Owner:  MegavaultMainAddress.String(),
		Number: 0,
	}
)
