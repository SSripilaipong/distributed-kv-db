package slc

import "distributed-kv-db/common/typ"

func ReduceFn[S ~[]E, E any](f func(E, E) E) func(xs S) E {
	return func(xs S) E {
		if len(xs) == 0 {
			return typ.Zero[E]()
		}
		return Fold(f, xs[0], xs[1:])
	}
}
