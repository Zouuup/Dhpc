package keeper

import (
	"context"

	"Decent/x/user/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DepositToken(goCtx context.Context, msg *types.MsgDepositToken) (*types.MsgDepositTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDepositTokenResponse{}, nil
}
