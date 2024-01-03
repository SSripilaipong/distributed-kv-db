package read

import (
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func readNodesToChannelCaptureNodes[Key, Data any](nodes *[]quorum.Node[Key, Data]) func([]quorum.Node[Key, Data]) <-chan Data {
	return func(n []quorum.Node[Key, Data]) <-chan Data {
		*nodes = n
		return chn.Repeat(typ.Zero[Data]())
	}
}

func readNodesToChannelWithResult[Key, Data any](ch <-chan Data) func([]quorum.Node[Key, Data]) <-chan Data {
	return func(n []quorum.Node[Key, Data]) <-chan Data {
		return ch
	}
}
