package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/DhpcChain/Dhpc/testutil/keeper"
	"github.com/DhpcChain/Dhpc/testutil/nullify"
	"github.com/DhpcChain/Dhpc/x/request/keeper"
	"github.com/DhpcChain/Dhpc/x/request/types"
)

func createTestTreasury(keeper *keeper.Keeper, ctx sdk.Context) types.Treasury {
	item := types.Treasury{}
	keeper.SetTreasury(ctx, item)
	return item
}

func TestTreasuryGet(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	item := createTestTreasury(keeper, ctx)
	rst, found := keeper.GetTreasury(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTreasuryRemove(t *testing.T) {
	keeper, ctx := keepertest.RequestKeeper(t)
	createTestTreasury(keeper, ctx)
	keeper.RemoveTreasury(ctx)
	_, found := keeper.GetTreasury(ctx)
	require.False(t, found)
}
