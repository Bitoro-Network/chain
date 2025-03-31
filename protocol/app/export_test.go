package app_test

import (
	"testing"

	"github.com/Bitoro-Network/chain/protocol/testutil/app"
	"github.com/stretchr/testify/require"
)

func TestExportAppStateAndValidators_Panics(t *testing.T) {
	bitoroApp := app.DefaultTestApp(nil)
	require.Panics(t, func() { bitoroApp.ExportAppStateAndValidators(false, nil, nil) }) // nolint:errcheck
}
