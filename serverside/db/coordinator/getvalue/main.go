package getvalue

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator"
)

func New(readRepair readRepairFunc) coordinator.GetValueFunc {
	return fn.Uncurry(func(ctx context.Context) partialFunc {
		return fn.Compose3(
			rslt.Fmap(responseFromValue), fn.Ctx(ctx, readRepair), keyOfRequest,
		)
	})
}

type partialFunc = func(coordinator.GetValueRequest) rslt.Of[coordinator.GetValueResponse]
