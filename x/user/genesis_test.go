package user_test

import (
	"testing"

	keepertest "Decent/testutil/keeper"
	"Decent/testutil/nullify"
	"Decent/x/user"
	"Decent/x/user/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LinkedRequestersList: []types.LinkedRequesters{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		SlashHistoryList: []types.SlashHistory{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.UserKeeper(t)
	user.InitGenesis(ctx, *k, genesisState)
	got := user.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LinkedRequestersList, got.LinkedRequestersList)
	require.ElementsMatch(t, genesisState.SlashHistoryList, got.SlashHistoryList)
	// this line is used by starport scaffolding # genesis/test/assert
}
