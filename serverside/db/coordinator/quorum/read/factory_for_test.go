package read

import (
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func newFuncWithDiscoveryAndReadNodeDataToChannelsAndLatestData[Key, Data any](discovery quorum.Discovery[Key, Data], readNodeDataToChannel func([]quorum.Node[Key, Data]) <-chan Data, latestData func([]Data) Data) quorum.ReadFunc[Key, Data] {
	return newFunc(discovery, readNodeDataToChannel, latestData)
}

func newFuncWithDiscoveryAndReadNodeDataToChannels[Key, Data any](discovery quorum.Discovery[Key, Data], readNodeDataToChannel func([]quorum.Node[Key, Data]) <-chan Data) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoveryAndReadNodeDataToChannelsAndLatestData(discovery, readNodeDataToChannel, latestDataDummy[Data])
}

func newFuncWithDiscovery[Key, Data any](discovery quorum.Discovery[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoveryAndReadNodeDataToChannels(discovery, readNodeDataToChannelDummy[Key, Data])
}

func newFuncWithLatestData[Key, Data any](latestData func([]Data) Data) quorum.ReadFunc[Key, Data] {
	return newFunc(discoveryDummy[Key, Data](), readNodeDataToChannelDummy[Key, Data], latestData)
}
