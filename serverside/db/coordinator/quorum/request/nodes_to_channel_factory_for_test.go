package request

import (
	"distributed-kv-db/common/fnopts"
	"distributed-kv-db/common/rslt"
)

type nodesToChannelTestDeps[Node, Data any] struct {
	request func(Node) rslt.Of[Data]
}

func nodesToChannelForTest[Node, Data any](options ...func(deps *nodesToChannelTestDeps[Node, Data])) func(nodes []Node) <-chan Data {
	deps := fnopts.Apply(nodesToChannelTestDeps[Node, Data]{
		request: requestDummy[Node, Data],
	}, options)
	return NodesToChannel[Node, Data](deps.request)
}
