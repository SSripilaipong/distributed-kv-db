package read

import (
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
)

func filterQuorumDummy[Data any](_ int) func(<-chan Data) rslt.Of[[]Data] {
	return func(data <-chan Data) rslt.Of[[]Data] {
		return rslt.Value(typ.Zero[[]Data]())
	}
}

func filterQuorumCaptureXs[Data any](xs *<-chan Data) func(n int) func(<-chan Data) rslt.Of[[]Data] {
	return func(n int) func(<-chan Data) rslt.Of[[]Data] {
		return func(d <-chan Data) rslt.Of[[]Data] {
			*xs = d
			return filterQuorumDummy[Data](n)(d)
		}
	}
}

func filterQuorumCaptureN[Data any](n *int) func(n int) func(<-chan Data) rslt.Of[[]Data] {
	return func(n_ int) func(<-chan Data) rslt.Of[[]Data] {
		return func(d <-chan Data) rslt.Of[[]Data] {
			*n = n_
			return filterQuorumDummy[Data](n_)(d)
		}
	}
}

func filterQuorumWithResult[Data any](result rslt.Of[[]Data]) func(n int) func(<-chan Data) rslt.Of[[]Data] {
	return func(n int) func(<-chan Data) rslt.Of[[]Data] {
		return func(data <-chan Data) rslt.Of[[]Data] {
			return result
		}
	}
}
