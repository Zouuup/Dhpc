package keeper

import (
	"context"
	"fmt"

	"Decent/x/request/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAddTreasury(goCtx context.Context, msg *types.MsgCreateAddTreasury) (*types.MsgCreateAddTreasuryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var addTreasury = types.AddTreasury{
		Creator: msg.Creator,
	}

	if k.GetAddTreasuryCount(ctx) == 0 {
		id := k.AppendAddTreasury(
			ctx,
			addTreasury,
		)
		return &types.MsgCreateAddTreasuryResponse{
			Id: id,
		}, nil
	}

	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Treasury already exists")
}

func (k msgServer) UpdateAddTreasury(goCtx context.Context, msg *types.MsgUpdateAddTreasury) (*types.MsgUpdateAddTreasuryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var addTreasury = types.AddTreasury{
		Creator: msg.Creator,
		Id:      msg.Id,
	}

	// Checks that the element exists
	val, found := k.GetAddTreasury(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAddTreasury(ctx, addTreasury)

	return &types.MsgUpdateAddTreasuryResponse{}, nil
}

func (k msgServer) DeleteAddTreasury(goCtx context.Context, msg *types.MsgDeleteAddTreasury) (*types.MsgDeleteAddTreasuryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetAddTreasury(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAddTreasury(ctx, msg.Id)

	return &types.MsgDeleteAddTreasuryResponse{}, nil
}
