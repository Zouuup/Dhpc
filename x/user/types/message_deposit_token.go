package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDepositToken = "deposit_token"

var _ sdk.Msg = &MsgDepositToken{}

func NewMsgDepositToken(creator string, deposit string) *MsgDepositToken {
	return &MsgDepositToken{
		Creator: creator,
		Deposit: deposit,
	}
}

func (msg *MsgDepositToken) Route() string {
	return RouterKey
}

func (msg *MsgDepositToken) Type() string {
	return TypeMsgDepositToken
}

func (msg *MsgDepositToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDepositToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDepositToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
