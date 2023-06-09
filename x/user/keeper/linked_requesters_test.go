package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/DhpcChain/Dhpc/testutil/keeper"
	"github.com/DhpcChain/Dhpc/testutil/nullify"
	"github.com/DhpcChain/Dhpc/x/user/keeper"
	"github.com/DhpcChain/Dhpc/x/user/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLinkedRequesters(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LinkedRequesters {
	items := make([]types.LinkedRequesters, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetLinkedRequesters(ctx, items[i])
	}
	return items
}

func TestLinkedRequestersGet(t *testing.T) {
	keeper, ctx := keepertest.UserKeeper(t)
	items := createNLinkedRequesters(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLinkedRequesters(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLinkedRequestersRemove(t *testing.T) {
	keeper, ctx := keepertest.UserKeeper(t)
	items := createNLinkedRequesters(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLinkedRequesters(ctx,
			item.Index,
		)
		_, found := keeper.GetLinkedRequesters(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestLinkedRequestersGetAll(t *testing.T) {
	keeper, ctx := keepertest.UserKeeper(t)
	items := createNLinkedRequesters(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLinkedRequesters(ctx)),
	)
}
