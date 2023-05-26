package keeper

import (
	"context"

	"Decent/x/request/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMinerResponse(goCtx context.Context, msg *types.MsgCreateMinerResponse) (*types.MsgCreateMinerResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetMinerResponse(
		ctx,
		msg.UUID,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var minerResponse = types.MinerResponse{
		Creator:     msg.Creator,
		UUID:        msg.UUID,
		RequestUUID: msg.RequestUUID,
		Hash:        msg.Hash,
		Answer:      msg.Answer,
	}
	requestRecord, isFound := k.GetRequestRecord(ctx, msg.UUID)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Specified request record does not exist")
	}

	if requestRecord.GetStage() != 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Request is not in zero stage, cannot add miner response")
	}

	requestRecord.Miners = append(requestRecord.Miners, &minerResponse)

	// k.SetMinerResponse(
	// 	ctx,
	// 	minerResponse,
	// )
	return &types.MsgCreateMinerResponseResponse{}, nil
}

func (k msgServer) UpdateMinerResponse(goCtx context.Context, msg *types.MsgUpdateMinerResponse) (*types.MsgUpdateMinerResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMinerResponse(
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

	requestRecord, isFound := k.GetRequestRecord(ctx, msg.UUID)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Specified request record does not exist")
	}

	if requestRecord.GetStage() != 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Request is not in zero stage, cannot add miner response")
	}

	var minerResponse = types.MinerResponse{
		Creator:     msg.Creator,
		UUID:        msg.UUID,
		RequestUUID: msg.RequestUUID,
		Hash:        msg.Hash,
		Answer:      msg.Answer,
	}

	k.SetMinerResponse(ctx, minerResponse)

	return &types.MsgUpdateMinerResponseResponse{}, nil
}

func (k msgServer) DeleteMinerResponse(goCtx context.Context, msg *types.MsgDeleteMinerResponse) (*types.MsgDeleteMinerResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMinerResponse(
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

	k.RemoveMinerResponse(
		ctx,
		msg.UUID,
	)

	return &types.MsgDeleteMinerResponseResponse{}, nil
}
