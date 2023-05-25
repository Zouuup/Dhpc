package keeper

import (
	"context"
	"fmt"

	"Decent/x/request/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAllowedOracles(goCtx context.Context, msg *types.MsgCreateAllowedOracles) (*types.MsgCreateAllowedOraclesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var allowedOracles = types.AllowedOracles{
		Creator: msg.Creator,
		Name:    msg.Name,
		Address: msg.Address,
	}

	if k.GetAllowedOraclesCount(ctx) == 0 {
		id := k.AppendAllowedOracles(
			ctx,
			allowedOracles,
		)
		return &types.MsgCreateAllowedOraclesResponse{
			Id: id,
		}, nil
	}

	currentOracles := k.GetAllAllowedOracles(ctx)
	for _, oracle := range currentOracles {
		if oracle.Address == msg.Creator {
			id := k.AppendAllowedOracles(
				ctx,
				allowedOracles,
			)
			return &types.MsgCreateAllowedOraclesResponse{
				Id: id,
			}, nil
		}
	}

	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Creator %s is not trusted", msg.Creator))
}

func (k msgServer) UpdateAllowedOracles(goCtx context.Context, msg *types.MsgUpdateAllowedOracles) (*types.MsgUpdateAllowedOraclesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var allowedOracles = types.AllowedOracles{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Address: msg.Address,
	}

	// Checks that the element exists
	val, found := k.GetAllowedOracles(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAllowedOracles(ctx, allowedOracles)

	return &types.MsgUpdateAllowedOraclesResponse{}, nil
}

func (k msgServer) DeleteAllowedOracles(goCtx context.Context, msg *types.MsgDeleteAllowedOracles) (*types.MsgDeleteAllowedOraclesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetAllowedOracles(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAllowedOracles(ctx, msg.Id)

	return &types.MsgDeleteAllowedOraclesResponse{}, nil
}
