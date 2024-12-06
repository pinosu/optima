package optima_test

import (
	"testing"

	keepertest "optima/testutil/keeper"
	"optima/testutil/nullify"
	optima "optima/x/optima/module"
	"optima/x/optima/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OptimaKeeper(t)
	optima.InitGenesis(ctx, k, genesisState)
	got := optima.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
