package keeper_test

import (
	"context"
	"testing"

	keepertest "Dhpc/testutil/keeper"
	"Dhpc/x/data/keeper"
	"Dhpc/x/data/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DataKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
