package strm

import (
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/syncval"
	"distributed-kv-db/common/typ"
)

func mapSliceFDummy[A, B any](_ A) rslt.Of[B] {
	return rslt.Value(typ.Zero[B]())
}

func mapSliceFCaptureAllX[A, B any](allX *[]A) func(A) rslt.Of[B] {

	return syncval.Func(func(x A) rslt.Of[B] {
		*allX = append(*allX, x)
		return mapSliceFDummy[A, B](x)
	})
}

func mapSliceFWithAllResult[A, B any](allResult []rslt.Of[B]) func(A) rslt.Of[B] {
	var count int
	return syncval.Func(func(A) rslt.Of[B] {
		result := allResult[count]
		count++
		return result
	})
}
