package slc

import (
	"distributed-kv-db/common/rslt"
	"errors"
)

func ReduceFn[S ~[]E, E any](f func(E, E) E) func(xs S) rslt.Of[E] {
	return func(xs S) rslt.Of[E] {
		if len(xs) == 0 {
			return rslt.Error[E](errors.New("empty slice"))
		}
		return rslt.Value(Fold(f, xs[0], xs[1:]))
	}
}
