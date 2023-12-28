package readrepair

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func New[Key, Data any](_ quorum.Discovery[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFunc[Key, Data](nil, nil)
}

func newFunc[Key, Data any](quorumRead quorum.ReadFunc[Key, Data], quorumWrite quorum.WriteFunc[Key, Data]) quorum.ReadFunc[Key, Data] {
	return func(ctx context.Context, key Key) rslt.Of[Data] {
		rslt.Fmap(fn.Bind(key, fn.Ctx2(nil, quorumWrite)))(quorumRead(nil, key))
		return rslt.Error[Data](nil)
	}
}
