package keeper

import (
	"context"
	"github.com/Skwizi_4/TrackerChain/x/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateBlock(goctx context.Context, Block *types.MsgCreateBlock) (*types.MsgCreateBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	lastBlockId := k.GetLastIndex(ctx)
	var result = types.UserResults{
		Creator:   Block.Creator,
		UserA:     Block.Body.UserA,
		UserB:     Block.Body.UserB,
		Winner:    Block.Body.Winner,
		Timestamp: Block.Body.Timestamp,
	}
	Id, err := k.SetResult(ctx, lastBlockId, result)
	if err != nil {
		return nil, err
	}
	return &types.MsgCreateBlockResponse{
		Id: Id,
	}, nil
}
