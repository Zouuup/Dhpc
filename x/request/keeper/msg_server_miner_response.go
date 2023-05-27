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
	requestRecord, isFound := k.GetRequestRecord(ctx, msg.RequestUUID)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Specified request record does not exist")
	}

	if requestRecord.GetStage() != 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Request is not in stage zero, cannot add miner response")
	}

	for _, miner := range requestRecord.Miners {
		if miner.GetCreator() == msg.Creator {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Miner already responded to this request")
		}
	}

	requestRecord.Miners = append(requestRecord.Miners, &minerResponse)
	k.SetRequestRecord(ctx, requestRecord)

	// if number of numbers of miners is more than 2/3 of the total number of miners, then change the stage to 1
	if len(requestRecord.Miners) > (1) {
		requestRecord.Stage = 1
		k.SetRequestRecord(ctx, requestRecord)
	}

	return &types.MsgCreateMinerResponseResponse{}, nil
}

func (k msgServer) UpdateMinerResponse(goCtx context.Context, msg *types.MsgUpdateMinerResponse) (*types.MsgUpdateMinerResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valFound, isFound := k.GetRequestRecord(
		ctx,
		msg.RequestUUID,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	if valFound.GetStage() != 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Request is not in stage one, cannot add miner response")
	}

	var minerResponse = types.MinerResponse{
		Creator:     msg.Creator,
		UUID:        msg.UUID,
		RequestUUID: msg.RequestUUID,
		Hash:        msg.Hash,
		Answer:      msg.Answer,
	}

	for i, miner := range valFound.Miners {
		if miner.GetCreator() == msg.Creator {
			valFound.Miners = append(valFound.Miners[:i], valFound.Miners[i+1:]...)
			valFound.Miners = append(valFound.Miners, &minerResponse)
			k.SetRequestRecord(ctx, valFound)
			break
		}
	}

	return &types.MsgUpdateMinerResponseResponse{}, nil
}

func (k msgServer) DeleteMinerResponse(goCtx context.Context, msg *types.MsgDeleteMinerResponse) (*types.MsgDeleteMinerResponseResponse, error) {
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

	for i, miner := range valFound.Miners {
		if miner.GetCreator() == msg.Creator {
			valFound.Miners = append(valFound.Miners[:i], valFound.Miners[i+1:]...)
			k.SetRequestRecord(ctx, valFound)
			break
		}
	}

	return &types.MsgDeleteMinerResponseResponse{}, nil
}
