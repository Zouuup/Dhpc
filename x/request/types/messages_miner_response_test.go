package types

import (
	"testing"

	"Dhpc/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateMinerResponse_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateMinerResponse
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateMinerResponse{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateMinerResponse{
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

func TestMsgUpdateMinerResponse_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateMinerResponse
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateMinerResponse{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateMinerResponse{
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

func TestMsgDeleteMinerResponse_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteMinerResponse
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteMinerResponse{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteMinerResponse{
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
