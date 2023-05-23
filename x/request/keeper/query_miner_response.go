package keeper

import (
	"context"

	"Decent/x/request/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MinerResponseAll(goCtx context.Context, req *types.QueryAllMinerResponseRequest) (*types.QueryAllMinerResponseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var minerResponses []types.MinerResponse
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	minerResponseStore := prefix.NewStore(store, types.KeyPrefix(types.MinerResponseKeyPrefix))

	pageRes, err := query.Paginate(minerResponseStore, req.Pagination, func(key []byte, value []byte) error {
		var minerResponse types.MinerResponse
		if err := k.cdc.Unmarshal(value, &minerResponse); err != nil {
			return err
		}

		minerResponses = append(minerResponses, minerResponse)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMinerResponseResponse{MinerResponse: minerResponses, Pagination: pageRes}, nil
}

func (k Keeper) MinerResponse(goCtx context.Context, req *types.QueryGetMinerResponseRequest) (*types.QueryGetMinerResponseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetMinerResponse(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMinerResponseResponse{MinerResponse: val}, nil
}
