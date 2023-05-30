package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "Decent/testutil/keeper"
	"Decent/testutil/nullify"
	"Decent/x/request/types"
)

func TestTreasuryQuery(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestTreasury(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTreasuryRequest
		response *types.QueryGetTreasuryResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTreasuryRequest{},
			response: &types.QueryGetTreasuryResponse{Treasury: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Treasury(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
