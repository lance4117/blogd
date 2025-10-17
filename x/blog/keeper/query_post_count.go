package keeper

import (
	"context"
	"github.com/lance4117/blogd/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PostCount(goCtx context.Context, req *types.QueryPostCountRequest) (*types.QueryPostCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	count := k.GetBlogCount(ctx)

	return &types.QueryPostCountResponse{Count: int64(count)}, nil
}
