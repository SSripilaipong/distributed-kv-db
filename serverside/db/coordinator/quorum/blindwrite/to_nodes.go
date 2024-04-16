package blindwrite

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/tuples"
	peerBlindWrite "distributed-kv-db/serverside/db/coordinator/peer/blindwrite"
	quorumFilter "distributed-kv-db/serverside/db/coordinator/quorum/filter"
)

func ToNodes[Key, Data any, Node peerBlindWrite.WritableNode[Key, Data]](nReplicas uint) func(ctx context.Context, nodes []Node, key Key, data Data) error {
	return composeToNodes[Key, Data, Node](
		quorumFilter.ErrorChannel(nReplicas),
		peerBlindWrite.ToNodes[Key, Data, Node](),
	)
}

func composeToNodes[Key, Data, Node any](
	qFilter func(<-chan error) error,
	pBlindWrite func(ctx context.Context, nodes []Node, key Key, data Data) <-chan error,
) func(ctx context.Context, nodes []Node, key Key, data Data) error {
	return tuples.ExplodeFn4(fn.Compose(
		qFilter, tuples.Fn4(pBlindWrite),
	))
}
