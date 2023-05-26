package request

import (
	"math/rand"

	"Decent/testutil/sample"
	requestsimulation "Decent/x/request/simulation"
	"Decent/x/request/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = requestsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateAllowedOracles = "op_weight_msg_allowed_oracles"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAllowedOracles int = 100

	opWeightMsgUpdateAllowedOracles = "op_weight_msg_allowed_oracles"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAllowedOracles int = 100

	opWeightMsgDeleteAllowedOracles = "op_weight_msg_allowed_oracles"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAllowedOracles int = 100

	opWeightMsgCreateMinerResponse = "op_weight_msg_miner_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMinerResponse int = 100

	opWeightMsgUpdateMinerResponse = "op_weight_msg_miner_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMinerResponse int = 100

	opWeightMsgDeleteMinerResponse = "op_weight_msg_miner_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMinerResponse int = 100

	opWeightMsgCreateRequestRecord = "op_weight_msg_request_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRequestRecord int = 100

	opWeightMsgUpdateRequestRecord = "op_weight_msg_request_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRequestRecord int = 100

	opWeightMsgDeleteRequestRecord = "op_weight_msg_request_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRequestRecord int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	requestGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AllowedOraclesList: []types.AllowedOracles{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		AllowedOraclesCount: 2,
		MinerResponseList: []types.MinerResponse{
			{
				Creator: sample.AccAddress(),
				UUID:    "0",
			},
			{
				Creator: sample.AccAddress(),
				UUID:    "1",
			},
		},
		RequestRecordList: []types.RequestRecord{
			{
				Creator: sample.AccAddress(),
				UUID:    "0",
			},
			{
				Creator: sample.AccAddress(),
				UUID:    "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&requestGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAllowedOracles int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAllowedOracles, &weightMsgCreateAllowedOracles, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAllowedOracles = defaultWeightMsgCreateAllowedOracles
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAllowedOracles,
		requestsimulation.SimulateMsgCreateAllowedOracles(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAllowedOracles int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAllowedOracles, &weightMsgUpdateAllowedOracles, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAllowedOracles = defaultWeightMsgUpdateAllowedOracles
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAllowedOracles,
		requestsimulation.SimulateMsgUpdateAllowedOracles(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAllowedOracles int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAllowedOracles, &weightMsgDeleteAllowedOracles, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAllowedOracles = defaultWeightMsgDeleteAllowedOracles
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAllowedOracles,
		requestsimulation.SimulateMsgDeleteAllowedOracles(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateMinerResponse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMinerResponse, &weightMsgCreateMinerResponse, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMinerResponse = defaultWeightMsgCreateMinerResponse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMinerResponse,
		requestsimulation.SimulateMsgCreateMinerResponse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMinerResponse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMinerResponse, &weightMsgUpdateMinerResponse, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMinerResponse = defaultWeightMsgUpdateMinerResponse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMinerResponse,
		requestsimulation.SimulateMsgUpdateMinerResponse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMinerResponse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMinerResponse, &weightMsgDeleteMinerResponse, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMinerResponse = defaultWeightMsgDeleteMinerResponse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMinerResponse,
		requestsimulation.SimulateMsgDeleteMinerResponse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateRequestRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateRequestRecord, &weightMsgCreateRequestRecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRequestRecord = defaultWeightMsgCreateRequestRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRequestRecord,
		requestsimulation.SimulateMsgCreateRequestRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRequestRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateRequestRecord, &weightMsgUpdateRequestRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRequestRecord = defaultWeightMsgUpdateRequestRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRequestRecord,
		requestsimulation.SimulateMsgUpdateRequestRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRequestRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteRequestRecord, &weightMsgDeleteRequestRecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRequestRecord = defaultWeightMsgDeleteRequestRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRequestRecord,
		requestsimulation.SimulateMsgDeleteRequestRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
