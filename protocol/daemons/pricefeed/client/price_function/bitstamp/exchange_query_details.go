package bitstamp

import (
	"github.com/bitoro-network/chain/protocol/daemons/pricefeed/client/constants/exchange_common"
	"github.com/bitoro-network/chain/protocol/daemons/pricefeed/client/types"
)

var (
	BitstampDetails = types.ExchangeQueryDetails{
		Exchange:      exchange_common.EXCHANGE_ID_BITSTAMP,
		Url:           "https://www.bitstamp.net/api/v2/ticker/",
		PriceFunction: BitstampPriceFunction,
		IsMultiMarket: true,
	}
)
