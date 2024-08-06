package core_test

import (
	"testing"

	keepertest "github.com/Skwizi_4/TrackerChain/testutil/keeper"
	"github.com/Skwizi_4/TrackerChain/testutil/nullify"
	core "github.com/Skwizi_4/TrackerChain/x/core/module"
	"github.com/Skwizi_4/TrackerChain/x/core/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CoreKeeper(t)
	core.InitGenesis(ctx, k, genesisState)
	got := core.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
