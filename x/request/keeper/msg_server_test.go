package keeper_test

import (
	"context"
	"testing"

	keepertest "Dhpc/testutil/keeper"
	"Dhpc/x/request/keeper"
	"Dhpc/x/request/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RequestKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
