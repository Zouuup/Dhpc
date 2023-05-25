package request_test

import (
	"testing"

	keepertest "Decent/testutil/keeper"
	"Decent/testutil/nullify"
	"Decent/x/request"
	"Decent/x/request/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MinerResponseList: []types.MinerResponse{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		RequestRecordList: []types.RequestRecord{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		AllowedOraclesList: []types.AllowedOracles{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AllowedOraclesCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RequestKeeper(t)
	request.InitGenesis(ctx, *k, genesisState)
	got := request.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MinerResponseList, got.MinerResponseList)
	require.ElementsMatch(t, genesisState.RequestRecordList, got.RequestRecordList)
	require.ElementsMatch(t, genesisState.AllowedOraclesList, got.AllowedOraclesList)
	require.Equal(t, genesisState.AllowedOraclesCount, got.AllowedOraclesCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
