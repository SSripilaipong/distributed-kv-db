package read

import (
	"context"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func newFuncWithDiscoverNodesAndReadNodesDataToChannelsAndLatestData[Key, Data any](discoverNodes quorum.DiscoverNodes[Key, quorum.ReadableNode[Key, Data]], readNodesDataToChannel func(context.Context, Key, []quorum.ReadableNode[Key, Data]) <-chan Data, latestData func([]Data) Data) quorum.ReadFunc[Key, Data] {
	return newFunc(discoverNodes, readNodesDataToChannel, latestData)
}

func newFuncWithDiscoverNodesAndReadNodesDataToChannels[Key, Data any](discoverNodes quorum.DiscoverNodes[Key, quorum.ReadableNode[Key, Data]], readNodesDataToChannel func(context.Context, Key, []quorum.ReadableNode[Key, Data]) <-chan Data) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoverNodesAndReadNodesDataToChannelsAndLatestData(discoverNodes, readNodesDataToChannel, latestDataDummy[Data])
}

func newFuncWithReadNodesDataToChannels[Key, Data any](readNodesDataToChannel func(context.Context, Key, []quorum.ReadableNode[Key, Data]) <-chan Data) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoverNodesAndReadNodesDataToChannelsAndLatestData(discoverNodesFuncDummy[Key, Data], readNodesDataToChannel, latestDataDummy[Data])
}

func newFuncWithDiscoverNodes[Key, Data any](discoverNodes quorum.DiscoverNodes[Key, quorum.ReadableNode[Key, Data]]) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoverNodesAndReadNodesDataToChannels(discoverNodes, readNodesDataToChannelDummy[Key, Data])
}

func newFuncWithLatestData[Key, Data any](latestData func([]Data) Data) quorum.ReadFunc[Key, Data] {
	return newFunc(discoverNodesFuncDummy[Key, Data], readNodesDataToChannelDummy[Key, Data], latestData)
}

func readNodesDataToChannelWithNodes[Key, Data any](nodes []quorum.ReadableNode[Key, Data]) <-chan Data {
	return readNodesDataToChannelWithContextAndNodes(context.Background(), nodes)
}

func readNodesDataToChannelWithContextAndNodes[Key, Data any](ctx context.Context, nodes []quorum.ReadableNode[Key, Data]) <-chan Data {
	return readNodesDataToChannel[Key, Data](ctx, typ.Zero[Key](), nodes)
}

func readWithKeyAndNodes[Key, Data any](key Key, nodes []quorum.ReadableNode[Key, Data]) <-chan Data {
	return readNodesDataToChannel[Key, Data](context.Background(), key, nodes)
}
