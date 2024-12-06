package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "optima/testutil/keeper"
	"optima/x/optima/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.OptimaKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
