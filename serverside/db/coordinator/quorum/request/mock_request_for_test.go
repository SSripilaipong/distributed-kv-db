package request

import (
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
)

func requestDummy[Node, Data any](_ Node) rslt.Of[Data] {
	return rslt.Value(typ.Zero[Data]())
}

func requestCaptureAllNode[Node, Data any](allNode *[]Node) func(Node) rslt.Of[Data] {
	return func(node Node) rslt.Of[Data] {
		*allNode = append(*allNode, node)
		return requestDummy[Node, Data](node)
	}
}
