package readlatest

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/peer/discovery"
	peerRead "distributed-kv-db/serverside/db/coordinator/quorum/read"
)

func newFuncWithDiscoverNodesAndReadNodesDataToChannelsAndLatestData[Key, Data, Node any](discoverNodes discovery.Func[Key, Node], readQuorumOfNodesData func(context.Context, Key, []Node) rslt.Of[[]Data], latestData func([]Data) rslt.Of[Data]) peerRead.Func[Key, Data] {
	return newFunc(discoverNodes, readQuorumOfNodesData, latestData)
}

func newFuncWithDiscoverNodesAndReadQuorumOfNodesData[Key, Data, Node any](discoverNodes discovery.Func[Key, Node], readQuorumOfNodesData func(context.Context, Key, []Node) rslt.Of[[]Data]) peerRead.Func[Key, Data] {
	return newFuncWithDiscoverNodesAndReadNodesDataToChannelsAndLatestData(discoverNodes, readQuorumOfNodesData, latestDataDummy[Data])
}

func newFuncWithReadQuorumOfNodesData[Key, Data, Node any](readNodesDataToChannel func(context.Context, Key, []Node) rslt.Of[[]Data]) peerRead.Func[Key, Data] {
	return newFuncWithDiscoverNodesAndReadNodesDataToChannelsAndLatestData(discoverNodesFuncDummy[Key, Node], readNodesDataToChannel, latestDataDummy[Data])
}

func newFuncWithReadQuorumOfNodesDataAndLatestData[Key, Data, Node any](readQuorumOfNodesData func(context.Context, Key, []Node) rslt.Of[[]Data], latestData func([]Data) rslt.Of[Data]) peerRead.Func[Key, Data] {
	return newFuncWithDiscoverNodesAndReadNodesDataToChannelsAndLatestData(discoverNodesFuncDummy[Key, Node], readQuorumOfNodesData, latestData)
}

func newFuncWithDiscoverNodes[Key, Data, Node any](discoverNodes discovery.Func[Key, Node]) peerRead.Func[Key, Data] {
	return newFuncWithDiscoverNodesAndReadQuorumOfNodesData(discoverNodes, readQuorumOfNodesDataDummy[Key, Data, Node])
}

func newFuncWithLatestData[Key, Data, Node any](latestData func([]Data) rslt.Of[Data]) peerRead.Func[Key, Data] {
	return newFunc(discoverNodesFuncDummy[Key, Node], readQuorumOfNodesDataDummy[Key, Data, Node], latestData)
}
