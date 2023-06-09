package keeper

import (
	"encoding/binary"

	"github.com/DhpcChain/Dhpc/x/request/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAllowedOraclesCount get the total number of allowedOracles
func (k Keeper) GetAllowedOraclesCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AllowedOraclesCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAllowedOraclesCount set the total number of allowedOracles
func (k Keeper) SetAllowedOraclesCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AllowedOraclesCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAllowedOracles appends a allowedOracles in the store with a new id and update the count
func (k Keeper) AppendAllowedOracles(
	ctx sdk.Context,
	allowedOracles types.AllowedOracles,
) uint64 {
	// Create the allowedOracles
	count := k.GetAllowedOraclesCount(ctx)

	// Set the ID of the appended value
	allowedOracles.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowedOraclesKey))
	appendedValue := k.cdc.MustMarshal(&allowedOracles)
	store.Set(GetAllowedOraclesIDBytes(allowedOracles.Id), appendedValue)

	// Update allowedOracles count
	k.SetAllowedOraclesCount(ctx, count+1)

	return count
}

// SetAllowedOracles set a specific allowedOracles in the store
func (k Keeper) SetAllowedOracles(ctx sdk.Context, allowedOracles types.AllowedOracles) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowedOraclesKey))
	b := k.cdc.MustMarshal(&allowedOracles)
	store.Set(GetAllowedOraclesIDBytes(allowedOracles.Id), b)
}

// GetAllowedOracles returns a allowedOracles from its id
func (k Keeper) GetAllowedOracles(ctx sdk.Context, id uint64) (val types.AllowedOracles, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowedOraclesKey))
	b := store.Get(GetAllowedOraclesIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAllowedOracles removes a allowedOracles from the store
func (k Keeper) RemoveAllowedOracles(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowedOraclesKey))
	store.Delete(GetAllowedOraclesIDBytes(id))
}

// GetAllAllowedOracles returns all allowedOracles
func (k Keeper) GetAllAllowedOracles(ctx sdk.Context) (list []types.AllowedOracles) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowedOraclesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AllowedOracles
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllowedOraclesIDBytes returns the byte representation of the ID
func GetAllowedOraclesIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAllowedOraclesIDFromBytes returns ID in uint64 format from a byte array
func GetAllowedOraclesIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
