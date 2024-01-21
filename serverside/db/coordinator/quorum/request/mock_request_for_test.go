package request

import (
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
)

func requestDummy[Node, Data any](_ Node) rslt.Of[Data] {
	return rslt.Value(typ.Zero[Data]())
}
