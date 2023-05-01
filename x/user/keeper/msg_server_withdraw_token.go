package keeper

import (
	"context"

	"Decent/x/user/types"

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
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, withdraw)
	if sdkError != nil {
		return nil, sdkError
	}

	k.GetOrCreateUser(ctx, msg.Creator)

	return &types.MsgWithdrawTokenResponse{}, nil
}
