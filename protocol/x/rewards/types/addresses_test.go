package types_test

import (
	"testing"

	"github.com/Bitoro-Network/chain/protocol/x/rewards/types"
	"github.com/stretchr/testify/require"
)

func TestTreasuryModuleAddress(t *testing.T) {
	require.Equal(t, "btoro16wrau2x4tsg033xfrrdpae6kxfn9kyuerr5jjp", types.TreasuryModuleAddress.String())
}
