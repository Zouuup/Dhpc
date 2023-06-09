package keeper

import (
	"context"

	"github.com/DhpcChain/Dhpc/x/request/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// CreateRequestRecord creates a request record
func (k msgServer) CreateRequestRecord(goCtx context.Context, msg *types.MsgCreateRequestRecord) (*types.MsgCreateRequestRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetRequestRecord(ctx, msg.UUID)
	if isFound {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// If From and Creator are not the same, check if the address is in the list of allowed oracles
	if msg.From != msg.Creator {
		allowedOracles := k.GetAllAllowedOracles(ctx)
		allowed := false
		for _, oracle := range allowedOracles {
			if oracle.Address == msg.Creator {
				allowed = true
			}
		}
		if !allowed {
			return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "not allowed as an oracle, submit the request with your own address")
		}
	}

	deposit, err := sdk.ParseCoinsNormalized(depositPerRequestToken)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInsufficientFunds, "Unable to parse coins for creating the request record")
	}

	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, deposit)
	if sdkError != nil {
		return nil, errors.Wrap(sdkerrors.ErrInsufficientFunds, "Unable to deposit coins for creating the request record")
	}

	var requestRecord = types.RequestRecord{
		Creator:      msg.Creator,
		UUID:         msg.UUID,
		TXhash:       msg.TXhash,
		Network:      msg.Network,
		From:         msg.From,
		Address:      msg.Address,
		Score:        msg.Score,
		Oracle:       msg.Oracle,
		Block:        msg.Block,
		Stage:        0,
		CreatedBlock: ctx.BlockHeight(),
	}

	ctx.Logger().Info("Creating Request record", "UUID", requestRecord.UUID)

	k.SetRequestRecord(ctx, requestRecord)

	return &types.MsgCreateRequestRecordResponse{}, nil
}

// UpdateRequestRecord updates a request record
func (k msgServer) UpdateRequestRecord(goCtx context.Context, msg *types.MsgUpdateRequestRecord) (*types.MsgUpdateRequestRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRequestRecord(ctx, msg.UUID)
	if !isFound {
		return nil, errors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Check if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	return nil, errors.Wrap(sdkerrors.ErrKeyNotFound, "Cannot update records after creation")
}

// DeleteRequestRecord deletes a request record
func (k msgServer) DeleteRequestRecord(goCtx context.Context, msg *types.MsgDeleteRequestRecord) (*types.MsgDeleteRequestRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRequestRecord(ctx, msg.UUID)
	if !isFound {
		return nil, errors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Check if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveRequestRecord(ctx, msg.UUID)

	return &types.MsgDeleteRequestRecordResponse{}, nil
}
