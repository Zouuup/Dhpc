package keeper

import (
	"context"

	"Decent/x/user/types"

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

	k.GetOrCreateUser(ctx, msg.Creator)

	return &types.MsgDepositTokenResponse{}, nil
}
