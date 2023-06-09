package keeper

import (
	"github.com/DhpcChain/Dhpc/x/user/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetSlashHistory set a specific slashHistory in the store from its index
func (k Keeper) SetSlashHistory(ctx sdk.Context, slashHistory types.SlashHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlashHistoryKeyPrefix))
	b := k.cdc.MustMarshal(&slashHistory)
	store.Set(types.SlashHistoryKey(
		slashHistory.Index,
	), b)
}

// GetSlashHistory returns a slashHistory from its index
func (k Keeper) GetSlashHistory(
	ctx sdk.Context,
	index string,

) (val types.SlashHistory, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlashHistoryKeyPrefix))

	b := store.Get(types.SlashHistoryKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSlashHistory removes a slashHistory from the store
func (k Keeper) RemoveSlashHistory(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlashHistoryKeyPrefix))
	store.Delete(types.SlashHistoryKey(
		index,
	))
}

// GetAllSlashHistory returns all slashHistory
func (k Keeper) GetAllSlashHistory(ctx sdk.Context) (list []types.SlashHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlashHistoryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SlashHistory
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
