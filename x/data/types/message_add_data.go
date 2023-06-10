package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddData = "add_data"

var _ sdk.Msg = &MsgAddData{}

func NewMsgAddData(creator string, address string, owner string, network string, event string, blockValidity uint64, score int32) *MsgAddData {
	return &MsgAddData{
		Creator:       creator,
		Address:       address,
		Owner:         owner,
		Network:       network,
		Event:         event,
		Blockvalidity: blockValidity,
		Score:         score,
	}
}

func (msg *MsgAddData) Route() string {
	return RouterKey
}

func (msg *MsgAddData) Type() string {
	return TypeMsgAddData
}

func (msg *MsgAddData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
