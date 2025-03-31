package marketmap

import (
	pricestypes "github.com/Bitoro-Network/chain/protocol/x/prices/types"
	bitoro "github.com/skip-mev/slinky/providers/apis/dydx"
	bitorotypes "github.com/skip-mev/slinky/providers/apis/dydx/types"
	marketmaptypes "github.com/skip-mev/slinky/x/marketmap/types"
)

// Construct a MarketMap struct from a slice of MarketParams
func ConstructMarketMapFromParams(
	allMarketParams []pricestypes.MarketParam,
) (marketmaptypes.MarketMap, error) {
	var mpr bitorotypes.QueryAllMarketParamsResponse
	for _, mp := range allMarketParams {
		mpr.MarketParams = append(mpr.MarketParams, bitorotypes.MarketParam{
			Id:                 mp.Id,
			Pair:               mp.Pair,
			Exponent:           mp.Exponent,
			MinExchanges:       mp.MinExchanges,
			MinPriceChangePpm:  mp.MinPriceChangePpm,
			ExchangeConfigJson: mp.ExchangeConfigJson,
		})
	}
	mm, err := bitoro.ConvertMarketParamsToMarketMap(mpr)
	if err != nil {
		return marketmaptypes.MarketMap{}, err
	}

	return mm.MarketMap, nil
}
