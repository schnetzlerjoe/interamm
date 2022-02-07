package keeper_test

import (
	"testing"

	testkeeper "github.com/schnetzlerjoe/interamm/testutil/keeper"
	"github.com/schnetzlerjoe/interamm/x/interamm/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.InterammKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
