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

func (k Keeper) RequestRecordAll(goCtx context.Context, req *types.QueryAllRequestRecordRequest) (*types.QueryAllRequestRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var requestRecords []types.RequestRecord
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	requestRecordStore := prefix.NewStore(store, types.KeyPrefix(types.RequestRecordKeyPrefix))

	pageRes, err := query.Paginate(requestRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var requestRecord types.RequestRecord
		if err := k.cdc.Unmarshal(value, &requestRecord); err != nil {
			return err
		}

		requestRecords = append(requestRecords, requestRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRequestRecordResponse{RequestRecord: requestRecords, Pagination: pageRes}, nil
}

func (k Keeper) RequestRecord(goCtx context.Context, req *types.QueryGetRequestRecordRequest) (*types.QueryGetRequestRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetRequestRecord(
		ctx,
		req.UUID,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRequestRecordResponse{RequestRecord: val}, nil
}
