package read

import (
	"context"
	"distributed-kv-db/common/fnopts"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
)

type nodesToDataSliceTestRunnerDeps[Key, Node any] struct {
	ctx   context.Context
	key   Key
	nodes []Node
}

func runNodesToDataSliceForTest[Key, Data, Node any](
	f func(ctx context.Context, key Key, nodes []Node) rslt.Of[[]Data],
	options ...func(*nodesToDataSliceTestRunnerDeps[Key, Node]),
) rslt.Of[[]Data] {
	deps := fnopts.Apply(nodesToDataSliceTestRunnerDeps[Key, Node]{
		ctx:   context.Background(),
		key:   typ.Zero[Key](),
		nodes: []Node{},
	}, options)
	return f(deps.ctx, deps.key, deps.nodes)
}

func withKey[Key, Node any](key Key) func(*nodesToDataSliceTestRunnerDeps[Key, Node]) {
	return func(deps *nodesToDataSliceTestRunnerDeps[Key, Node]) {
		deps.key = key
	}
}
