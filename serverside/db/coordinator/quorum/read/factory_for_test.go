package read

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func newFuncWithDiscoverNodesAndReadNodesDataToChannelsAndLatestData[Key, Data, Node any](discoverNodes quorum.DiscoverNodes[Key, Node], readNodesDataToChannel func(context.Context, Key, []Node) <-chan Data, latestData func([]Data) rslt.Of[Data]) quorum.ReadFunc[Key, Data] {
	return newFunc(discoverNodes, readNodesDataToChannel, latestData)
}

func newFuncWithDiscoverNodesAndReadNodesDataToChannels[Key, Data, Node any](discoverNodes quorum.DiscoverNodes[Key, Node], readNodesDataToChannel func(context.Context, Key, []Node) <-chan Data) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoverNodesAndReadNodesDataToChannelsAndLatestData(discoverNodes, readNodesDataToChannel, latestDataDummy[Data])
}

func newFuncWithReadNodesDataToChannels[Key, Data, Node any](readNodesDataToChannel func(context.Context, Key, []Node) <-chan Data) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoverNodesAndReadNodesDataToChannelsAndLatestData(discoverNodesFuncDummy[Key, Node], readNodesDataToChannel, latestDataDummy[Data])
}

func newFuncWithDiscoverNodes[Key, Data, Node any](discoverNodes quorum.DiscoverNodes[Key, Node]) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoverNodesAndReadNodesDataToChannels(discoverNodes, readNodesDataToChannelDummy[Key, Data, Node])
}

func newFuncWithLatestData[Key, Data, Node any](latestData func([]Data) rslt.Of[Data]) quorum.ReadFunc[Key, Data] {
	return newFunc(discoverNodesFuncDummy[Key, Node], readNodesDataToChannelDummy[Key, Data, Node], latestData)
}

func readNodesDataToChannelWithNodes[Key, Data any, Node quorum.ReadNode[Key, Data]](nodes []Node) <-chan Data {
	return readNodesDataToChannelWithContextAndNodes[Key, Data](context.Background(), nodes)
}

func readNodesDataToChannelWithContextAndNodes[Key, Data any, Node quorum.ReadNode[Key, Data]](ctx context.Context, nodes []Node) <-chan Data {
	return readNodesDataToChannel[Key, Data](ctx, typ.Zero[Key](), nodes)
}

func readWithKeyAndNodes[Key, Data any, Node quorum.ReadNode[Key, Data]](key Key, nodes []Node) <-chan Data {
	return readNodesDataToChannel[Key, Data](context.Background(), key, nodes)
}
