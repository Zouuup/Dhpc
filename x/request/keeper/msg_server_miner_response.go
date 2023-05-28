package keeper

import (
	"context"
	"strconv"

	"Decent/x/request/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/crypto/sha3"
)

const (
	MinimumMiners = 2
	BlackWait     = 5
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
		DataUsed:    msg.DataUsed,
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

	if requestRecord.GetScore() != 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Score must be set to zero at this stage")
	}

	requestRecord.Miners = append(requestRecord.Miners, &minerResponse)
	k.SetRequestRecord(ctx, requestRecord)

	// TODO: value should come from the configuration
	if len(requestRecord.Miners) > (MinimumMiners) {
		if ctx.BlockHeight() > int64(requestRecord.GetBlock())+BlackWait {
			requestRecord.Stage = 1
			k.SetRequestRecord(ctx, requestRecord)
		}
	}

	return &types.MsgCreateMinerResponseResponse{}, nil
}

func (k msgServer) UpdateMinerResponse(goCtx context.Context, msg *types.MsgUpdateMinerResponse) (*types.MsgUpdateMinerResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	requestRecord, isFound := k.GetRequestRecord(
		ctx,
		msg.RequestUUID,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	if requestRecord.GetStage() != 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Request is not in stage one, cannot add miner response")
	}

	var minerResponse = types.MinerResponse{
		Creator:     msg.Creator,
		UUID:        msg.UUID,
		RequestUUID: msg.RequestUUID,
		Hash:        msg.Hash,
		Answer:      msg.Answer,
		DataUsed:    msg.DataUsed,
		Salt:        msg.Salt,
	}

	for i, miner := range requestRecord.Miners {
		if miner.GetCreator() == msg.Creator {
			requestRecord.Miners = append(requestRecord.Miners[:i], requestRecord.Miners[i+1:]...)
			requestRecord.Miners = append(requestRecord.Miners, &minerResponse)
			k.SetRequestRecord(ctx, requestRecord)
			break
		}
	}

	// if number of miners responses where answer is not zero is more than 2/3 of the total number of miners, then change the stage to 2
	// create a list of miners who have responded with non zero answer
	var nonZeroAnswerMiners []types.MinerResponse
	for _, miner := range requestRecord.Miners {
		if miner.GetAnswer() != 0 {
			// IMPORTANT TODO: in fact it's crucial that we pay all miners as soon as we find a hash matching most of the answers, if we don't do this, bad players
			// will copy both hash and answers from other miners and will get paid without doing any work

			// check if the hash matches the answer, hash should be sha3 of the answer
			h := sha3.New256()
			h.Write([]byte(strconv.Itoa(int(miner.Answer))))
			answerHash := string(h.Sum(nil))
			if miner.GetHash() == answerHash {
				nonZeroAnswerMiners = append(nonZeroAnswerMiners, *miner)
			} else {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Hash does not match the answer")
			}
		}
	}
	// if the number of miners who have responded with non zero answer is more than 2/3 of the total number of miners, then change the stage to 2
	if len(nonZeroAnswerMiners) > (len(requestRecord.Miners) * 2 / 3) {
		// if requestrecord block is older than 5 blocks compared to current block, proceed
		// TODO: 5 blocks is arbitrary, should be configurable
		if ctx.BlockHeight() > int64(requestRecord.GetBlock())+5 {
			requestRecord.Stage = 2
			k.SetRequestRecord(ctx, requestRecord)

			// TODO: pay reward to miners who have responded with non zero answer
			// TODO: set the score of the request record to the average of the non zero answers
			// TODO: pay data providers, if a datarecord is specific in all of the miners responses, it should be paid
		}
	}
	return &types.MsgUpdateMinerResponseResponse{}, nil
}

func (k msgServer) DeleteMinerResponse(goCtx context.Context, msg *types.MsgDeleteMinerResponse) (*types.MsgDeleteMinerResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	requestRecord, isFound := k.GetRequestRecord(
		ctx,
		msg.UUID,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != requestRecord.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
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
