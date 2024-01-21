package read

import (
	"context"
	"distributed-kv-db/common/fnopts"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/peer/discovery"
)

type nodesToDataSliceTestDeps[Key, Data, Node any] struct {
	quorumOfData  func(n int) func(<-chan Data) rslt.Of[[]Data]
	readNode      func(ctx context.Context, key Key, nodes []Node) <-chan Data
	discoverNodes discovery.Func[Key, Node]
}

func newNodesToDataSliceForTest[Key, Data, Node any](options ...func(*nodesToDataSliceTestDeps[Key, Data, Node])) func(ctx context.Context, key Key, nodes []Node) rslt.Of[[]Data] {
	deps := fnopts.Apply(nodesToDataSliceTestDeps[Key, Data, Node]{
		quorumOfData:  nil,
		readNode:      nil,
		discoverNodes: nil,
	}, options)
	return composeNodesToDataSlice[Key, Data, Node](deps.quorumOfData, deps.readNode, deps.discoverNodes)
}

func withDiscoverNodes[Key, Data, Node any](f discovery.Func[Key, Node]) func(*nodesToDataSliceTestDeps[Key, Data, Node]) {
	return func(deps *nodesToDataSliceTestDeps[Key, Data, Node]) {
		deps.discoverNodes = f
	}
}
