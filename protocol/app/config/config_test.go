package config_test

import (
	"testing"

	"github.com/bitoro-network/chain/protocol/app/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestSetupConfig_SealsConfig(t *testing.T) {
	sdkConfig := sdk.GetConfig()

	// A successful set confirms the config is not yet sealed
	sdkConfig.SetPurpose(0)
	require.Equal(t, uint32(0), sdkConfig.GetPurpose(), "Expected purpose to match set value")

	// Should set default app values and seal the config
	config.SetupConfig()

	require.Panicsf(t, func() { sdkConfig.SetPurpose(0) }, "Expected config to be sealed after SetupConfig")
}

func TestSetAddressPrefixes(t *testing.T) {
	sdkConfig := sdk.GetConfig()

	require.Equal(t, "btoro", sdkConfig.GetBech32AccountAddrPrefix())
	require.Equal(t, "btoropub", sdkConfig.GetBech32AccountPubPrefix())

	require.Equal(t, "btorovaloper", sdkConfig.GetBech32ValidatorAddrPrefix())
	require.Equal(t, "btorovaloperpub", sdkConfig.GetBech32ValidatorPubPrefix())

	require.Equal(t, "btorovalcons", sdkConfig.GetBech32ConsensusAddrPrefix())
	require.Equal(t, "btorovalconspub", sdkConfig.GetBech32ConsensusPubPrefix())
}
