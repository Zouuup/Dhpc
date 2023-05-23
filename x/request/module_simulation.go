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
	opWeightMsgCreateMinerResponse = "op_weight_msg_miner_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMinerResponse int = 100

	opWeightMsgUpdateMinerResponse = "op_weight_msg_miner_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMinerResponse int = 100

	opWeightMsgDeleteMinerResponse = "op_weight_msg_miner_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMinerResponse int = 100

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
		MinerResponseList: []types.MinerResponse{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
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

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
