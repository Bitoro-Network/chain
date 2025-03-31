package ws

import (
	"github.com/gorilla/websocket"

	"github.com/Bitoro-Network/chain/protocol/streaming/types"
	clobtypes "github.com/Bitoro-Network/chain/protocol/x/clob/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

var _ types.OutgoingMessageSender = (*WebsocketMessageSender)(nil)

type WebsocketMessageSender struct {
	cdc codec.JSONCodec

	conn *websocket.Conn
}

func (wms *WebsocketMessageSender) Send(
	response *clobtypes.StreamOrderbookUpdatesResponse,
) (err error) {
	responseJson, err := wms.cdc.MarshalJSON(response)
	if err != nil {
		return err
	}
	return wms.conn.WriteMessage(websocket.TextMessage, responseJson)
}
