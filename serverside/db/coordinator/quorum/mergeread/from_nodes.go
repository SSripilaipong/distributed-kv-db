package mergeread

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/slc"
	"distributed-kv-db/common/tuples"
)

func FromNodes[Key, Data, Node any](
	merge func(x, y Data) Data,
) func(ctx context.Context, node []Node, key Key) rslt.Of[Data] {
	return composeFromNodes[Key, Data, Node](
		merge, nil, // TODO inject qRead
	)
}

func composeFromNodes[Key, Data, Node any](
	merge func(x, y Data) Data,
	qRead func(ctx context.Context, nodes []Node, key Key) rslt.Of[[]Data],
) func(ctx context.Context, nodes []Node, key Key) rslt.Of[Data] {

	return tuples.ExplodeFn3(fn.Compose(
		rslt.Fmap(slc.ReduceFn[[]Data](merge)), tuples.Fn3(qRead),
	))
}
