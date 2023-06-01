package keeper

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"

	"Dhpc/x/request/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MinimumMiners defines the minimum number of miners required
const MinimumMiners = 2

// BlackWait defines the wait time in blocks
const BlackWait = 10

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

	if len(requestRecord.Miners) >= MinimumMiners || ctx.BlockHeight() > int64(requestRecord.GetCreatedBlock())+BlackWait {
		requestRecord.Stage = 1
		requestRecord.UpdatedBlock = ctx.BlockHeight()
		k.SetRequestRecord(ctx, requestRecord)
	}

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

	// If the number of miners' responses where the answer is not zero is more than 2/3 of the total number of miners,
	// then change the stage to 2
	// Create a list of miners who have responded with a non-zero answer
	var nonZeroAnswerMiners []types.MinerResponse
	var rewardedMiners []types.MinerResponse
	for _, miner := range requestRecord.Miners {
		if miner.GetAnswer() != 0 {
			sum := miner.Answer + miner.Salt
			sumStr := strconv.Itoa(int(sum))
			hasher := md5.New()
			hasher.Write([]byte(sumStr))
			answerHash := hex.EncodeToString(hasher.Sum(nil))

			if miner.GetHash() == answerHash {
				nonZeroAnswerMiners = append(nonZeroAnswerMiners, *miner)
			} else {
				return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Hash does not match the answer")
			}
		}
	}
	if len(nonZeroAnswerMiners) > (len(requestRecord.Miners)*2/3) || ctx.BlockHeight() > int64(requestRecord.GetUpdatedBlock())+BlackWait {
		// If more than half of the miners have the same answer, then we can switch to stage 2

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
			return nil, errors.Wrap(sdkerrors.ErrInsufficientFunds, "Unable to parse coins for creating request record")
		}

		// Get 40% of the deposit for miners
		minerAmountCoin := sdk.NewCoin("dhpc", deposit[0].Amount.MulRaw(40).QuoRaw(100).QuoRaw(int64(len(rewardedMiners))))
		minerAmount := sdk.NewCoins(minerAmountCoin)

		// Iterate through rewardedMiners and pay each miner amount
		for _, miner := range rewardedMiners {
			minerAddress, err := sdk.AccAddressFromBech32(miner.Creator)
			if err != nil {
				return nil, errors.Wrap(sdkerrors.ErrInvalidAddress, "Unable to parse miner address")
			}
			sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, minerAddress, minerAmount)
			if sdkError != nil {
				return nil, sdkError
			}
		}

		// Iterate through rewardedMiners' dataused, find data that's used by 80% of the miners, and pay data providers for that data
		dataCounts := make(map[string]int)
		for _, miner := range rewardedMiners {
			// Dataused is an MD5 string divided by commas, split it and iterate through it
			dataUsed := strings.Split(miner.DataUsed, ",")
			for _, data := range dataUsed {
				found, _ := k.dataKeeper.GetOwnerByHash(ctx, data)
				if found {
					dataCounts[data]++
				}
			}
		}

		// Create a list of data that's used by 80% of the miners, which means the count is equal to or more than 80% of rewardedMiners
		var dataUsedBy80Percent []string
		for data, count := range dataCounts {
			if count >= len(rewardedMiners)*80/100 {
				dataUsedBy80Percent = append(dataUsedBy80Percent, data)
			}
		}

		// Get 55% of the deposit for data providers
		dataProviderAmountCoin := sdk.NewCoin("dhpc", deposit[0].Amount.MulRaw(55).QuoRaw(100).QuoRaw(int64(len(dataUsedBy80Percent))))
		dataProviderAmount := sdk.NewCoins(dataProviderAmountCoin)

		// Iterate through dataUsedBy80Percent and pay each data provider amount
		for _, data := range dataUsedBy80Percent {
			_, owner := k.dataKeeper.GetOwnerByHash(ctx, data)
			ownerAddress, err := sdk.AccAddressFromBech32(owner)
			if err != nil {
				return nil, errors.Wrap(sdkerrors.ErrInvalidAddress, "Unable to parse miner address")
			}

			sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddress, dataProviderAmount)
			if sdkError != nil {
				return nil, sdkError
			}
		}

		// TODO: Not ideal, instead we should send all remaining tokens to the treasury
		treasuryAmountCoin := sdk.NewCoin("dhpc", deposit[0].Amount.MulRaw(5).QuoRaw(100))
		treasuryAmount := sdk.NewCoins(treasuryAmountCoin)

		treasury, found := k.GetTreasury(ctx)
		if !found {
			return nil, errors.Wrap(sdkerrors.ErrInvalidAddress, "Unable to find treasury address")
		}

		treasuryAddress, err := sdk.AccAddressFromBech32(treasury.Address)
		if err != nil {
			return nil, errors.Wrap(sdkerrors.ErrInvalidAddress, "Unable to parse treasury address")
		}

		sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, treasuryAddress, treasuryAmount)
		if sdkError != nil {
			return nil, sdkError
		}

		ctx.Logger().Info("Finishing Request", "UUID", requestRecord.UUID, "miners", len(rewardedMiners), "minerAmount", minerAmount, "dataProviders", len(dataUsedBy80Percent), "dataAmount", dataProviderAmount, "treasury", treasuryAmount)

		// Switch to Stage 2
		requestRecord.Stage = 2
		requestRecord.Score = maxAnswer
		requestRecord.UpdatedBlock = ctx.BlockHeight()
		k.SetRequestRecord(ctx, requestRecord)
	}

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
