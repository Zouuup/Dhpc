package keeper

import (
	"context"
	"strings"

	"github.com/DhpcChain/Dhpc/x/request/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// depositPerRequestToken is the deposit amount per request token
const depositPerRequestToken = "1000dhpc"

// CreateMinerResponse creates a miner response
func (k msgServer) CreateMinerResponse(goCtx context.Context, msg *types.MsgCreateMinerResponse) (*types.MsgCreateMinerResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetMinerResponse(ctx, msg.UUID)
	if isFound {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// Iterate through dataused and make sure they are all unique
	// TODO: protect against zero length dataused
	if strings.Contains(msg.DataUsed, ",") {
		dataUsedList := strings.Split(msg.DataUsed, ",")
		for i := 0; i < len(dataUsedList); i++ {
			for j := i + 1; j < len(dataUsedList); j++ {
				if dataUsedList[i] == dataUsedList[j] {
					return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Provided Data used should be unique")
				}
			}
		}
	}

	var minerResponse = types.MinerResponse{
		Creator:     msg.Creator,
		UUID:        msg.UUID,
		RequestUUID: msg.RequestUUID,
		Hash:        msg.Hash,
		Answer:      msg.Answer,
		DataUsed:    msg.DataUsed,
	}
	requestRecord, isFound := k.GetRequestRecord(ctx, msg.RequestUUID)
	if !isFound {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Specified request record does not exist")
	}

	if requestRecord.GetStage() != 0 {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Request is not in stage zero, cannot add miner response")
	}

	for _, miner := range requestRecord.Miners {
		if miner.GetCreator() == msg.Creator {
			return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Miner already responded to this request")
		}
	}

	if requestRecord.GetScore() != 0 {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Score must be set to zero at this stage")
	}

	requestRecord.Miners = append(requestRecord.Miners, &minerResponse)
	k.SetRequestRecord(ctx, requestRecord)

	return &types.MsgCreateMinerResponseResponse{}, nil
}

// UpdateMinerResponse updates a miner response
func (k msgServer) UpdateMinerResponse(goCtx context.Context, msg *types.MsgUpdateMinerResponse) (*types.MsgUpdateMinerResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	requestRecord, isFound := k.GetRequestRecord(ctx, msg.RequestUUID)
	if !isFound {
		return nil, errors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	if requestRecord.GetStage() != 1 {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Request is not in stage one, cannot add miner response")
	}

	var minerResponse *types.MinerResponse

	for _, miner := range requestRecord.Miners {
		if miner.GetCreator() == msg.Creator {
			minerResponse = miner
			break
		}
	}

	// Stage 2.0:
	// Start accepting salt and answers
	// If 5 blocks have passed since switch to stage 1.0
	// If at least half + 1 miners provided a response

	// Stage 2.1:
	// Start processing answers and salts, and payments
	minerResponse.Answer = msg.Answer
	minerResponse.Salt = msg.Salt
	k.SetRequestRecord(ctx, requestRecord)

	return &types.MsgUpdateMinerResponseResponse{}, nil
}

// DeleteMinerResponse deletes a miner response
func (k msgServer) DeleteMinerResponse(goCtx context.Context, msg *types.MsgDeleteMinerResponse) (*types.MsgDeleteMinerResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	requestRecord, isFound := k.GetRequestRecord(ctx, msg.UUID)
	if !isFound {
		return nil, errors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != requestRecord.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	for i, miner := range requestRecord.Miners {
		if miner.GetCreator() == msg.Creator {
			requestRecord.Miners = append(requestRecord.Miners[:i], requestRecord.Miners[i+1:]...)
			k.SetRequestRecord(ctx, requestRecord)
			break
		}
	}

	return &types.MsgDeleteMinerResponseResponse{}, nil
}
