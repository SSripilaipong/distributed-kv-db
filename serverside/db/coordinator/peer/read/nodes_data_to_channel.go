package read

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/strm"
)

func NodesDataToChannel[Key, Data any, Node ReadableNode[Key, Data]](ctx context.Context, key Key, nodes []Node) <-chan Data {
	return strm.MapResultFromSlice(nodeToData[Key, Data, Node](ctx, key), ctx, nodes)
}

func nodeToData[Key, Data any, Node ReadableNode[Key, Data]](ctx context.Context, key Key) func(node Node) rslt.Of[Data] {
	return func(node Node) rslt.Of[Data] {
		return node.Read(ctx, key)
	}
}
