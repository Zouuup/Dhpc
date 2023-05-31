package keeper

import (
	"context"

	"Dhpc/x/request/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllowedOraclesAll(goCtx context.Context, req *types.QueryAllAllowedOraclesRequest) (*types.QueryAllAllowedOraclesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var allowedOracless []types.AllowedOracles
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	allowedOraclesStore := prefix.NewStore(store, types.KeyPrefix(types.AllowedOraclesKey))

	pageRes, err := query.Paginate(allowedOraclesStore, req.Pagination, func(key []byte, value []byte) error {
		var allowedOracles types.AllowedOracles
		if err := k.cdc.Unmarshal(value, &allowedOracles); err != nil {
			return err
		}

		allowedOracless = append(allowedOracless, allowedOracles)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAllowedOraclesResponse{AllowedOracles: allowedOracless, Pagination: pageRes}, nil
}

func (k Keeper) AllowedOracles(goCtx context.Context, req *types.QueryGetAllowedOraclesRequest) (*types.QueryGetAllowedOraclesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	allowedOracles, found := k.GetAllowedOracles(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAllowedOraclesResponse{AllowedOracles: allowedOracles}, nil
}
