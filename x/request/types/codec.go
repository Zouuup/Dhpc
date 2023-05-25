package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateMinerResponse{}, "request/CreateMinerResponse", nil)
	cdc.RegisterConcrete(&MsgUpdateMinerResponse{}, "request/UpdateMinerResponse", nil)
	cdc.RegisterConcrete(&MsgDeleteMinerResponse{}, "request/DeleteMinerResponse", nil)
	cdc.RegisterConcrete(&MsgCreateRequestRecord{}, "request/CreateRequestRecord", nil)
	cdc.RegisterConcrete(&MsgUpdateRequestRecord{}, "request/UpdateRequestRecord", nil)
	cdc.RegisterConcrete(&MsgDeleteRequestRecord{}, "request/DeleteRequestRecord", nil)
	cdc.RegisterConcrete(&MsgCreateAllowedOracles{}, "request/CreateAllowedOracles", nil)
	cdc.RegisterConcrete(&MsgUpdateAllowedOracles{}, "request/UpdateAllowedOracles", nil)
	cdc.RegisterConcrete(&MsgDeleteAllowedOracles{}, "request/DeleteAllowedOracles", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMinerResponse{},
		&MsgUpdateMinerResponse{},
		&MsgDeleteMinerResponse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRequestRecord{},
		&MsgUpdateRequestRecord{},
		&MsgDeleteRequestRecord{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAllowedOracles{},
		&MsgUpdateAllowedOracles{},
		&MsgDeleteAllowedOracles{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
