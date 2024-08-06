package keeper

import (
	"context"
	"github.com/Skwizi_4/TrackerChain/x/core/types"
)

func (k msgServer) ValidateBlock(context.Context, *types.MsgValidateBlock) (*types.MsgValidateBlockResponse, error) {
	return &types.MsgValidateBlockResponse{}, nil
}
