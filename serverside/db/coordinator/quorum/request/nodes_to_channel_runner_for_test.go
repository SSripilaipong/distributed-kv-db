package request

import "distributed-kv-db/common/fnopts"

type nodeToChannelsTestRunnerDeps[Node any] struct {
	nodes []Node
}

func runNodesToChannelForTest[Node, Data any](
	f func(nodes []Node) <-chan Data,
	options ...func(*nodeToChannelsTestRunnerDeps[Node]),
) <-chan Data {
	deps := fnopts.Apply(nodeToChannelsTestRunnerDeps[Node]{
		nodes: []Node{},
	}, options)
	return f(deps.nodes)
}

func withNodes[Node any](nodes []Node) func(deps *nodeToChannelsTestRunnerDeps[Node]) {
	return func(deps *nodeToChannelsTestRunnerDeps[Node]) {
		deps.nodes = nodes
	}
}
