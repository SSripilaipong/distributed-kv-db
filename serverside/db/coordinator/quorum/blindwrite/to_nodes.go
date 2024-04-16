package blindwrite

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/tuples"
	peerBlindWrite "distributed-kv-db/serverside/db/coordinator/peer/blindwrite"
)

func ToNodes[Key, Data, Node peerBlindWrite.WritableNode[Key, Data]](nReplicas uint) func(ctx context.Context, nodes []Node, key Key, data Data) error {
	return composeToNodes[Key, Data, Node](
		nil,                                     // TODO inject
		peerBlindWrite.ToAll[Key, Data, Node](), // TODO inject
	)
}

func composeToNodes[Key, Data, Node any](
	qFilter func(<-chan error) error,
	pBlindWrite func(ctx context.Context, key Key, data Data, nodes []Node) <-chan error,
) func(ctx context.Context, nodes []Node, key Key, data Data) error {
	return tuples.ExplodeFn4(fn.Compose(
		qFilter, tuples.Fn4(swapArg(pBlindWrite)),
	))
}

func swapArg[A, B, C, D, E any](f func(a A, b B, c C, d D) E) func(a A, d D, b B, c C) E {
	return func(a A, d D, b B, c C) E {
		return f(a, b, c, d)
	}
}
