package read

import (
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func readNodeDataToChannelDummy[Key, Data any](_ []quorum.Node[Key, Data]) <-chan Data {
	return chn.Repeat(typ.Zero[Data]())
}

func readNodeDataToChannelCaptureNodes[Key, Data any](nodes *[]quorum.Node[Key, Data]) func([]quorum.Node[Key, Data]) <-chan Data {
	return func(n []quorum.Node[Key, Data]) <-chan Data {
		*nodes = n
		return chn.Repeat(typ.Zero[Data]())
	}
}

func readNodeDataToChannelWithResult[Key, Data any](ch <-chan Data) func([]quorum.Node[Key, Data]) <-chan Data {
	return fn.Const[[]quorum.Node[Key, Data]](ch)
}
