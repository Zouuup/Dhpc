package keeper_test

import (
	"strconv"
	"testing"

	keepertest "Decent/testutil/keeper"
	"Decent/testutil/nullify"
	"Decent/x/user/keeper"
	"Decent/x/user/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSlashHistory(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SlashHistory {
	items := make([]types.SlashHistory, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetSlashHistory(ctx, items[i])
	}
	return items
}

func TestSlashHistoryGet(t *testing.T) {
	keeper, ctx := keepertest.UserKeeper(t)
	items := createNSlashHistory(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSlashHistory(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSlashHistoryRemove(t *testing.T) {
	keeper, ctx := keepertest.UserKeeper(t)
	items := createNSlashHistory(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSlashHistory(ctx,
			item.Index,
		)
		_, found := keeper.GetSlashHistory(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestSlashHistoryGetAll(t *testing.T) {
	keeper, ctx := keepertest.UserKeeper(t)
	items := createNSlashHistory(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSlashHistory(ctx)),
	)
}
