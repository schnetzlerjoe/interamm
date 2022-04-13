package keeper

import (
	"context"

	"github.com/schnetzlerjoe/interamm/x/interamm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) SwapCosmos(goCtx context.Context, msg *types.MsgSwapCosmos) (*types.MsgSwapCosmosResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSwapCosmosResponse{}, nil
}
