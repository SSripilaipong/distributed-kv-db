package blindwrite

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/peer/blindwrite"
)

func New[Key, Data any, Node WritableNode[Key, Data]](nReplicas int, discoverNodes discoverNodesFunc[Key, Node]) func(ctx context.Context, key Key, data Data) error {
	return composeBlindWrite[Key, Data, Node](
		nil, // TODO inject
		blindwrite.ToAll[Key, Data, Node](),
		discoverNodes,
	)
}

func composeBlindWrite[Key, Data, Node any](quorumFilter quorumFilterFunc, writeNodes writeNodesFunc[Key, Data, Node], discoverNodes discoverNodesFunc[Key, Node]) func(ctx context.Context, key Key, data Data) error {
	return func(ctx context.Context, key Key, data Data) error {
		return nil
	}
}

type WritableNode[Key, Data any] interface {
	blindwrite.WritableNode[Key, Data]
}

type quorumFilterFunc func(<-chan error) error
type writeNodesFunc[Key, Data, Node any] func(ctx context.Context, key Key, data Data, nodes []Node) <-chan error
type discoverNodesFunc[Key, Node any] func(ctx context.Context, key Key) rslt.Of[[]Node]
