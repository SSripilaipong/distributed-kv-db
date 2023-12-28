package readrepair

import (
	"context"
	"distributed-kv-db/common/result"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func New[Key, Data any](_ quorum.Discovery[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFunc[Key, Data](nil)
}

func newFunc[Key, Data any](quorumRead quorum.ReadFunc[Key, Data]) quorum.ReadFunc[Key, Data] {
	return func(ctx context.Context, key Key) result.Of[Data] {
		quorumRead(nil, key)
		return result.Error[Data](nil)
	}
}
