package ante

import (
	errorsmod "cosmossdk.io/errors"
	accountpluskeeper "github.com/bitoro-network/chain/protocol/x/accountplus/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

// IsTimestampNonceTx returns `true` if the supplied `tx` consist of a single signature that uses a timestamp nonce
// value for sequence
func IsTimestampNonceTx(ctx sdk.Context, tx sdk.Tx) (bool, error) {
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return false, errorsmod.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}
	signatures, err := sigTx.GetSignaturesV2()
	if err != nil {
		return false, err
	}

	// multi signature cannot contain timestamp nonce
	if len(signatures) > 1 {
		for _, sig := range signatures {
			if accountpluskeeper.IsTimestampNonce(sig.Sequence) {
				return false, errorsmod.Wrap(sdkerrors.ErrTxDecode, "multi signature contains timestampnonce")
			}
		}
	}

	return len(signatures) == 1 && accountpluskeeper.IsTimestampNonce(signatures[0].Sequence), nil
}
