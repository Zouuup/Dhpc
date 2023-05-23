package keeper

import (
	"Decent/x/request/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetMinerResponse set a specific minerResponse in the store from its index
func (k Keeper) SetMinerResponse(ctx sdk.Context, minerResponse types.MinerResponse) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerResponseKeyPrefix))
	b := k.cdc.MustMarshal(&minerResponse)
	store.Set(types.MinerResponseKey(
		minerResponse.Index,
	), b)
}

// GetMinerResponse returns a minerResponse from its index
func (k Keeper) GetMinerResponse(
	ctx sdk.Context,
	index string,

) (val types.MinerResponse, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerResponseKeyPrefix))

	b := store.Get(types.MinerResponseKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMinerResponse removes a minerResponse from the store
func (k Keeper) RemoveMinerResponse(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerResponseKeyPrefix))
	store.Delete(types.MinerResponseKey(
		index,
	))
}

// GetAllMinerResponse returns all minerResponse
func (k Keeper) GetAllMinerResponse(ctx sdk.Context) (list []types.MinerResponse) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerResponseKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MinerResponse
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
