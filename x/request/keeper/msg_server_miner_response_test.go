package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "Decent/testutil/keeper"
	"Decent/x/request/keeper"
	"Decent/x/request/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMinerResponseMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RequestKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateMinerResponse{Creator: creator,
			UUID: strconv.Itoa(i),
		}
		_, err := srv.CreateMinerResponse(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetMinerResponse(ctx,
			expected.UUID,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestMinerResponseMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateMinerResponse
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateMinerResponse{Creator: creator,
				UUID: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateMinerResponse{Creator: "B",
				UUID: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateMinerResponse{Creator: creator,
				UUID: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RequestKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateMinerResponse{Creator: creator,
				UUID: strconv.Itoa(0),
			}
			_, err := srv.CreateMinerResponse(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateMinerResponse(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetMinerResponse(ctx,
					expected.UUID,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestMinerResponseMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteMinerResponse
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteMinerResponse{Creator: creator,
				UUID: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteMinerResponse{Creator: "B",
				UUID: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteMinerResponse{Creator: creator,
				UUID: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RequestKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateMinerResponse(wctx, &types.MsgCreateMinerResponse{Creator: creator,
				UUID: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteMinerResponse(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetMinerResponse(ctx,
					tc.request.UUID,
				)
				require.False(t, found)
			}
		})
	}
}
