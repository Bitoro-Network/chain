package indexer_manager

import "github.com/bitoro-network/chain/protocol/indexer/msgsender"

func NewIndexerEventManagerNoop() IndexerEventManager {
	return NewIndexerEventManager(
		msgsender.NewIndexerMessageSenderNoop(),
		nil,
		false,
	)
}

func NewIndexerEventManagerNoopEnabled() IndexerEventManager {
	return NewIndexerEventManager(
		msgsender.NewIndexerMessageSenderNoopEnabled(),
		nil,
		false,
	)
}
