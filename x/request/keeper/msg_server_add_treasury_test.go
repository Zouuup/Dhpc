package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"Decent/x/request/types"
)

func TestAddTreasuryMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateAddTreasury(ctx, &types.MsgCreateAddTreasury{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestAddTreasuryMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateAddTreasury
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateAddTreasury{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAddTreasury{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAddTreasury{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateAddTreasury(ctx, &types.MsgCreateAddTreasury{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateAddTreasury(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAddTreasuryMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteAddTreasury
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteAddTreasury{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteAddTreasury{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteAddTreasury{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateAddTreasury(ctx, &types.MsgCreateAddTreasury{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteAddTreasury(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
