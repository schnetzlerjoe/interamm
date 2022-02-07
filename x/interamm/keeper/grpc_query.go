package keeper

import (
	"github.com/schnetzlerjoe/interamm/x/interamm/types"
)

var _ types.QueryServer = Keeper{}
