package keeper

import (
	"context"

	"Decent/x/request/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateRequestRecord(goCtx context.Context, msg *types.MsgCreateRequestRecord) (*types.MsgCreateRequestRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetRequestRecord(
		ctx,
		msg.UUID,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var requestRecord = types.RequestRecord{
		Creator: msg.Creator,
		UUID:    msg.UUID,
		TXhash:  msg.TXhash,
		Network: msg.Network,
		From:    msg.From,
		Address: msg.Address,
		Score:   msg.Score,
		Oracle:  msg.Oracle,
		Block:   msg.Block,
		Stage:   msg.Stage,
	}

	k.SetRequestRecord(
		ctx,
		requestRecord,
	)
	return &types.MsgCreateRequestRecordResponse{}, nil
}

func (k msgServer) UpdateRequestRecord(goCtx context.Context, msg *types.MsgUpdateRequestRecord) (*types.MsgUpdateRequestRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRequestRecord(
		ctx,
		msg.UUID,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var requestRecord = types.RequestRecord{
		Creator: msg.Creator,
		UUID:    msg.UUID,
		TXhash:  msg.TXhash,
		Network: msg.Network,
		From:    msg.From,
		Address: msg.Address,
		Score:   msg.Score,
		Oracle:  msg.Oracle,
		Block:   msg.Block,
		Stage:   msg.Stage,
	}

	k.SetRequestRecord(ctx, requestRecord)

	return &types.MsgUpdateRequestRecordResponse{}, nil
}

func (k msgServer) DeleteRequestRecord(goCtx context.Context, msg *types.MsgDeleteRequestRecord) (*types.MsgDeleteRequestRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRequestRecord(
		ctx,
		msg.UUID,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveRequestRecord(
		ctx,
		msg.UUID,
	)

	return &types.MsgDeleteRequestRecordResponse{}, nil
}
