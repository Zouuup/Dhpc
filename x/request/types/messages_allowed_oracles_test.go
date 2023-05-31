package types

import (
	"testing"

	"Dhpc/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateAllowedOracles_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateAllowedOracles
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateAllowedOracles{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateAllowedOracles{
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

func TestMsgUpdateAllowedOracles_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateAllowedOracles
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateAllowedOracles{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateAllowedOracles{
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

func TestMsgDeleteAllowedOracles_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteAllowedOracles
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteAllowedOracles{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteAllowedOracles{
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
