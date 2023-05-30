package keeper

import (
	"Decent/x/request/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTreasury set treasury in the store
func (k Keeper) SetTreasury(ctx sdk.Context, treasury types.Treasury) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TreasuryKey))
	b := k.cdc.MustMarshal(&treasury)
	store.Set([]byte{0}, b)
}

// GetTreasury returns treasury
func (k Keeper) GetTreasury(ctx sdk.Context) (val types.Treasury, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TreasuryKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTreasury removes treasury from the store
func (k Keeper) RemoveTreasury(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TreasuryKey))
	store.Delete([]byte{0})
}
