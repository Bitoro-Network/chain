package kraken

import (
	"github.com/bitoro-network/chain/protocol/daemons/pricefeed/client/constants/exchange_common"
	"github.com/bitoro-network/chain/protocol/daemons/pricefeed/client/types"
)

var (
	KrakenDetails = types.ExchangeQueryDetails{
		Exchange:      exchange_common.EXCHANGE_ID_KRAKEN,
		Url:           "https://api.kraken.com/0/public/Ticker",
		PriceFunction: KrakenPriceFunction,
		IsMultiMarket: true,
	}
)
