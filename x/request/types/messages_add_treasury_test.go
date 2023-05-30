package types

import (
	"testing"

	"Decent/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateAddTreasury_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateAddTreasury
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateAddTreasury{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateAddTreasury{
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

func TestMsgUpdateAddTreasury_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateAddTreasury
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateAddTreasury{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateAddTreasury{
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

func TestMsgDeleteAddTreasury_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteAddTreasury
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteAddTreasury{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteAddTreasury{
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
