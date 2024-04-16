package blindwrite

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/strm"
)

func ToAll[Key, Data any, Node WritableNode[Key, Data]]() func(ctx context.Context, key Key, data Data, nodes []Node) <-chan error {
	return composeToAll[Key, Data, Node](writeToNode[Key, Data, Node])
}

func composeToAll[Key, Data, Node any](
	write func(ctx context.Context, key Key, data Data, node Node) error,
) func(ctx context.Context, key Key, data Data, nodes []Node) <-chan error {
	return func(ctx context.Context, key Key, data Data, nodes []Node) <-chan error {
		return strm.MapSlice(fn.WithArg3(ctx, key, data, write), ctx, nodes)
	}
}
