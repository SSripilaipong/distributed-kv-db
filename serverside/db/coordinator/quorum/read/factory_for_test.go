package read

import "distributed-kv-db/serverside/db/coordinator/quorum"

func newFuncWithDiscoveryAndReadNodesToChannels[Key, Data any](discovery quorum.Discovery[Key, Data], readNodesToChannel func([]quorum.Node[Key, Data]) <-chan Data) quorum.ReadFunc[Key, Data] {
	return newFunc(discovery, readNodesToChannel)
}

func newFuncWithReadNodesToChannels[Key, Data any](readNodesToChannel func([]quorum.Node[Key, Data]) <-chan Data) quorum.ReadFunc[Key, Data] {
	return newFunc(discovery[Key, Data](), readNodesToChannel)
}
