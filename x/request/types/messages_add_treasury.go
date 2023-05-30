package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAddTreasury = "create_add_treasury"
	TypeMsgUpdateAddTreasury = "update_add_treasury"
	TypeMsgDeleteAddTreasury = "delete_add_treasury"
)

var _ sdk.Msg = &MsgCreateAddTreasury{}

func NewMsgCreateAddTreasury(creator string) *MsgCreateAddTreasury {
	return &MsgCreateAddTreasury{
		Creator: creator,
	}
}

func (msg *MsgCreateAddTreasury) Route() string {
	return RouterKey
}

func (msg *MsgCreateAddTreasury) Type() string {
	return TypeMsgCreateAddTreasury
}

func (msg *MsgCreateAddTreasury) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAddTreasury) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAddTreasury) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAddTreasury{}

func NewMsgUpdateAddTreasury(creator string, id uint64) *MsgUpdateAddTreasury {
	return &MsgUpdateAddTreasury{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgUpdateAddTreasury) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAddTreasury) Type() string {
	return TypeMsgUpdateAddTreasury
}

func (msg *MsgUpdateAddTreasury) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAddTreasury) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAddTreasury) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAddTreasury{}

func NewMsgDeleteAddTreasury(creator string, id uint64) *MsgDeleteAddTreasury {
	return &MsgDeleteAddTreasury{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteAddTreasury) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAddTreasury) Type() string {
	return TypeMsgDeleteAddTreasury
}

func (msg *MsgDeleteAddTreasury) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAddTreasury) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAddTreasury) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
