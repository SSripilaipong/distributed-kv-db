package read

import (
	"context"
	"distributed-kv-db/common/fnopts"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
)

type nodesToDataSliceTestRunnerDeps[Key, Node any] struct {
	ctx context.Context
	key Key
}

func runNodesToDataSliceForTest[Key, Data, Node any](
	f func(ctx context.Context, key Key) rslt.Of[[]Data],
	options ...func(*nodesToDataSliceTestRunnerDeps[Key, Node]),
) rslt.Of[[]Data] {
	deps := fnopts.Apply(nodesToDataSliceTestRunnerDeps[Key, Node]{
		ctx: context.Background(),
		key: typ.Zero[Key](),
	}, options)
	return f(deps.ctx, deps.key)
}

func withKey[Key, Node any](key Key) func(*nodesToDataSliceTestRunnerDeps[Key, Node]) {
	return func(deps *nodesToDataSliceTestRunnerDeps[Key, Node]) {
		deps.key = key
	}
}

func withContext[Key, Node any](ctx context.Context) func(*nodesToDataSliceTestRunnerDeps[Key, Node]) {
	return func(deps *nodesToDataSliceTestRunnerDeps[Key, Node]) {
		deps.ctx = ctx
	}
}
