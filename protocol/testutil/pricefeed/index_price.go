package pricefeed

import (
	"testing"
	"time"

	"github.com/bitoro-network/chain/protocol/testutil/constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/bitoro-network/chain/protocol/app"
	pricefeedapi "github.com/bitoro-network/chain/protocol/daemons/pricefeed/api"
)

func GetTestMarketPriceUpdates(n int) (indexPrices []*pricefeedapi.MarketPriceUpdate) {
	for i := 0; i < n; i++ {
		indexPrices = append(
			indexPrices,
			&pricefeedapi.MarketPriceUpdate{
				MarketId: uint32(i),
				ExchangePrices: []*pricefeedapi.ExchangePrice{
					constants.Exchange1_Price1_TimeT,
					constants.Exchange2_Price2_TimeT,
				},
			},
		)
	}
	return indexPrices
}

func UpdateIndexPrice(
	t testing.TB,
	ctx sdk.Context,
	tApp *app.App,
	marketId uint32,
	price uint64,
	lastUpdatedTime time.Time,
) {
	_, err := tApp.Server.UpdateMarketPrices(
		ctx,
		&pricefeedapi.UpdateMarketPricesRequest{
			MarketPriceUpdates: []*pricefeedapi.MarketPriceUpdate{
				{
					MarketId: marketId,
					ExchangePrices: []*pricefeedapi.ExchangePrice{
						{
							ExchangeId:     "exchange-a",
							Price:          price,
							LastUpdateTime: &lastUpdatedTime,
						},
						{
							ExchangeId:     "exchange-b",
							Price:          price,
							LastUpdateTime: &lastUpdatedTime,
						},
						{
							ExchangeId:     "exchange-c",
							Price:          price,
							LastUpdateTime: &lastUpdatedTime,
						},
					},
				},
			},
		},
	)
	require.NoError(t, err)
}
