package keeper_test

import (
	"testing"

	keepertest "Decent/testutil/keeper"
	"Decent/testutil/nullify"
	"Decent/x/request/keeper"
	"Decent/x/request/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNAddTreasury(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AddTreasury {
	items := make([]types.AddTreasury, n)
	for i := range items {
		items[i].Id = keeper.AppendAddTreasury(ctx, items[i])
	}
	return items
}

func TestAddTreasuryGet(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNAddTreasury(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAddTreasury(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAddTreasuryRemove(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNAddTreasury(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAddTreasury(ctx, item.Id)
		_, found := keeper.GetAddTreasury(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAddTreasuryGetAll(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNAddTreasury(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAddTreasury(ctx)),
	)
}

func TestAddTreasuryCount(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNAddTreasury(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAddTreasuryCount(ctx))
}
