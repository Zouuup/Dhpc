package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/DhpcChain/Dhpc/testutil/network"
	"github.com/DhpcChain/Dhpc/testutil/nullify"
	"github.com/DhpcChain/Dhpc/x/request/client/cli"
	"github.com/DhpcChain/Dhpc/x/request/types"
)

func networkWithTreasuryObjects(t *testing.T) (*network.Network, types.Treasury) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	treasury := &types.Treasury{}
	nullify.Fill(&treasury)
	state.Treasury = treasury
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.Treasury
}

func TestShowTreasury(t *testing.T) {
	net, obj := networkWithTreasuryObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.Treasury
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowTreasury(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetTreasuryResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Treasury)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Treasury),
				)
			}
		})
	}
}
