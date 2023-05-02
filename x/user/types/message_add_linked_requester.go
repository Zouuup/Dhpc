package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddLinkedRequester = "add_linked_requester"

var _ sdk.Msg = &MsgAddLinkedRequester{}

func NewMsgAddLinkedRequester(creator string, network string, address string) *MsgAddLinkedRequester {
	return &MsgAddLinkedRequester{
		Creator: creator,
		Network: network,
		Address: address,
	}
}

func (msg *MsgAddLinkedRequester) Route() string {
	return RouterKey
}

func (msg *MsgAddLinkedRequester) Type() string {
	return TypeMsgAddLinkedRequester
}

func (msg *MsgAddLinkedRequester) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddLinkedRequester) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddLinkedRequester) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
