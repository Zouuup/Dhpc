package keeper

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"

	"Decent/x/request/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/davecgh/go-spew/spew"
)

const (
	MinimumMiners          = 5
	BlackWait              = 10
	depositPerRequestToken = "1000dhpc"
	treasury               = "decent1cd5e5vr9ftlpzkw45n9d4snpf3vuanl4cf5yyc"
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

	// stage 0:
	// on giving answers
	// if 5 blocks passed since creation record
	// if at least 10 miners provided response
	// switch to stage 1

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

	//TODO: make sure Dataused doesn't have any duplicated fields

	requestRecord.Miners = append(requestRecord.Miners, &minerResponse)
	k.SetRequestRecord(ctx, requestRecord)

	// TODO: value should come from the configuration
	if len(requestRecord.Miners) >= (MinimumMiners) || ctx.BlockHeight() > int64(requestRecord.GetCreatedBlock())+BlackWait {
		requestRecord.Stage = 1
		requestRecord.UpdatedBlock = ctx.BlockHeight()
		k.SetRequestRecord(ctx, requestRecord)
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

	var minerResponse *types.MinerResponse

	for _, miner := range requestRecord.Miners {
		if miner.GetCreator() == msg.Creator {
			minerResponse = miner
			break
		}
	}
	// stage 2.0:
	// start accepting salt and answers
	// if 5 blocks since switch to stage 1.0
	// if at least half + 1 miners provided response

	// stage 2.1:
	// start processing answers and salts, and payments

	minerResponse.Answer = msg.Answer
	minerResponse.Salt = msg.Salt
	k.SetRequestRecord(ctx, requestRecord)

	// if number of miners responses where answer is not zero is more than 2/3 of the total number of miners, then change the stage to 2
	// create a list of miners who have responded with non zero answer
	var nonZeroAnswerMiners []types.MinerResponse
	var rewardedMiners []types.MinerResponse
	for _, miner := range requestRecord.Miners {
		if miner.GetAnswer() != 0 {
			// IMPORTANT TODO: in fact it's crucial that we pay all miners as soon as we find a hash matching most of the answers, if we don't do this, bad players
			// will copy both hash and answers from other miners and will get paid without doing any work
			//
			// check if the hash matches the answer, hash should be sha3 of the answer
			sum := miner.Answer + miner.Salt
			sumStr := strconv.Itoa(int(sum))
			hasher := md5.New()
			hasher.Write([]byte(sumStr))
			answerHash := hex.EncodeToString(hasher.Sum(nil))

			if miner.GetHash() == answerHash {
				nonZeroAnswerMiners = append(nonZeroAnswerMiners, *miner)
			} else {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Hash does not match the answer")
			}
		}
	}
	if len(nonZeroAnswerMiners) > (len(requestRecord.Miners)*2/3) || ctx.BlockHeight() > int64(requestRecord.GetUpdatedBlock())+BlackWait {

		// if more than half of the miners have the same answer, then we can switch to stage 2

		// Find the most common answer
		answerCount := make(map[int32]int)
		for _, miner := range nonZeroAnswerMiners {
			answerCount[miner.Answer]++
		}
		var maxAnswerCount int
		var maxAnswer int32
		for answer, count := range answerCount {
			if count > maxAnswerCount {
				maxAnswerCount = count
				maxAnswer = answer
			}
		}

		for _, miner := range nonZeroAnswerMiners {
			if miner.Answer == maxAnswer {
				rewardedMiners = append(rewardedMiners, miner)
			}
		}

		deposit, err := sdk.ParseCoinsNormalized(depositPerRequestToken)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Unable to parse coins for creating request record")
		}
		// get 40% of the deposit for miners
		minerAmountCoin := sdk.NewCoin("dhpc", deposit[0].Amount.MulRaw(40).QuoRaw(100).QuoRaw(int64(len(rewardedMiners))))
		minerAmount := sdk.NewCoins(minerAmountCoin)

		// iterate through nonZeroAnswerMiners and pay each miner amount
		for _, miner := range rewardedMiners {
			minerAddress, err := sdk.AccAddressFromBech32(miner.Creator)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Unable to parse miner address")
			}
			sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, minerAddress, minerAmount)
			if sdkError != nil {
				return nil, sdkError
			}
		}
		spew.Dump("1******************************************************")
		// iterate through rewardedMiners dataused, find data that's used by 80% of the miners and pay data providers for that data
		dataCounts := make(map[string]int)
		for _, miner := range rewardedMiners {
			// Dataused is MD5 string divided by comma, split it and iterate through it
			dataUsed := strings.Split(miner.DataUsed, ",")
			for _, data := range dataUsed {
				found, _ := k.dataKeeper.GetOwnerByHash(ctx, data)
				if found {
					dataCounts[data]++
				}
			}
		}

		// create a list of data that's used by 80% of the miners, which means count of it equal or more than 80% of the rewardedMiners
		var dataUsedBy80Percent []string
		for data, count := range dataCounts {
			if count >= len(rewardedMiners)*80/100 {
				dataUsedBy80Percent = append(dataUsedBy80Percent, data)
			}
		}

		// get 55% of the deposit for data providers
		dataProviderAmountCoin := sdk.NewCoin("dhpc", deposit[0].Amount.MulRaw(55).QuoRaw(100).QuoRaw(int64(len(dataUsedBy80Percent))))
		dataProviderAmount := sdk.NewCoins(dataProviderAmountCoin)

		// iterate through dataUsedBy80Percent and pay each data provider amount
		for _, data := range dataUsedBy80Percent {
			_, owner := k.dataKeeper.GetOwnerByHash(ctx, data)
			ownerAddress, err := sdk.AccAddressFromBech32(owner)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Unable to parse miner address")
			}

			sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddress, dataProviderAmount)
			if sdkError != nil {
				return nil, sdkError
			}
		}

		treasuryAmount := sdk.NewCoin("dhpc", deposit[0].Amount.Sub(minerAmountCoin.Amount).Sub(dataProviderAmountCoin.Amount))
		treasuryAmountCoin := sdk.NewCoins(treasuryAmount)
		// send all remains to treasury

		treasuryAddress, err := sdk.AccAddressFromBech32(treasury)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Unable to parse treasury address")
		}

		sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, treasuryAddress, treasuryAmountCoin)
		if sdkError != nil {
			return nil, sdkError
		}

		ctx.Logger().Info("Finishing Request", "UUID", requestRecord.UUID, "miners", len(rewardedMiners), "minerAmount", minerAmount, "dataProviders", len(dataUsedBy80Percent), "dataAmount", dataProviderAmount)
		// switch to Stage 2
		requestRecord.Stage = 2
		requestRecord.Score = maxAnswer
		requestRecord.UpdatedBlock = ctx.BlockHeight()
		k.SetRequestRecord(ctx, requestRecord)

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
