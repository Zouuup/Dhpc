package keeper

import (
	"github.com/DhpcChain/Dhpc/x/request/types"
)

var _ types.QueryServer = Keeper{}
