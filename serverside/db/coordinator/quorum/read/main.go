package read

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func New[Key, Data any](discovery quorum.Discovery[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFunc(
		discovery,
		nil,
	)
}

func newFunc[Key, Data any](discovery quorum.Discovery[Key, Data], readNodesToChannel func([]quorum.Node[Key, Data]) <-chan Data) quorum.ReadFunc[Key, Data] {
	return func(ctx context.Context, key Key) rslt.Of[Data] {
		rslt.Fmap(readNodesToChannel)(discovery.Nodes(nil))
		return rslt.Error[Data](nil)
	}
}
