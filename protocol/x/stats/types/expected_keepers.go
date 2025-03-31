package types

import (
	"context"

	epochtypes "github.com/bitoro-network/chain/protocol/x/epochs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// EpochsKeeper defines the expected epochs keeper to get epoch info.
type EpochsKeeper interface {
	MustGetStatsEpochInfo(ctx sdk.Context) epochtypes.EpochInfo
}

type StakingKeeper interface {
	GetDelegatorDelegations(ctx context.Context,
		delegator sdk.AccAddress, maxRetrieve uint16) ([]stakingtypes.Delegation, error)
}
