package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDepositToken{}, "user/DepositToken", nil)
	cdc.RegisterConcrete(&MsgWithdrawToken{}, "user/WithdrawToken", nil)
	cdc.RegisterConcrete(&MsgAddLinkedRequester{}, "user/AddLinkedRequester", nil)
	cdc.RegisterConcrete(&MsgRemoveLinkedRequester{}, "user/RemoveLinkedRequester", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDepositToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddLinkedRequester{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveLinkedRequester{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
