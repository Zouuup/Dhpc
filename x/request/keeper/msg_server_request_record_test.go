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

func TestRequestRecordMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RequestKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateRequestRecord{Creator: creator,
			UUID: strconv.Itoa(i),
		}
		_, err := srv.CreateRequestRecord(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetRequestRecord(ctx,
			expected.UUID,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestRequestRecordMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateRequestRecord
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateRequestRecord{Creator: creator,
				UUID: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateRequestRecord{Creator: "B",
				UUID: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateRequestRecord{Creator: creator,
				UUID: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RequestKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateRequestRecord{Creator: creator,
				UUID: strconv.Itoa(0),
			}
			_, err := srv.CreateRequestRecord(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateRequestRecord(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetRequestRecord(ctx,
					expected.UUID,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestRequestRecordMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteRequestRecord
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteRequestRecord{Creator: creator,
				UUID: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteRequestRecord{Creator: "B",
				UUID: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteRequestRecord{Creator: creator,
				UUID: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RequestKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateRequestRecord(wctx, &types.MsgCreateRequestRecord{Creator: creator,
				UUID: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteRequestRecord(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetRequestRecord(ctx,
					tc.request.UUID,
				)
				require.False(t, found)
			}
		})
	}
}
