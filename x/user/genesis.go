package user

import (
	"github.com/DhpcChain/Dhpc/x/user/keeper"
	"github.com/DhpcChain/Dhpc/x/user/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the linkedRequesters
	for _, elem := range genState.LinkedRequestersList {
		k.SetLinkedRequesters(ctx, elem)
	}
	// Set all the slashHistory
	for _, elem := range genState.SlashHistoryList {
		k.SetSlashHistory(ctx, elem)
	}
	// Set all the user
	for _, elem := range genState.UserList {
		k.SetUser(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.LinkedRequestersList = k.GetAllLinkedRequesters(ctx)
	genesis.SlashHistoryList = k.GetAllSlashHistory(ctx)
	genesis.UserList = k.GetAllUser(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
