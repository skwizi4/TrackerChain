package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/Skwizi_4/TrackerChain/testutil/keeper"
	"github.com/Skwizi_4/TrackerChain/x/core/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CoreKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
