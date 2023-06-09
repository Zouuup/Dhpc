package keeper

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/DhpcChain/Dhpc/x/data/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddData(goCtx context.Context, msg *types.MsgAddData) (*types.MsgAddDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	data := new(types.Data)
	data.Address = msg.Address
	data.Owner = msg.Creator
	data.Network = msg.Network
	data.Event = msg.Event
	data.Score = msg.Score
	data.DateAdded = uint64(ctx.BlockTime().Unix())
	data.DateUpdated = uint64(ctx.BlockTime().Unix())
	// data.Uuid is md5sum of address + network + owner + method
	hash := md5.Sum([]byte(data.Address + data.Network + data.Owner + data.Event))
	hashString := hex.EncodeToString(hash[:])
	data.Hash = hashString

	k.SetData(ctx, *data)

	return &types.MsgAddDataResponse{}, nil
}
