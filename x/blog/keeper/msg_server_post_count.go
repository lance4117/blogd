package keeper

import (
	"context"

	"github.com/lance4117/blogd/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PostCount(goCtx context.Context, msg *types.MsgPostCount) (*types.MsgPostCountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgPostCountResponse{}, nil
}
