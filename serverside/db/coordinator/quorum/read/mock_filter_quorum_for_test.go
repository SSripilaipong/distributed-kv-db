package read

import (
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
)

func filterQuorumDummy[Data any](_ <-chan Data) rslt.Of[[]Data] {
	return rslt.Value(typ.Zero[[]Data]())
}

func filterQuorumCaptureXs[Data any](xs *<-chan Data) func(<-chan Data) rslt.Of[[]Data] {
	return func(d <-chan Data) rslt.Of[[]Data] {
		*xs = d
		return filterQuorumDummy[Data](d)
	}
}

func filterQuorumWithResult[Data any](result rslt.Of[[]Data]) func(<-chan Data) rslt.Of[[]Data] {
	return func(data <-chan Data) rslt.Of[[]Data] {
		return result
	}
}
