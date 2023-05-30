package keeper

import (
	"context"

	"Decent/x/request/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateRequestRecord(goCtx context.Context, msg *types.MsgCreateRequestRecord) (*types.MsgCreateRequestRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetRequestRecord(
		ctx,
		msg.UUID,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// TODO: if From and Creator are not the same, check if address is a in the list of allowed oracles
	deposit, err := sdk.ParseCoinsNormalized(depositPerRequestToken)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Unable to parse coins for creating request record")
	}

	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, deposit)
	if sdkError != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Unable to deposit coins for creating request record")
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

	k.SetRequestRecord(
		ctx,
		requestRecord,
	)
	return &types.MsgCreateRequestRecordResponse{}, nil
}

func (k msgServer) UpdateRequestRecord(goCtx context.Context, msg *types.MsgUpdateRequestRecord) (*types.MsgUpdateRequestRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRequestRecord(
		ctx,
		msg.UUID,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Can't update records after creation")
}

func (k msgServer) DeleteRequestRecord(goCtx context.Context, msg *types.MsgDeleteRequestRecord) (*types.MsgDeleteRequestRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRequestRecord(
		ctx,
		msg.UUID,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveRequestRecord(
		ctx,
		msg.UUID,
	)

	return &types.MsgDeleteRequestRecordResponse{}, nil
}
