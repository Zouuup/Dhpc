package request

import (
	"Decent/x/request/keeper"
	"Decent/x/request/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the allowedOracles
	for _, elem := range genState.AllowedOraclesList {
		k.SetAllowedOracles(ctx, elem)
	}

	// Set allowedOracles count
	k.SetAllowedOraclesCount(ctx, genState.AllowedOraclesCount)
	// Set all the minerResponse
	for _, elem := range genState.MinerResponseList {
		k.SetMinerResponse(ctx, elem)
	}
	// Set all the requestRecord
	for _, elem := range genState.RequestRecordList {
		k.SetRequestRecord(ctx, elem)
	}
	// Set all the addTreasury
	for _, elem := range genState.AddTreasuryList {
		k.SetAddTreasury(ctx, elem)
	}

	// Set addTreasury count
	k.SetAddTreasuryCount(ctx, genState.AddTreasuryCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AllowedOraclesList = k.GetAllAllowedOracles(ctx)
	genesis.AllowedOraclesCount = k.GetAllowedOraclesCount(ctx)
	genesis.MinerResponseList = k.GetAllMinerResponse(ctx)
	genesis.RequestRecordList = k.GetAllRequestRecord(ctx)
	genesis.AddTreasuryList = k.GetAllAddTreasury(ctx)
	genesis.AddTreasuryCount = k.GetAddTreasuryCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
