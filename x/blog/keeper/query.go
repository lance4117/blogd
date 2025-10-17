package keeper

import (
	"github.com/lance4117/blogd/x/blog/types"
)

var _ types.QueryServer = Keeper{}
