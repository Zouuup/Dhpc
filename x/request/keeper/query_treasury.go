package keeper

import (
	"context"

	"github.com/DhpcChain/Dhpc/x/request/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Treasury(goCtx context.Context, req *types.QueryGetTreasuryRequest) (*types.QueryGetTreasuryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetTreasury(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTreasuryResponse{Treasury: val}, nil
}
