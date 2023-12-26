package getvalue

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/result"
	"distributed-kv-db/serverside/db/coordinator"
)

func New(readRepair readRepairFunc) coordinator.GetValueFunc {
	return fn.Uncurry(func(ctx context.Context) partialFunc {
		return fn.Compose3(
			result.Fmap(responseFromValue), fn.Ctx(ctx, readRepair), keyOfRequest,
		)
	})
}

type partialFunc = func(coordinator.GetValueRequest) result.Of[coordinator.GetValueResponse]
