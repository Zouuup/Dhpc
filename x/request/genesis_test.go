package request_test

import (
	"testing"

	keepertest "github.com/DhpcChain/Dhpc/testutil/keeper"
	"github.com/DhpcChain/Dhpc/testutil/nullify"
	"github.com/DhpcChain/Dhpc/x/request"
	"github.com/DhpcChain/Dhpc/x/request/types"
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
		Treasury: &types.Treasury{
			Address: "18",
		},
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
	require.Equal(t, genesisState.Treasury, got.Treasury)
	// this line is used by starport scaffolding # genesis/test/assert
}
