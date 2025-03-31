package coinbase_pro

import (
	"github.com/Bitoro-Network/chain/protocol/daemons/pricefeed/client/constants/exchange_common"
	"github.com/Bitoro-Network/chain/protocol/daemons/pricefeed/client/types"
)

var (
	CoinbaseProDetails = types.ExchangeQueryDetails{
		Exchange:      exchange_common.EXCHANGE_ID_COINBASE_PRO,
		Url:           "https://api.pro.coinbase.com/products/$/ticker",
		PriceFunction: CoinbaseProPriceFunction,
	}
)
