package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveLinkedRequester = "remove_linked_requester"

var _ sdk.Msg = &MsgRemoveLinkedRequester{}

func NewMsgRemoveLinkedRequester(creator string, network string, address string) *MsgRemoveLinkedRequester {
	return &MsgRemoveLinkedRequester{
		Creator: creator,
		Network: network,
		Address: address,
	}
}

func (msg *MsgRemoveLinkedRequester) Route() string {
	return RouterKey
}

func (msg *MsgRemoveLinkedRequester) Type() string {
	return TypeMsgRemoveLinkedRequester
}

func (msg *MsgRemoveLinkedRequester) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveLinkedRequester) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveLinkedRequester) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
