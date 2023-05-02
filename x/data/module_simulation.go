package data

import (
	"math/rand"

	"Decent/testutil/sample"
	datasimulation "Decent/x/data/simulation"
	"Decent/x/data/types"
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
	_ = datasimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgAddData = "op_weight_msg_add_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddData int = 100

	opWeightMsgDeleteData = "op_weight_msg_delete_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteData int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dataGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dataGenesis)
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

	var weightMsgAddData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddData, &weightMsgAddData, nil,
		func(_ *rand.Rand) {
			weightMsgAddData = defaultWeightMsgAddData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddData,
		datasimulation.SimulateMsgAddData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteData, &weightMsgDeleteData, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteData = defaultWeightMsgDeleteData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteData,
		datasimulation.SimulateMsgDeleteData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
