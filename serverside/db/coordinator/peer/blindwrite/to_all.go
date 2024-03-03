package blindwrite

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/strm"
)

func ToAll[Key, Data, Node any]() func(ctx context.Context, key Key, data Data, nodes []Node) <-chan error {
	return composeToAll[Key, Data, Node](nil)
}

func composeToAll[Key, Data, Node any](write writeFunc[Key, Data, Node]) func(ctx context.Context, key Key, data Data, nodes []Node) <-chan error {
	return func(ctx context.Context, key Key, data Data, nodes []Node) <-chan error {
		return strm.MapSlice(fn.WithArg3(ctx, key, data, write), ctx, nodes)
	}
}

type writeFunc[Key, Data, Node any] func(ctx context.Context, key Key, data Data, node Node) error
