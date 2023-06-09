package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/DhpcChain/Dhpc/testutil/keeper"
	"github.com/DhpcChain/Dhpc/testutil/nullify"
	"github.com/DhpcChain/Dhpc/x/data/keeper"
	"github.com/DhpcChain/Dhpc/x/data/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Data {
	items := make([]types.Data, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].Owner = strconv.Itoa(i)
		items[i].Network = strconv.Itoa(i)

		keeper.SetData(ctx, items[i])
	}
	return items
}

func TestDataGet(t *testing.T) {
	keeper, ctx := keepertest.DataKeeper(t)
	items := createNData(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetData(ctx,
			item.Address,
			item.Owner,
			item.Network,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDataRemove(t *testing.T) {
	keeper, ctx := keepertest.DataKeeper(t)
	items := createNData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveData(ctx,
			item.Address,
			item.Owner,
			item.Network,
		)
		_, found := keeper.GetData(ctx,
			item.Address,
			item.Owner,
			item.Network,
		)
		require.False(t, found)
	}
}

func TestDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.DataKeeper(t)
	items := createNData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllData(ctx)),
	)
}
