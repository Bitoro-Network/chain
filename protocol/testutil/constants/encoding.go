package constants

import "github.com/bitoro-network/chain/protocol/testutil/encoding"

var (
	TestEncodingCfg = encoding.GetTestEncodingCfg()
	TestTxBuilder   = TestEncodingCfg.TxConfig.NewTxBuilder()
)
