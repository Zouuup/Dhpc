package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteData = "delete_data"

var _ sdk.Msg = &MsgDeleteData{}

func NewMsgDeleteData(creator string, address string, owner string, network string) *MsgDeleteData {
	return &MsgDeleteData{
		Creator: creator,
		Address: address,
		Owner:   owner,
		Network: network,
	}
}

func (msg *MsgDeleteData) Route() string {
	return RouterKey
}

func (msg *MsgDeleteData) Type() string {
	return TypeMsgDeleteData
}

func (msg *MsgDeleteData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
