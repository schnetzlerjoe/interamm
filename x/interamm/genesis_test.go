package interamm_test

import (
	"testing"

	keepertest "github.com/schnetzlerjoe/interamm/testutil/keeper"
	"github.com/schnetzlerjoe/interamm/testutil/nullify"
	"github.com/schnetzlerjoe/interamm/x/interamm"
	"github.com/schnetzlerjoe/interamm/x/interamm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.InterammKeeper(t)
	interamm.InitGenesis(ctx, *k, genesisState)
	got := interamm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
