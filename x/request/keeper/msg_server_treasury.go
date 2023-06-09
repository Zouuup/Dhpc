package keeper

import (
	"context"

	"github.com/DhpcChain/Dhpc/x/request/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateTreasury(goCtx context.Context, msg *types.MsgCreateTreasury) (*types.MsgCreateTreasuryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetTreasury(ctx)
	if isFound {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var treasury = types.Treasury{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	k.SetTreasury(
		ctx,
		treasury,
	)
	return &types.MsgCreateTreasuryResponse{}, nil
}

func (k msgServer) UpdateTreasury(goCtx context.Context, msg *types.MsgUpdateTreasury) (*types.MsgUpdateTreasuryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTreasury(ctx)
	if !isFound {
		return nil, errors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var treasury = types.Treasury{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	k.SetTreasury(ctx, treasury)

	return &types.MsgUpdateTreasuryResponse{}, nil
}

func (k msgServer) DeleteTreasury(goCtx context.Context, msg *types.MsgDeleteTreasury) (*types.MsgDeleteTreasuryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTreasury(ctx)
	if !isFound {
		return nil, errors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTreasury(ctx)

	return &types.MsgDeleteTreasuryResponse{}, nil
}
