package user

import (
	"math/rand"

	"github.com/DhpcChain/Dhpc/testutil/sample"
	usersimulation "github.com/DhpcChain/Dhpc/x/user/simulation"
	"github.com/DhpcChain/Dhpc/x/user/types"
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
	_ = usersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgDepositToken = "op_weight_msg_deposit_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDepositToken int = 100

	opWeightMsgWithdrawToken = "op_weight_msg_withdraw_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawToken int = 100

	opWeightMsgAddLinkedRequester = "op_weight_msg_add_linked_requester"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddLinkedRequester int = 100

	opWeightMsgRemoveLinkedRequester = "op_weight_msg_remove_linked_requester"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveLinkedRequester int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	userGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&userGenesis)
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

	var weightMsgDepositToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDepositToken, &weightMsgDepositToken, nil,
		func(_ *rand.Rand) {
			weightMsgDepositToken = defaultWeightMsgDepositToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDepositToken,
		usersimulation.SimulateMsgDepositToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawToken, &weightMsgWithdrawToken, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawToken = defaultWeightMsgWithdrawToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawToken,
		usersimulation.SimulateMsgWithdrawToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddLinkedRequester int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddLinkedRequester, &weightMsgAddLinkedRequester, nil,
		func(_ *rand.Rand) {
			weightMsgAddLinkedRequester = defaultWeightMsgAddLinkedRequester
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddLinkedRequester,
		usersimulation.SimulateMsgAddLinkedRequester(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveLinkedRequester int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveLinkedRequester, &weightMsgRemoveLinkedRequester, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveLinkedRequester = defaultWeightMsgRemoveLinkedRequester
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveLinkedRequester,
		usersimulation.SimulateMsgRemoveLinkedRequester(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
