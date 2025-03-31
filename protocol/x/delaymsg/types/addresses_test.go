package types_test

import (
	"testing"

	"github.com/Bitoro-Network/chain/protocol/x/delaymsg/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "btoro1mkkvp26dngu6n8rmalaxyp3gwkjuzztq5zx6tr", types.ModuleAddress.String())
}
