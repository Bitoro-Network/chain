package lib_test

import (
	"testing"

	"github.com/bitoro-network/chain/protocol/lib"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMustParseCoinsNormalized(t *testing.T) {
	// Test case: valid coin string
	coinStr := "100atom"
	expectedCoins := sdk.Coins{sdk.NewInt64Coin("atom", 100)}
	result := lib.MustParseCoinsNormalized(coinStr)
	require.Equal(t, expectedCoins, result)

	// Test case: invalid coin string
	invalidCoinStr := "invalidcoin"
	require.Panics(t, func() { lib.MustParseCoinsNormalized(invalidCoinStr) })
}
