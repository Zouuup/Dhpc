package keeper

import (
	"Dhpc/x/user/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetUser set a specific user in the store from its index
func (k Keeper) SetUser(ctx sdk.Context, user types.User) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKeyPrefix))
	b := k.cdc.MustMarshal(&user)
	store.Set(types.UserKey(
		user.Account,
	), b)
}

// GetUser returns a user from its index
func (k Keeper) GetUser(
	ctx sdk.Context,
	account string,

) (val types.User, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKeyPrefix))

	b := store.Get(types.UserKey(
		account,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetOrCreateUser(
	ctx sdk.Context,
	account string,

) (val types.User, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKeyPrefix))

	b := store.Get(types.UserKey(
		account,
	))
	if b == nil {
		user := new(types.User)
		user.Account = account
		user.Deposit = []sdk.Coin{}
		user.Reputation = 0
		user.LinkedRequester = []*types.LinkedRequesters{}
		user.SlashHistory = []*types.SlashHistory{}
		k.SetUser(ctx, *user)
		b = store.Get(types.UserKey(
			account,
		))

	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUser removes a user from the store
func (k Keeper) RemoveUser(
	ctx sdk.Context,
	account string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKeyPrefix))
	store.Delete(types.UserKey(
		account,
	))
}

// GetAllUser returns all user
func (k Keeper) GetAllUser(ctx sdk.Context) (list []types.User) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.User
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
