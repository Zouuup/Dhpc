package keeper

import (
	"Dhpc/x/user/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLinkedRequesters set a specific linkedRequesters in the store from its index
func (k Keeper) SetLinkedRequesters(ctx sdk.Context, linkedRequesters types.LinkedRequesters) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LinkedRequestersKeyPrefix))
	b := k.cdc.MustMarshal(&linkedRequesters)
	store.Set(types.LinkedRequestersKey(
		linkedRequesters.Index,
	), b)
}

// GetLinkedRequesters returns a linkedRequesters from its index
func (k Keeper) GetLinkedRequesters(
	ctx sdk.Context,
	index string,

) (val types.LinkedRequesters, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LinkedRequestersKeyPrefix))

	b := store.Get(types.LinkedRequestersKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLinkedRequesters removes a linkedRequesters from the store
func (k Keeper) RemoveLinkedRequesters(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LinkedRequestersKeyPrefix))
	store.Delete(types.LinkedRequestersKey(
		index,
	))
}

// GetAllLinkedRequesters returns all linkedRequesters
func (k Keeper) GetAllLinkedRequesters(ctx sdk.Context) (list []types.LinkedRequesters) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LinkedRequestersKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LinkedRequesters
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
