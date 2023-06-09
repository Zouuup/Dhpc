package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/DhpcChain/Dhpc/testutil/keeper"
	"github.com/DhpcChain/Dhpc/testutil/nullify"
	"github.com/DhpcChain/Dhpc/x/request/keeper"
	"github.com/DhpcChain/Dhpc/x/request/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMinerResponse(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MinerResponse {
	items := make([]types.MinerResponse, n)
	for i := range items {
		items[i].UUID = strconv.Itoa(i)

		keeper.SetMinerResponse(ctx, items[i])
	}
	return items
}

func TestMinerResponseGet(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNMinerResponse(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMinerResponse(ctx,
			item.UUID,
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
			item.UUID,
		)
		_, found := keeper.GetMinerResponse(ctx,
			item.UUID,
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
