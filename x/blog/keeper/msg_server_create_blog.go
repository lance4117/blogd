package keeper

import (
	"blog/x/blog/types"
	"context"

	"gitee.com/lance4117/GoFuse/times"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateBlog(goCtx context.Context, msg *types.MsgCreateBlog) (*types.MsgCreateBlogResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	blogItem := types.BlogItem{
		Creator: msg.Creator,
		Author:  msg.Creator,
		Ctm:     times.NowMilli(),
		Title:   msg.Title,
		Content: msg.Content,
	}

	id := k.AppendBlog(ctx, blogItem)

	return &types.MsgCreateBlogResponse{
		Id: id,
	}, nil
}
