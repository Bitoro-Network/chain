package lib_test

import (
	"testing"

	"github.com/Bitoro-Network/chain/protocol/lib"
	"github.com/stretchr/testify/require"
)

func TestGovModuleAddress(t *testing.T) {
	require.Equal(t, "btoro10d07y265gmmuvt4z0w9aw880jnsr700jnmapky", lib.GovModuleAddress.String())
}
