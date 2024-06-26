package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "fttm/testutil/keeper"
	"fttm/x/fttm/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.FttmKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
