package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTreasury = "create_treasury"
	TypeMsgUpdateTreasury = "update_treasury"
	TypeMsgDeleteTreasury = "delete_treasury"
)

var _ sdk.Msg = &MsgCreateTreasury{}

func NewMsgCreateTreasury(creator string, address string) *MsgCreateTreasury {
	return &MsgCreateTreasury{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgCreateTreasury) Route() string {
	return RouterKey
}

func (msg *MsgCreateTreasury) Type() string {
	return TypeMsgCreateTreasury
}

func (msg *MsgCreateTreasury) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTreasury) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTreasury) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTreasury{}

func NewMsgUpdateTreasury(creator string, address string) *MsgUpdateTreasury {
	return &MsgUpdateTreasury{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgUpdateTreasury) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTreasury) Type() string {
	return TypeMsgUpdateTreasury
}

func (msg *MsgUpdateTreasury) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTreasury) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTreasury) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTreasury{}

func NewMsgDeleteTreasury(creator string) *MsgDeleteTreasury {
	return &MsgDeleteTreasury{
		Creator: creator,
	}
}
func (msg *MsgDeleteTreasury) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTreasury) Type() string {
	return TypeMsgDeleteTreasury
}

func (msg *MsgDeleteTreasury) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTreasury) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTreasury) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
