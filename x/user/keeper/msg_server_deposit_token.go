package keeper

import (
	"context"

	"Dhpc/x/user/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DepositToken(goCtx context.Context, msg *types.MsgDepositToken) (*types.MsgDepositTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	deposit, err := sdk.ParseCoinsNormalized(msg.Deposit)
	if err != nil {
		panic(err)
	}
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, deposit)
	if sdkError != nil {
		return nil, sdkError
	}

	for _, denom := range deposit {

		account, _ := k.GetOrCreateUser(ctx, msg.Creator)
		updated := false

		for i, userDenom := range account.Deposit {

			// If the user already has a deposit of this denom, add the amount to the existing deposit
			if userDenom.Denom == denom.Denom {
				account.Deposit[i] = account.Deposit[i].AddAmount(denom.Amount)
				updated = true
				k.SetUser(ctx, account)
			}

		}

		// If the user does not have a deposit of this denom, add the denom to the user's deposit
		if !updated {
			account.Deposit = append(account.Deposit, denom)
			k.SetUser(ctx, account)
		}

	}

	return &types.MsgDepositTokenResponse{}, nil
}
