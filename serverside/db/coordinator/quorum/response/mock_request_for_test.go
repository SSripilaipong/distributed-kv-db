package response

import (
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/syncval"
	"distributed-kv-db/common/typ"
)

func requestDummy[Response, Node any](_ Node) rslt.Of[Response] {
	return rslt.Value(typ.Zero[Response]())
}

func requestCaptureAllNodes[Response, Node any](nodes *[]Node) func(Node) rslt.Of[Response] {

	return syncval.Func(func(node Node) rslt.Of[Response] {
		*nodes = append(*nodes, node)
		return requestDummy[Response, Node](node)
	})
}

func requestWithResults[Response, Node any](results []rslt.Of[Response]) func(Node) rslt.Of[Response] {
	var count int
	return syncval.Func(func(node Node) rslt.Of[Response] {
		result := results[count]
		count++
		return result
	})
}
