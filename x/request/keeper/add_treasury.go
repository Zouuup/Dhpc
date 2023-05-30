package keeper

import (
	"encoding/binary"

	"Decent/x/request/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAddTreasuryCount get the total number of addTreasury
func (k Keeper) GetAddTreasuryCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AddTreasuryCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAddTreasuryCount set the total number of addTreasury
func (k Keeper) SetAddTreasuryCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AddTreasuryCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAddTreasury appends a addTreasury in the store with a new id and update the count
func (k Keeper) AppendAddTreasury(
	ctx sdk.Context,
	addTreasury types.AddTreasury,
) uint64 {
	// Create the addTreasury
	count := k.GetAddTreasuryCount(ctx)

	// Set the ID of the appended value
	addTreasury.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddTreasuryKey))
	appendedValue := k.cdc.MustMarshal(&addTreasury)
	store.Set(GetAddTreasuryIDBytes(addTreasury.Id), appendedValue)

	// Update addTreasury count
	k.SetAddTreasuryCount(ctx, count+1)

	return count
}

// SetAddTreasury set a specific addTreasury in the store
func (k Keeper) SetAddTreasury(ctx sdk.Context, addTreasury types.AddTreasury) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddTreasuryKey))
	b := k.cdc.MustMarshal(&addTreasury)
	store.Set(GetAddTreasuryIDBytes(addTreasury.Id), b)
}

// GetAddTreasury returns a addTreasury from its id
func (k Keeper) GetAddTreasury(ctx sdk.Context, id uint64) (val types.AddTreasury, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddTreasuryKey))
	b := store.Get(GetAddTreasuryIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAddTreasury removes a addTreasury from the store
func (k Keeper) RemoveAddTreasury(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddTreasuryKey))
	store.Delete(GetAddTreasuryIDBytes(id))
}

// GetAllAddTreasury returns all addTreasury
func (k Keeper) GetAllAddTreasury(ctx sdk.Context) (list []types.AddTreasury) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddTreasuryKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AddTreasury
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAddTreasuryIDBytes returns the byte representation of the ID
func GetAddTreasuryIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAddTreasuryIDFromBytes returns ID in uint64 format from a byte array
func GetAddTreasuryIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
