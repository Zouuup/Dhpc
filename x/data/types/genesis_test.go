package types_test

import (
	"testing"

	"Decent/x/data/types"
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated data",
			genState: &types.GenesisState{
				DataList: []types.Data{
					{
						Address: "0",
						Owner:   "0",
						Network: "0",
					},
					{
						Address: "0",
						Owner:   "0",
						Network: "0",
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
