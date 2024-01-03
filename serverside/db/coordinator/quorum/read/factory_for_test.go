package read

import (
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func newFuncWithDiscoveryAndReadNodesToChannelsAndLatestData[Key, Data any](discovery quorum.Discovery[Key, Data], readNodesToChannel func([]quorum.Node[Key, Data]) <-chan Data, latestData func([]Data) Data) quorum.ReadFunc[Key, Data] {
	return newFunc(discovery, readNodesToChannel, latestData)
}

func newFuncWithDiscoveryAndReadNodesToChannels[Key, Data any](discovery quorum.Discovery[Key, Data], readNodesToChannel func([]quorum.Node[Key, Data]) <-chan Data) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoveryAndReadNodesToChannelsAndLatestData(discovery, readNodesToChannel, latestDataDummy[Data])
}

func newFuncWithDiscovery[Key, Data any](discovery quorum.Discovery[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFuncWithDiscoveryAndReadNodesToChannels(discovery, readNodesToChannelDummy[Key, Data])
}
