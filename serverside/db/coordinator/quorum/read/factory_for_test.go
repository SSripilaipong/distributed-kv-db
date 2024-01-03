package read

import (
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func newFuncWithDiscoveryAndReadNodesToChannels[Key, Data any](discovery quorum.Discovery[Key, Data], readNodesToChannel func([]quorum.Node[Key, Data]) <-chan Data) quorum.ReadFunc[Key, Data] {
	return newFunc(discovery, readNodesToChannel)
}

func newFuncWithDiscovery[Key, Data any](discovery quorum.Discovery[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFunc(discovery, func([]quorum.Node[Key, Data]) <-chan Data {
		return chn.Repeat(typ.Zero[Data]())
	})
}
