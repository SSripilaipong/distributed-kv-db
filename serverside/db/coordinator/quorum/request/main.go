package request

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/strm"
)

func NodesToChannel[Node, Data any](request func(Node) rslt.Of[Data]) func(nodes []Node) <-chan Data {
	return func(nodes []Node) <-chan Data {
		return strm.MapSlice(request, context.Background(), nodes)
	}
}
