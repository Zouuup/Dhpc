package types_test

import (
	"testing"

	"Decent/x/request/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated minerResponse",
			genState: &types.GenesisState{
				MinerResponseList: []types.MinerResponse{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated requestRecord",
			genState: &types.GenesisState{
				RequestRecordList: []types.RequestRecord{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated allowedOracles",
			genState: &types.GenesisState{
				AllowedOraclesList: []types.AllowedOracles{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid allowedOracles count",
			genState: &types.GenesisState{
				AllowedOraclesList: []types.AllowedOracles{
					{
						Id: 1,
					},
				},
				AllowedOraclesCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
