package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "Dhpc/testutil/keeper"
	"Dhpc/x/request/keeper"
	"Dhpc/x/request/types"
)

func TestTreasuryMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RequestKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateTreasury{Creator: creator}
	_, err := srv.CreateTreasury(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetTreasury(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestTreasuryMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateTreasury
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateTreasury{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTreasury{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RequestKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateTreasury{Creator: creator}
			_, err := srv.CreateTreasury(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateTreasury(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetTreasury(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestTreasuryMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteTreasury
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteTreasury{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteTreasury{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RequestKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateTreasury(wctx, &types.MsgCreateTreasury{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteTreasury(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetTreasury(ctx)
				require.False(t, found)
			}
		})
	}
}
