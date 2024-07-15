package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/Skwizi_4/TrackerChain/testutil/keeper"
	"github.com/Skwizi_4/TrackerChain/x/trackerchain/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.TrackerchainKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
