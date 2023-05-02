package types

import (
	"testing"

	"Decent/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgAddLinkedRequester_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddLinkedRequester
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddLinkedRequester{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddLinkedRequester{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
