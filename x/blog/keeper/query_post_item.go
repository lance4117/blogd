package keeper

import (
	"context"

	"github.com/lance4117/blogd/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PostItem(goCtx context.Context, req *types.QueryPostItemRequest) (*types.QueryPostItemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	val, _ := k.GetBlog(ctx, req.Id)

	return &types.QueryPostItemResponse{Blog: &val}, nil
}
