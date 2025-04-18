package types_test

import (
	"testing"

	"github.com/bitoro-network/chain/protocol/x/feetiers/types"
	"github.com/stretchr/testify/require"
)

func TestModuleKeys(t *testing.T) {
	require.Equal(t, "feetiers", types.ModuleName)
	require.Equal(t, "feetiers", types.StoreKey)
}

func TestStateKeys(t *testing.T) {
	require.Equal(t, "PerpParams", types.PerpetualFeeParamsKey)
}
