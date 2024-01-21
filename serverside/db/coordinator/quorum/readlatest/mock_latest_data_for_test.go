package readlatest

import (
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
)

func latestDataDummy[Data any](_ []Data) rslt.Of[Data] {
	return rslt.Value(typ.Zero[Data]())
}

func latestDataCaptureXs[Data any](xs *[]Data) func([]Data) rslt.Of[Data] {
	return func(ys []Data) rslt.Of[Data] {
		*xs = ys
		return latestDataDummy[Data](nil)
	}
}

func latestDataWithResult[Data any](result Data) func([]Data) rslt.Of[Data] {
	return fn.Const[[]Data](rslt.Value(result))
}
