package keeper

import (
	"github.com/DhpcChain/Dhpc/x/user/types"
)

var _ types.QueryServer = Keeper{}
