package trackerchain_test

import (
	"testing"

	keepertest "github.com/Skwizi_4/TrackerChain/testutil/keeper"
	"github.com/Skwizi_4/TrackerChain/testutil/nullify"
	trackerchain "github.com/Skwizi_4/TrackerChain/x/trackerchain/module"
	"github.com/Skwizi_4/TrackerChain/x/trackerchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TrackerchainKeeper(t)
	trackerchain.InitGenesis(ctx, k, genesisState)
	got := trackerchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
