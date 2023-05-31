package keeper

import (
	"context"

	"Dhpc/x/user/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) RemoveLinkedRequester(goCtx context.Context, msg *types.MsgRemoveLinkedRequester) (*types.MsgRemoveLinkedRequesterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	account, _ := k.GetOrCreateUser(ctx, msg.Creator)
	updated := false

	newLinkedRequest := types.LinkedRequesters{
		Network: msg.Network,
		Address: msg.Address,
	}

	for i, linkedRequest := range account.LinkedRequester {
		if linkedRequest.Network == newLinkedRequest.Network && linkedRequest.Address == newLinkedRequest.Address {
			account.LinkedRequester = append(account.LinkedRequester[:i], account.LinkedRequester[i+1:]...)
			updated = true
		}
	}

	if !updated {
		return nil, status.Error(codes.InvalidArgument, "Address does not exist")
	}

	k.SetUser(ctx, account)

	return &types.MsgRemoveLinkedRequesterResponse{}, nil
}
