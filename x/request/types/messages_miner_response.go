package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMinerResponse = "create_miner_response"
	TypeMsgUpdateMinerResponse = "update_miner_response"
	TypeMsgDeleteMinerResponse = "delete_miner_response"
)

var _ sdk.Msg = &MsgCreateMinerResponse{}

func NewMsgCreateMinerResponse(
	creator string,
	uUID string,
	requestUUID string,
	hash string,
	dataUsed string,

) *MsgCreateMinerResponse {
	return &MsgCreateMinerResponse{
		Creator:     creator,
		UUID:        uUID,
		RequestUUID: requestUUID,
		Hash:        hash,
		DataUsed:    dataUsed,
	}
}

func (msg *MsgCreateMinerResponse) Route() string {
	return RouterKey
}

func (msg *MsgCreateMinerResponse) Type() string {
	return TypeMsgCreateMinerResponse
}

func (msg *MsgCreateMinerResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMinerResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMinerResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMinerResponse{}

func NewMsgUpdateMinerResponse(
	creator string,
	uUID string,
	requestUUID string,
	answer int32,
	salt int32,

) *MsgUpdateMinerResponse {
	return &MsgUpdateMinerResponse{
		Creator:     creator,
		UUID:        uUID,
		RequestUUID: requestUUID,
		Answer:      answer,
		Salt:        salt,
	}
}

func (msg *MsgUpdateMinerResponse) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMinerResponse) Type() string {
	return TypeMsgUpdateMinerResponse
}

func (msg *MsgUpdateMinerResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMinerResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMinerResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMinerResponse{}

func NewMsgDeleteMinerResponse(
	creator string,
	uUID string,

) *MsgDeleteMinerResponse {
	return &MsgDeleteMinerResponse{
		Creator: creator,
		UUID:    uUID,
	}
}
func (msg *MsgDeleteMinerResponse) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMinerResponse) Type() string {
	return TypeMsgDeleteMinerResponse
}

func (msg *MsgDeleteMinerResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMinerResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMinerResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
