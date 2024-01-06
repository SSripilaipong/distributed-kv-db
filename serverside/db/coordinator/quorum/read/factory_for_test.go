package read

import (
	"context"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func newFuncWithDiscoveryAndReadNodesDataToChannelsAndLatestData[Key, Data any](discovery quorum.Discovery[Key, Data], readNodesDataToChannel func(context.Context, []quorum.Node[Key, Data]) <-chan Data, latestData func([]Data) Data) quorum.ReadFunc[Key, Data] {
	return newFunc(discovery, readNodesDataToChannel, latestData)
}

func newFuncWithDiscoveryAndReadNodesDataToChannels[Key, Data any](discovery quorum.Discovery[Key, Data], readNodesDataToChannel func(context.Context, []quorum.Node[Key, Data]) <-chan Data) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoveryAndReadNodesDataToChannelsAndLatestData(discovery, readNodesDataToChannel, latestDataDummy[Data])
}

func newFuncWithDiscovery[Key, Data any](discovery quorum.Discovery[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoveryAndReadNodesDataToChannels(discovery, readNodesDataToChannelDummy[Key, Data])
}

func newFuncWithLatestData[Key, Data any](latestData func([]Data) Data) quorum.ReadFunc[Key, Data] {
	return newFunc(discoveryDummy[Key, Data](), readNodesDataToChannelDummy[Key, Data], latestData)
}
