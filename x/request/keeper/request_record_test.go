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

func createNRequestRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RequestRecord {
	items := make([]types.RequestRecord, n)
	for i := range items {
		items[i].UUID = strconv.Itoa(i)

		keeper.SetRequestRecord(ctx, items[i])
	}
	return items
}

func TestRequestRecordGet(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNRequestRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRequestRecord(ctx,
			item.UUID,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRequestRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNRequestRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRequestRecord(ctx,
			item.UUID,
		)
		_, found := keeper.GetRequestRecord(ctx,
			item.UUID,
		)
		require.False(t, found)
	}
}

func TestRequestRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	items := createNRequestRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRequestRecord(ctx)),
	)
}
