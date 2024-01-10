package read

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
	"distributed-kv-db/serverside/db/coordinator/quorum/response"
)

func readNodesDataToChannel[Key, Data any, Node quorum.ReadNode[Key, Data]](ctx context.Context, key Key, nodes []Node) <-chan Data {
	return response.Channel(ctx, readNode[Key, Data, Node](ctx, key), nodes)
}

func readNode[Key, Data any, Node quorum.ReadNode[Key, Data]](ctx context.Context, key Key) func(node Node) rslt.Of[Data] {
	return func(node Node) rslt.Of[Data] {
		return node.Read(ctx, key)
	}
}
