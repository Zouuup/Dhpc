package keeper

import (
	"context"

	"Decent/x/data/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

func (k msgServer) AddData(goCtx context.Context, msg *types.MsgAddData) (*types.MsgAddDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	data := new(types.Data)
	data.Address = msg.Address
	data.Owner = msg.Creator
	data.Network = msg.Network
	data.Method = msg.Method
	data.Score = msg.Score
	data.DateAdded = uint64(ctx.BlockTime().Unix())
	data.DateUpdated = uint64(ctx.BlockTime().Unix())
	data.Uuid = uuid.New().String()

	k.SetData(ctx, *data)

	return &types.MsgAddDataResponse{}, nil
}
