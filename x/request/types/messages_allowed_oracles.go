package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAllowedOracles = "create_allowed_oracles"
	TypeMsgUpdateAllowedOracles = "update_allowed_oracles"
	TypeMsgDeleteAllowedOracles = "delete_allowed_oracles"
)

var _ sdk.Msg = &MsgCreateAllowedOracles{}

func NewMsgCreateAllowedOracles(creator string, name string, address string) *MsgCreateAllowedOracles {
	return &MsgCreateAllowedOracles{
		Creator: creator,
		Name:    name,
		Address: address,
	}
}

func (msg *MsgCreateAllowedOracles) Route() string {
	return RouterKey
}

func (msg *MsgCreateAllowedOracles) Type() string {
	return TypeMsgCreateAllowedOracles
}

func (msg *MsgCreateAllowedOracles) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAllowedOracles) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAllowedOracles) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAllowedOracles{}

func NewMsgUpdateAllowedOracles(creator string, id uint64, name string, address string) *MsgUpdateAllowedOracles {
	return &MsgUpdateAllowedOracles{
		Id:      id,
		Creator: creator,
		Name:    name,
		Address: address,
	}
}

func (msg *MsgUpdateAllowedOracles) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAllowedOracles) Type() string {
	return TypeMsgUpdateAllowedOracles
}

func (msg *MsgUpdateAllowedOracles) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAllowedOracles) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAllowedOracles) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAllowedOracles{}

func NewMsgDeleteAllowedOracles(creator string, id uint64) *MsgDeleteAllowedOracles {
	return &MsgDeleteAllowedOracles{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteAllowedOracles) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAllowedOracles) Type() string {
	return TypeMsgDeleteAllowedOracles
}

func (msg *MsgDeleteAllowedOracles) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAllowedOracles) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAllowedOracles) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
