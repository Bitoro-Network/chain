package types_test

import (
	"testing"

	"github.com/Bitoro-Network/chain/protocol/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "btoro1zlefkpe3g0vvm9a4h0jf9000lmqutlh9jwjnsv", types.ModuleAddress.String())
}
