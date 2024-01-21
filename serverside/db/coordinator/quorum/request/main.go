package request

import "distributed-kv-db/common/rslt"

func NodesToChannel[Node, Data any](request func(Node) rslt.Of[Data]) func(nodes []Node) <-chan Data {
	return func(nodes []Node) <-chan Data {
		return make(chan Data)
	}
}
