package keeper

import (
	"context"

	"Decent/x/request/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AddTreasuryAll(goCtx context.Context, req *types.QueryAllAddTreasuryRequest) (*types.QueryAllAddTreasuryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var addTreasurys []types.AddTreasury
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	addTreasuryStore := prefix.NewStore(store, types.KeyPrefix(types.AddTreasuryKey))

	pageRes, err := query.Paginate(addTreasuryStore, req.Pagination, func(key []byte, value []byte) error {
		var addTreasury types.AddTreasury
		if err := k.cdc.Unmarshal(value, &addTreasury); err != nil {
			return err
		}

		addTreasurys = append(addTreasurys, addTreasury)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAddTreasuryResponse{AddTreasury: addTreasurys, Pagination: pageRes}, nil
}

func (k Keeper) AddTreasury(goCtx context.Context, req *types.QueryGetAddTreasuryRequest) (*types.QueryGetAddTreasuryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	addTreasury, found := k.GetAddTreasury(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAddTreasuryResponse{AddTreasury: addTreasury}, nil
}
