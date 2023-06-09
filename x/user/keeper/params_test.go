package keeper_test

import (
	"testing"

	testkeeper "github.com/DhpcChain/Dhpc/testutil/keeper"
	"github.com/DhpcChain/Dhpc/x/user/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.UserKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
