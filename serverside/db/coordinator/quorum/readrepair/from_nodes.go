package readrepair

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
)

func FromNodes[Key any, Data any, Node any](
	qMergeRead func(ctx context.Context, nodes []Node, key Key) rslt.Of[Data],
	qBlindWrite func(ctx context.Context, nodes []Node, key Key, data Data) error,
) func(context.Context, []Node, Key) rslt.Of[Data] {
	return composeFromNodes[Key, Data, Node](qMergeRead, qBlindWrite) // TODO inject real functions
}

func composeFromNodes[Key any, Data any, Node any](
	qMergeRead func(ctx context.Context, nodes []Node, key Key) rslt.Of[Data],
	qBlindWrite func(ctx context.Context, nodes []Node, key Key, data Data) error,
) func(context.Context, []Node, Key) rslt.Of[Data] {
	return func(ctx context.Context, nodes []Node, key Key) rslt.Of[Data] {
		writeFn := fmapResultToError(fn.WithArg3(ctx, nodes, key, qBlindWrite))

		result := qMergeRead(ctx, nodes, key)
		return rslt.ResultOrError(result, writeFn(result))
	}
}
