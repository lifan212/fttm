package fttm_test

import (
	"testing"

	keepertest "fttm/testutil/keeper"
	"fttm/testutil/nullify"
	fttm "fttm/x/fttm/module"
	"fttm/x/fttm/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FttmKeeper(t)
	fttm.InitGenesis(ctx, k, genesisState)
	got := fttm.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
