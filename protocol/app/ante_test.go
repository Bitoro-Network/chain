package app_test

import (
	"testing"

	"cosmossdk.io/store/rootmulti"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/Bitoro-Network/chain/protocol/app"
	testApp "github.com/Bitoro-Network/chain/protocol/testutil/app"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/stretchr/testify/require"
)

func newHandlerOptions() app.HandlerOptions {
	encodingConfig := app.GetEncodingConfig()
	bitoroApp := testApp.DefaultTestApp(nil)
	return app.HandlerOptions{
		HandlerOptions: ante.HandlerOptions{
			AccountKeeper:   bitoroApp.AccountKeeper,
			BankKeeper:      bitoroApp.BankKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:  bitoroApp.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
		AccountplusKeeper: &bitoroApp.AccountPlusKeeper,
		ClobKeeper:        bitoroApp.ClobKeeper,
		Codec:             encodingConfig.Codec,
		AuthStoreKey:      bitoroApp.CommitMultiStore().(*rootmulti.Store).StoreKeysByName()[authtypes.StoreKey],
		PerpetualsKeeper:  bitoroApp.PerpetualsKeeper,
		PricesKeeper:      bitoroApp.PricesKeeper,
		MarketMapKeeper:   &bitoroApp.MarketMapKeeper,
	}
}

func TestNewAnteHandler(t *testing.T) {
	handlerOptions := newHandlerOptions()
	anteHandler, err := app.NewAnteHandler(handlerOptions)
	require.NoError(t, err, "NewAnteHandler call failed")
	require.NotNil(t, anteHandler, "expected non-nil AnteHandler function")
}

func TestNewAnteHandler_Error(t *testing.T) {
	tests := map[string]struct {
		handlerMutation func(*app.HandlerOptions)
		errorMsg        string
	}{
		"nil handlerOptions.AccountKeeper": {
			handlerMutation: func(options *app.HandlerOptions) { options.AccountKeeper = nil },
			errorMsg:        "account keeper is required for ante builder",
		},
		"nil handlerOptions.BankKeeper": {
			handlerMutation: func(options *app.HandlerOptions) { options.BankKeeper = nil },
			errorMsg:        "bank keeper is required for ante builder",
		},
		"nil PerpetualsKeeper": {
			handlerMutation: func(options *app.HandlerOptions) { options.PerpetualsKeeper = nil },
			errorMsg:        "perpetuals keeper is required for ante builder",
		},
		"nil PricesKeeper": {
			handlerMutation: func(options *app.HandlerOptions) { options.PricesKeeper = nil },
			errorMsg:        "prices keeper is required for ante builder",
		},
		"nil MarketMapKeeper": {
			handlerMutation: func(options *app.HandlerOptions) { options.MarketMapKeeper = nil },
			errorMsg:        "market map keeper is required for ante builder",
		},
		"nil handlerOptions.SignModeHandler": {
			handlerMutation: func(options *app.HandlerOptions) { options.SignModeHandler = nil },
			errorMsg:        "sign mode handler is required for ante builder",
		},
		"nil ClobKeeper": {
			handlerMutation: func(options *app.HandlerOptions) { options.ClobKeeper = nil },
			errorMsg:        "clob keeper is required for ante builder",
		},
		"nil Codec": {
			handlerMutation: func(options *app.HandlerOptions) { options.Codec = nil },
			errorMsg:        "codec is required for ante builder",
		},
		"nil AuthStoreKey": {
			handlerMutation: func(options *app.HandlerOptions) { options.AuthStoreKey = nil },
			errorMsg:        "auth store key is required for ante builder",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handlerOptions := newHandlerOptions()
			tc.handlerMutation(&handlerOptions)

			anteHandler, err := app.NewAnteHandler(handlerOptions)
			require.Nil(t, anteHandler, "Expected Ante Handler creation to error")
			require.Errorf(t, err, tc.errorMsg)
		})
	}
}
