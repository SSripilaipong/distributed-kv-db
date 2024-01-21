package getvalue

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
)

func New(readRepair readRepairFunc) Func {
	return fn.Uncurry(func(ctx context.Context) partialFunc {
		return fn.Compose3(
			rslt.Fmap(responseFromValue), fn.Ctx(ctx, readRepair), keyOfRequest,
		)
	})
}

type partialFunc = func(Request) rslt.Of[Response]
