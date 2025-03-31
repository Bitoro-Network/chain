package types_test

import (
	"testing"

	"github.com/Bitoro-Network/chain/protocol/x/govplus/types"
	"github.com/stretchr/testify/require"
)

func TestModuleKeys(t *testing.T) {
	require.Equal(t, "govplus", types.ModuleName)
	require.Equal(t, "bitorogovplus", types.StoreKey)
}
