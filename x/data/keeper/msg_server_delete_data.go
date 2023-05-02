package keeper

import (
	"context"

	"Decent/x/data/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteData(goCtx context.Context, msg *types.MsgDeleteData) (*types.MsgDeleteDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.RemoveData(ctx, msg.Address, msg.Creator, msg.Network)

	return &types.MsgDeleteDataResponse{}, nil
}
