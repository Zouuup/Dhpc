package keeper

import (
	"context"

	"github.com/DhpcChain/Dhpc/x/user/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LinkedRequestersAll(goCtx context.Context, req *types.QueryAllLinkedRequestersRequest) (*types.QueryAllLinkedRequestersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var linkedRequesterss []types.LinkedRequesters
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	linkedRequestersStore := prefix.NewStore(store, types.KeyPrefix(types.LinkedRequestersKeyPrefix))

	pageRes, err := query.Paginate(linkedRequestersStore, req.Pagination, func(key []byte, value []byte) error {
		var linkedRequesters types.LinkedRequesters
		if err := k.cdc.Unmarshal(value, &linkedRequesters); err != nil {
			return err
		}

		linkedRequesterss = append(linkedRequesterss, linkedRequesters)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLinkedRequestersResponse{LinkedRequesters: linkedRequesterss, Pagination: pageRes}, nil
}

func (k Keeper) LinkedRequesters(goCtx context.Context, req *types.QueryGetLinkedRequestersRequest) (*types.QueryGetLinkedRequestersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetLinkedRequesters(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLinkedRequestersResponse{LinkedRequesters: val}, nil
}
