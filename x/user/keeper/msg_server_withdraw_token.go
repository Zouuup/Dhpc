package keeper

import (
	"context"

	"Dhpc/x/user/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WithdrawToken(goCtx context.Context, msg *types.MsgWithdrawToken) (*types.MsgWithdrawTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	withdraw, err := sdk.ParseCoinsNormalized(msg.Withdraw)
	if err != nil {
		panic(err)
	}

	account, _ := k.GetOrCreateUser(ctx, msg.Creator)
	updatedCoin := []sdk.Coin{}
	for _, denom := range withdraw {

		for i, userDenom := range account.Deposit {

			if userDenom.Denom == denom.Denom {
				account.Deposit[i] = account.Deposit[i].SubAmount(denom.Amount)
				coin := sdk.NewCoin(denom.Denom, denom.Amount)
				updatedCoin = append(updatedCoin, coin)
				k.SetUser(ctx, account)
			}
		}

	}

	// TODO: Better error handling when the user tries to withdraw more than they have, or coins they don't have at all
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, updatedCoin)
	if sdkError != nil {
		return nil, sdkError
	}

	return &types.MsgWithdrawTokenResponse{}, nil
}
