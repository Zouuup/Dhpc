package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateRequestRecord = "create_request_record"
	TypeMsgUpdateRequestRecord = "update_request_record"
	TypeMsgDeleteRequestRecord = "delete_request_record"
)

var _ sdk.Msg = &MsgCreateRequestRecord{}

func NewMsgCreateRequestRecord(
	creator string,
	uUID string,
	tXhash string,
	network string,
	from string,
	address string,
	oracle string,
	block int64,

) *MsgCreateRequestRecord {
	return &MsgCreateRequestRecord{
		Creator: creator,
		UUID:    uUID,
		TXhash:  tXhash,
		Network: network,
		From:    from,
		Address: address,
		Oracle:  oracle,
		Block:   block,
	}
}

func (msg *MsgCreateRequestRecord) Route() string {
	return RouterKey
}

func (msg *MsgCreateRequestRecord) Type() string {
	return TypeMsgCreateRequestRecord
}

func (msg *MsgCreateRequestRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRequestRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRequestRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRequestRecord{}

func NewMsgUpdateRequestRecord(
	creator string,
	uUID string,
	tXhash string,
	network string,
	from string,
	address string,
	score int32,
	oracle string,
	block int64,
	stage int32,
	miners *MinerResponse,

) *MsgUpdateRequestRecord {
	return &MsgUpdateRequestRecord{
		Creator: creator,
		UUID:    uUID,
		TXhash:  tXhash,
		Network: network,
		From:    from,
		Address: address,
		Score:   score,
		Oracle:  oracle,
		Block:   block,
		Stage:   stage,
		Miners:  miners,
	}
}

func (msg *MsgUpdateRequestRecord) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRequestRecord) Type() string {
	return TypeMsgUpdateRequestRecord
}

func (msg *MsgUpdateRequestRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRequestRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRequestRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRequestRecord{}

func NewMsgDeleteRequestRecord(
	creator string,
	uUID string,

) *MsgDeleteRequestRecord {
	return &MsgDeleteRequestRecord{
		Creator: creator,
		UUID:    uUID,
	}
}
func (msg *MsgDeleteRequestRecord) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRequestRecord) Type() string {
	return TypeMsgDeleteRequestRecord
}

func (msg *MsgDeleteRequestRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRequestRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRequestRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
