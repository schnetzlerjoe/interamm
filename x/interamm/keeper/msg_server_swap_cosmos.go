package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/schnetzlerjoe/interamm/x/interamm/types"
)

func (k msgServer) SwapCosmos(goCtx context.Context, msg *types.MsgSwapCosmos) (*types.MsgSwapCosmosResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSwapCosmosResponse{}, nil
}
