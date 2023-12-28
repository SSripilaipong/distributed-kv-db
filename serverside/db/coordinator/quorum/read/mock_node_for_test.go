package read

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

type nodeWithId[Key, Data any] struct {
	id int
}

func (n nodeWithId[Key, Data]) Read(_ context.Context, _ Key) rslt.Of[Data] {
	return rslt.Value(typ.Zero[Data]())
}

var _ quorum.Node[int, int] = nodeWithId[int, int]{}
