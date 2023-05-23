package keeper_test

import (
	"strconv"
	"testing"

	keepertest "Decent/testutil/keeper"
	"Decent/testutil/nullify"
	"Decent/x/request/keeper"
	"Decent/x/request/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMinerResponse(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MinerResponse {
	items := make([]types.MinerResponse, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetMinerResponse(ctx, items[i])
	}
	return items
}

func TestMinerResponseGet(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNMinerResponse(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMinerResponse(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMinerResponseRemove(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNMinerResponse(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMinerResponse(ctx,
			item.Index,
		)
		_, found := keeper.GetMinerResponse(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestMinerResponseGetAll(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNMinerResponse(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMinerResponse(ctx)),
	)
}
