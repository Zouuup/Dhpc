package keeper

import (
	"context"

	"Decent/x/user/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) AddLinkedRequester(goCtx context.Context, msg *types.MsgAddLinkedRequester) (*types.MsgAddLinkedRequesterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	account, _ := k.GetOrCreateUser(ctx, msg.Creator)

	newLinkedRequest := types.LinkedRequesters{
		Network: msg.Network,
		Address: msg.Address,
	}

	for _, linkedRequest := range account.LinkedRequester {
		if linkedRequest.Network == newLinkedRequest.Network && linkedRequest.Address == newLinkedRequest.Address {
			return nil, status.Error(codes.InvalidArgument, "Address is already linked")
		}
	}
	account.LinkedRequester = append(account.LinkedRequester, &newLinkedRequest)
	k.SetUser(ctx, account)
	return &types.MsgAddLinkedRequesterResponse{}, nil
}
