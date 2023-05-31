package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"Dhpc/x/request/types"
)

func TestAllowedOraclesMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateAllowedOracles(ctx, &types.MsgCreateAllowedOracles{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestAllowedOraclesMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateAllowedOracles
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateAllowedOracles{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAllowedOracles{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAllowedOracles{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateAllowedOracles(ctx, &types.MsgCreateAllowedOracles{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateAllowedOracles(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAllowedOraclesMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteAllowedOracles
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteAllowedOracles{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteAllowedOracles{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteAllowedOracles{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateAllowedOracles(ctx, &types.MsgCreateAllowedOracles{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteAllowedOracles(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
