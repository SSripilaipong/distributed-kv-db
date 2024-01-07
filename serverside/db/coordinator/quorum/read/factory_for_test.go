package read

import (
	"context"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func newFuncWithDiscoveryAndReadNodesDataToChannelsAndLatestData[Key, Data any](discovery quorum.Discovery[Key, Data], readNodesDataToChannel func(context.Context, Key, []quorum.Node[Key, Data]) <-chan Data, latestData func([]Data) Data) quorum.ReadFunc[Key, Data] {
	return newFunc(discovery, readNodesDataToChannel, latestData)
}

func newFuncWithDiscoveryAndReadNodesDataToChannels[Key, Data any](discovery quorum.Discovery[Key, Data], readNodesDataToChannel func(context.Context, Key, []quorum.Node[Key, Data]) <-chan Data) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoveryAndReadNodesDataToChannelsAndLatestData(discovery, readNodesDataToChannel, latestDataDummy[Data])
}

func newFuncWithReadNodesDataToChannels[Key, Data any](readNodesDataToChannel func(context.Context, Key, []quorum.Node[Key, Data]) <-chan Data) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoveryAndReadNodesDataToChannelsAndLatestData(discoveryDummy[Key, Data](), readNodesDataToChannel, latestDataDummy[Data])
}

func newFuncWithDiscovery[Key, Data any](discovery quorum.Discovery[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoveryAndReadNodesDataToChannels(discovery, readNodesDataToChannelDummy[Key, Data])
}

func newFuncWithLatestData[Key, Data any](latestData func([]Data) Data) quorum.ReadFunc[Key, Data] {
	return newFunc(discoveryDummy[Key, Data](), readNodesDataToChannelDummy[Key, Data], latestData)
}

func readNodesDataToChannelWithNodes[Key, Data any](nodes []quorum.Node[Key, Data]) <-chan Data {
	return readNodesDataToChannel[Key, Data](context.Background(), typ.Zero[Key](), nodes)
}
