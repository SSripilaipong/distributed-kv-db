package read

import (
	"context"
	"distributed-kv-db/common/fnopts"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/peer/discovery"
)

type nodesToDataSliceTestDeps[Key, Data, Node any] struct {
	quorumOfData  func(n int) func(<-chan Data) rslt.Of[[]Data]
	readNodes     func(ctx context.Context, key Key, nodes []Node) <-chan Data
	discoverNodes discovery.Func[Key, Node]
}

func newNodesToDataSliceForTest[Key, Data, Node any](options ...func(*nodesToDataSliceTestDeps[Key, Data, Node])) func(ctx context.Context, key Key) rslt.Of[[]Data] {
	deps := fnopts.Apply(nodesToDataSliceTestDeps[Key, Data, Node]{
		quorumOfData:  nil,
		readNodes:     readNodesDummy[Key, Data, Node],
		discoverNodes: discoverNodesDummy[Key, Node],
	}, options)
	return composeNodesToDataSlice[Key, Data, Node](deps.quorumOfData, deps.readNodes, deps.discoverNodes)
}

func withDiscoverNodes[Key, Data, Node any](f discovery.Func[Key, Node]) func(*nodesToDataSliceTestDeps[Key, Data, Node]) {
	return func(deps *nodesToDataSliceTestDeps[Key, Data, Node]) {
		deps.discoverNodes = f
	}
}

func withReadNodes[Key, Data, Node any](readNodes func(ctx context.Context, key Key, nodes []Node) <-chan Data) func(*nodesToDataSliceTestDeps[Key, Data, Node]) {
	return func(deps *nodesToDataSliceTestDeps[Key, Data, Node]) {
		deps.readNodes = readNodes
	}
}
