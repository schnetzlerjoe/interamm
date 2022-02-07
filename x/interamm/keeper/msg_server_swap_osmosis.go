package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/schnetzlerjoe/interamm/x/interamm/types"
)

func (k msgServer) SwapOsmosis(goCtx context.Context, msg *types.MsgSwapOsmosis) (*types.MsgSwapOsmosisResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSwapOsmosisResponse{}, nil
}
