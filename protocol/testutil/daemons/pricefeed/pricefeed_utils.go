package pricefeed

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/exp/maps"

	"github.com/bitoro-network/chain/protocol/testutil/json"
	"github.com/stretchr/testify/require"
)

// ErrorMapsEqual is a testing method that takes any two maps of keys to errors and asserts that they have the same
// sets of keys, and that each associated error value has the same rendered message.
func ErrorMapsEqual[K comparable](t testing.TB, expected map[K]error, actual map[K]error) {
	require.Equal(t, len(expected), len(actual))
	for key, expectedError := range expected {
		error, ok := actual[key]
		require.True(t, ok)
		require.EqualError(t, error, expectedError.Error())
	}
}

// ErrorsEqual is a testing method that takes any two slices of errors and asserts that each actual error has
// the same rendered message as the expected error.
func ErrorsEqual(t testing.TB, expected []error, actual []error) {
	require.Equal(t, len(expected), len(actual))
	for i, expectedError := range expected {
		require.EqualError(t, expectedError, actual[i].Error())
	}
}

// ReadJsonTestFile takes a test file with human-readable, formatted JSON, load it, and compacts it.
// The purpose is to remove the formatting (e.g. newlines, tabs, etc) and return a string that would match an
// unmarshaled object string generated by a Go program natively.
func ReadJsonTestFile(t testing.TB, fileName string) string {
	fileBytes, err := os.ReadFile(fmt.Sprintf("testdata/%v", fileName))
	require.NoError(t, err, "Error reading test file")
	return json.CompactJsonString(t, string(fileBytes))
}

// MarketParamErrorsEqual is a testing method that takes any two maps of market ids to errors and asserts that they
// have the same sets of keys, and that each associated error value has the same rendered message
func MarketParamErrorsEqual(
	t testing.TB,
	expectedMarketParamErrors map[uint32]error,
	actualMarketParamErrors map[uint32]error,
) {
	require.Equal(t, maps.Keys(expectedMarketParamErrors), maps.Keys(actualMarketParamErrors))
	for marketId, expectedErr := range expectedMarketParamErrors {
		actualErr := actualMarketParamErrors[marketId]
		require.ErrorContains(
			t,
			actualErr,
			expectedErr.Error(),
			"Errors for market id %v do not match",
			marketId,
		)
	}
}
