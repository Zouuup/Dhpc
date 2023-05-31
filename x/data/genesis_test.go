package data_test

import (
	"testing"

	keepertest "Dhpc/testutil/keeper"
	"Dhpc/testutil/nullify"
	"Dhpc/x/data"
	"Dhpc/x/data/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DataList: []types.Data{
			{
				Address: "0",
				Owner:   "0",
				Network: "0",
			},
			{
				Address: "1",
				Owner:   "1",
				Network: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DataKeeper(t)
	data.InitGenesis(ctx, *k, genesisState)
	got := data.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DataList, got.DataList)
	// this line is used by starport scaffolding # genesis/test/assert
}
