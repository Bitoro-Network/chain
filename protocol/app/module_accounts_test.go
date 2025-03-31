package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bitoro-network/chain/protocol/app"
	bridgemoduletypes "github.com/bitoro-network/chain/protocol/x/bridge/types"
	perpetualsmoduletypes "github.com/bitoro-network/chain/protocol/x/perpetuals/types"
	rewardsmoduletypes "github.com/bitoro-network/chain/protocol/x/rewards/types"
	satypes "github.com/bitoro-network/chain/protocol/x/subaccounts/types"
	vaultmoduletypes "github.com/bitoro-network/chain/protocol/x/vault/types"
	vestmoduletypes "github.com/bitoro-network/chain/protocol/x/vest/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	marketmapmoduletypes "github.com/skip-mev/slinky/x/marketmap/types"
)

func TestModuleAccountsToAddresses(t *testing.T) {
	expectedModuleAccToAddresses := map[string]string{
		authtypes.FeeCollectorName:                   "btoro17xpfvakm2amg962yls6f84z3kell8c5leqdyt2",
		bridgemoduletypes.ModuleName:                 "btoro1zlefkpe3g0vvm9a4h0jf9000lmqutlh9jwjnsv",
		distrtypes.ModuleName:                        "btoro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8wx2cfg",
		stakingtypes.BondedPoolName:                  "btoro1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3uz8teq",
		stakingtypes.NotBondedPoolName:               "btoro1tygms3xhhs3yv487phx3dw4a95jn7t7lgzm605",
		govtypes.ModuleName:                          "btoro10d07y265gmmuvt4z0w9aw880jnsr700jnmapky",
		ibctransfertypes.ModuleName:                  "btoro1yl6hdjhmkf37639730gffanpzndzdpmh8xcdh5",
		satypes.ModuleName:                           "btoro1v88c3xv9xyv3eetdx0tvcmq7ung3dywp5upwc6",
		perpetualsmoduletypes.InsuranceFundName:      "btoro1c7ptc87hkd54e3r7zjy92q29xkq7t79w64slrq",
		rewardsmoduletypes.TreasuryAccountName:       "btoro16wrau2x4tsg033xfrrdpae6kxfn9kyuerr5jjp",
		rewardsmoduletypes.VesterAccountName:         "btoro1ltyc6y4skclzafvpznpt2qjwmfwgsndp458rmp",
		vestmoduletypes.CommunityTreasuryAccountName: "btoro15ztc7xy42tn2ukkc0qjthkucw9ac63pgp70urn",
		vestmoduletypes.CommunityVesterAccountName:   "btoro1wxje320an3karyc6mjw4zghs300dmrjkwn7xtk",
		icatypes.ModuleName:                          "btoro1vlthgax23ca9syk7xgaz347xmf4nunefw3cnv8",
		marketmapmoduletypes.ModuleName:              "btoro16j3d86dww8p2rzdlqsv7wle98cxzjxw6gjjyzn",
		vaultmoduletypes.MegavaultAccountName:        "btoro18tkxrnrkqc2t0lr3zxr5g6a4hdvqksylxqje4r",
	}

	require.True(t, len(expectedModuleAccToAddresses) == len(app.GetMaccPerms()),
		"expected %d, got %d", len(expectedModuleAccToAddresses), len(app.GetMaccPerms()))
	for acc, address := range expectedModuleAccToAddresses {
		expectedAddr := authtypes.NewModuleAddress(acc).String()
		require.Equal(t, address, expectedAddr, "module (%v) should have address (%s)", acc, expectedAddr)
	}
}

func TestBlockedAddresses(t *testing.T) {
	expectedBlockedAddresses := map[string]bool{
		"btoro17xpfvakm2amg962yls6f84z3kell8c5leqdyt2": true,
		"btoro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8wx2cfg": true,
		"btoro1tygms3xhhs3yv487phx3dw4a95jn7t7lgzm605": true,
		"btoro1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3uz8teq": true,
		"btoro1yl6hdjhmkf37639730gffanpzndzdpmh8xcdh5": true,
		"btoro1vlthgax23ca9syk7xgaz347xmf4nunefw3cnv8": true,
	}
	require.Equal(t, expectedBlockedAddresses, app.BlockedAddresses())
}

func TestMaccPerms(t *testing.T) {
	maccPerms := app.GetMaccPerms()
	expectedMaccPerms := map[string][]string{
		"bonded_tokens_pool":     {"burner", "staking"},
		"bridge":                 {"minter"},
		"distribution":           nil,
		"fee_collector":          nil,
		"gov":                    {"burner"},
		"insurance_fund":         nil,
		"not_bonded_tokens_pool": {"burner", "staking"},
		"subaccounts":            nil,
		"transfer":               {"minter", "burner"},
		"interchainaccounts":     nil,
		"rewards_treasury":       nil,
		"rewards_vester":         nil,
		"community_treasury":     nil,
		"community_vester":       nil,
		"marketmap":              nil,
		"megavault":              nil,
	}
	require.Equal(t, expectedMaccPerms, maccPerms, "default macc perms list does not match expected")
}

func TestModuleAccountAddrs(t *testing.T) {
	expectedModuleAccAddresses := map[string]bool{
		"btoro17xpfvakm2amg962yls6f84z3kell8c5leqdyt2": true, // x/auth.FeeCollector
		"btoro1zlefkpe3g0vvm9a4h0jf9000lmqutlh9jwjnsv": true, // x/bridge
		"btoro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8wx2cfg": true, // x/distribution
		"btoro1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3uz8teq": true, // x/staking.bondedPool
		"btoro1tygms3xhhs3yv487phx3dw4a95jn7t7lgzm605": true, // x/staking.notBondedPool
		"btoro10d07y265gmmuvt4z0w9aw880jnsr700jnmapky": true, // x/ gov
		"btoro1yl6hdjhmkf37639730gffanpzndzdpmh8xcdh5": true, // ibc transfer
		"btoro1vlthgax23ca9syk7xgaz347xmf4nunefw3cnv8": true, // interchainaccounts
		"btoro1v88c3xv9xyv3eetdx0tvcmq7ung3dywp5upwc6": true, // x/subaccount
		"btoro1c7ptc87hkd54e3r7zjy92q29xkq7t79w64slrq": true, // x/clob.insuranceFund
		"btoro16wrau2x4tsg033xfrrdpae6kxfn9kyuerr5jjp": true, // x/rewards.treasury
		"btoro1ltyc6y4skclzafvpznpt2qjwmfwgsndp458rmp": true, // x/rewards.vester
		"btoro15ztc7xy42tn2ukkc0qjthkucw9ac63pgp70urn": true, // x/vest.communityTreasury
		"btoro1wxje320an3karyc6mjw4zghs300dmrjkwn7xtk": true, // x/vest.communityVester
		"btoro16j3d86dww8p2rzdlqsv7wle98cxzjxw6gjjyzn": true, // x/marketmap
		"btoro18tkxrnrkqc2t0lr3zxr5g6a4hdvqksylxqje4r": true, // x/vault.megavault
	}

	require.Equal(t, expectedModuleAccAddresses, app.ModuleAccountAddrs())
}
