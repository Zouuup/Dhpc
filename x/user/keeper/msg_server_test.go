package keeper_test

import (
	"context"
	"testing"

	keepertest "Dhpc/testutil/keeper"
	"Dhpc/x/user/keeper"
	"Dhpc/x/user/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.UserKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
