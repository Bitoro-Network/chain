package types_test

import (
	"testing"

	"github.com/Bitoro-Network/chain/protocol/x/subaccounts/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "btoro1v88c3xv9xyv3eetdx0tvcmq7ung3dywp5upwc6", types.ModuleAddress.String())
}
