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

		AllowedOraclesList: []types.AllowedOracles{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AllowedOraclesCount: 2,
		MinerResponseList: []types.MinerResponse{
			{
				UUID: "0",
			},
			{
				UUID: "1",
			},
		},
		RequestRecordList: []types.RequestRecord{
			{
				UUID: "0",
			},
			{
				UUID: "1",
			},
		},
		AddTreasuryList: []types.AddTreasury{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AddTreasuryCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RequestKeeper(t)
	request.InitGenesis(ctx, *k, genesisState)
	got := request.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AllowedOraclesList, got.AllowedOraclesList)
	require.Equal(t, genesisState.AllowedOraclesCount, got.AllowedOraclesCount)
	require.ElementsMatch(t, genesisState.MinerResponseList, got.MinerResponseList)
	require.ElementsMatch(t, genesisState.RequestRecordList, got.RequestRecordList)
	require.ElementsMatch(t, genesisState.AddTreasuryList, got.AddTreasuryList)
	require.Equal(t, genesisState.AddTreasuryCount, got.AddTreasuryCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
