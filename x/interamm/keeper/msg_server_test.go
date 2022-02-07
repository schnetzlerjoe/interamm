package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/schnetzlerjoe/interamm/testutil/keeper"
	"github.com/schnetzlerjoe/interamm/x/interamm/keeper"
	"github.com/schnetzlerjoe/interamm/x/interamm/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.InterammKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
