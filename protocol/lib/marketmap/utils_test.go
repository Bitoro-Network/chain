package marketmap_test

import (
	"testing"

	"github.com/Bitoro-Network/chain/protocol/lib/marketmap"
	pricestypes "github.com/Bitoro-Network/chain/protocol/x/prices/types"
	slinkytypes "github.com/skip-mev/slinky/pkg/types"
	marketmaptypes "github.com/skip-mev/slinky/x/marketmap/types"
	"github.com/stretchr/testify/require"
)

func TestConstructMarketMapFromParams(t *testing.T) {
	marketParams := []pricestypes.MarketParam{
		{
			Pair:               "BTC-USD",
			Exponent:           -8,
			MinExchanges:       1,
			MinPriceChangePpm:  10,
			ExchangeConfigJson: `{"exchanges":[{"exchangeName":"Binance","ticker":"BTCUSDT"}]}`,
		},
	}

	expectedMarketMap := marketmaptypes.MarketMap{
		Markets: map[string]marketmaptypes.Market{
			"BTC/USD": {
				Ticker: marketmaptypes.Ticker{
					CurrencyPair:     slinkytypes.CurrencyPair{Base: "BTC", Quote: "USD"},
					Decimals:         8,
					MinProviderCount: 1,
					Enabled:          true,
					Metadata_JSON:    "",
				},
				ProviderConfigs: []marketmaptypes.ProviderConfig{
					{Name: "binance_ws", OffChainTicker: "BTCUSDT"},
				},
			},
		},
	}
	marketMap, err := marketmap.ConstructMarketMapFromParams(marketParams)
	require.NoError(t, err)
	require.NotNil(t, marketMap)
	require.Equal(t, expectedMarketMap, marketMap)
}
