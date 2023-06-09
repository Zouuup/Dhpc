package types_test

import (
	"testing"

	"github.com/DhpcChain/Dhpc/x/user/types"
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
				UserList: []types.User{
					{
						Account: "0",
					},
					{
						Account: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated linkedRequesters",
			genState: &types.GenesisState{
				LinkedRequestersList: []types.LinkedRequesters{
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
			desc: "duplicated slashHistory",
			genState: &types.GenesisState{
				SlashHistoryList: []types.SlashHistory{
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
			desc: "duplicated user",
			genState: &types.GenesisState{
				UserList: []types.User{
					{
						Account: "0",
					},
					{
						Account: "0",
					},
				},
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
