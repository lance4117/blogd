package keeper

import (
	"context"

	"github.com/lance4117/blogd/x/blog/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PostList(goCtx context.Context, req *types.QueryPostListRequest) (*types.QueryPostListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) 读取总量，用于边界修正与返回 total
	total := k.GetBlogCount(ctx)
	if total == 0 {
		return &types.QueryPostListResponse{Blogs: nil, Total: 0}, nil
	}

	// 2) 解析并修正区间 [startID, endID]，均为“包含式”语义
	startID := req.StartId // 默认 0
	endID := req.EndId     // 0 或未填视为直到最后一条
	if endID == 0 || endID >= total-1 {
		endID = total - 1
	}
	// 若传反了，自动交换
	if startID > endID {
		startID, endID = endID, startID
	}
	// 进一步夹紧到合法范围
	if startID >= total {
		// 起点越界，直接返回空
		return &types.QueryPostListResponse{Blogs: nil, Total: total}, nil
	}

	// 3) 前缀化 store，避免扫全库
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	pstore := prefix.NewStore(storeAdapter, []byte(types.BlogKey))

	// 4) 计算迭代器边界：
	//    Cosmos KV Iterator 是 [start, end) 半开区间。
	//    -> startKey  = key(startID)
	//    -> endKey    = key(endID+1)；若 endID 已是最后一条，则 endKey = nil 表示到前缀结尾
	startKey := GetBlogIDBytes(startID)

	var endKey []byte
	if endID >= total-1 {
		endKey = nil // 直到该前缀末尾
	} else {
		endKey = GetBlogIDBytes(endID + 1)
	}

	it := pstore.Iterator(startKey, endKey)
	defer it.Close()

	res := types.QueryPostListResponse{Blogs: make([]*types.BlogItem, 0, endID-startID+1), Total: total}

	for ; it.Valid(); it.Next() {
		var blog types.BlogItem
		// 你上面 AppendBlog 用的是 k.cdc.MustMarshal，这里对应 MustUnmarshal
		k.cdc.MustUnmarshal(it.Value(), &blog)
		// 若你有 utils.UnmarshalTo 也可用：
		// blog, err := utils.UnmarshalTo[types.BlogItem](it.Value())
		// if err != nil { return nil, err }

		res.Blogs = append(res.Blogs, &blog)
	}

	return &res, nil
}
