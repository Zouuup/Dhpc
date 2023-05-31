package keeper

import (
	"context"

	"Dhpc/x/user/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SlashHistoryAll(goCtx context.Context, req *types.QueryAllSlashHistoryRequest) (*types.QueryAllSlashHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var slashHistorys []types.SlashHistory
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	slashHistoryStore := prefix.NewStore(store, types.KeyPrefix(types.SlashHistoryKeyPrefix))

	pageRes, err := query.Paginate(slashHistoryStore, req.Pagination, func(key []byte, value []byte) error {
		var slashHistory types.SlashHistory
		if err := k.cdc.Unmarshal(value, &slashHistory); err != nil {
			return err
		}

		slashHistorys = append(slashHistorys, slashHistory)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSlashHistoryResponse{SlashHistory: slashHistorys, Pagination: pageRes}, nil
}

func (k Keeper) SlashHistory(goCtx context.Context, req *types.QueryGetSlashHistoryRequest) (*types.QueryGetSlashHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetSlashHistory(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSlashHistoryResponse{SlashHistory: val}, nil
}
