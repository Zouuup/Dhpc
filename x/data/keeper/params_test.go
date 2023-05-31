package keeper_test

import (
	"testing"

	testkeeper "Dhpc/testutil/keeper"
	"Dhpc/x/data/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DataKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
