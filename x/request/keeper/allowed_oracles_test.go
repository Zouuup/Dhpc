package keeper_test

import (
	"testing"

	keepertest "github.com/DhpcChain/Dhpc/testutil/keeper"
	"github.com/DhpcChain/Dhpc/testutil/nullify"
	"github.com/DhpcChain/Dhpc/x/request/keeper"
	"github.com/DhpcChain/Dhpc/x/request/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNAllowedOracles(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AllowedOracles {
	items := make([]types.AllowedOracles, n)
	for i := range items {
		items[i].Id = keeper.AppendAllowedOracles(ctx, items[i])
	}
	return items
}

func TestAllowedOraclesGet(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNAllowedOracles(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAllowedOracles(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAllowedOraclesRemove(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNAllowedOracles(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAllowedOracles(ctx, item.Id)
		_, found := keeper.GetAllowedOracles(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAllowedOraclesGetAll(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNAllowedOracles(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAllowedOracles(ctx)),
	)
}

func TestAllowedOraclesCount(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNAllowedOracles(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAllowedOraclesCount(ctx))
}
