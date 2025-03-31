package grpc

import (
	bridgetypes "github.com/Bitoro-Network/chain/protocol/daemons/bridge/api"
	liquidationtypes "github.com/Bitoro-Network/chain/protocol/daemons/liquidation/api"
	pricefeedtypes "github.com/Bitoro-Network/chain/protocol/daemons/pricefeed/api"
	blocktimetypes "github.com/Bitoro-Network/chain/protocol/x/blocktime/types"
	clobtypes "github.com/Bitoro-Network/chain/protocol/x/clob/types"
	perptypes "github.com/Bitoro-Network/chain/protocol/x/perpetuals/types"
	pricetypes "github.com/Bitoro-Network/chain/protocol/x/prices/types"
	satypes "github.com/Bitoro-Network/chain/protocol/x/subaccounts/types"
)

// QueryClient combines all the query clients used in testing into a single mock interface for testing convenience.
type QueryClient interface {
	blocktimetypes.QueryClient
	satypes.QueryClient
	clobtypes.QueryClient
	perptypes.QueryClient
	pricetypes.QueryClient
	bridgetypes.BridgeServiceClient
	liquidationtypes.LiquidationServiceClient
	pricefeedtypes.PriceFeedServiceClient
}
