package keeper

import (
	"Decent/x/request/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRequestRecord set a specific requestRecord in the store from its index
func (k Keeper) SetRequestRecord(ctx sdk.Context, requestRecord types.RequestRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestRecordKeyPrefix))
	b := k.cdc.MustMarshal(&requestRecord)
	store.Set(types.RequestRecordKey(
		requestRecord.UUID,
	), b)
}

// GetRequestRecord returns a requestRecord from its index
func (k Keeper) GetRequestRecord(
	ctx sdk.Context,
	uUID string,

) (val types.RequestRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestRecordKeyPrefix))

	b := store.Get(types.RequestRecordKey(
		uUID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRequestRecord removes a requestRecord from the store
func (k Keeper) RemoveRequestRecord(
	ctx sdk.Context,
	uUID string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestRecordKeyPrefix))
	store.Delete(types.RequestRecordKey(
		uUID,
	))
}

// GetAllRequestRecord returns all requestRecord
func (k Keeper) GetAllRequestRecord(ctx sdk.Context) (list []types.RequestRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RequestRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
