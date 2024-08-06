package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/Skwizi_4/TrackerChain/testutil/keeper"
	"github.com/Skwizi_4/TrackerChain/x/core/keeper"
	"github.com/Skwizi_4/TrackerChain/x/core/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.CoreKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
