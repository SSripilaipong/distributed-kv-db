package strm

import (
	"context"
	"distributed-kv-db/common/rslt"
)

func mapSliceWithContextAndXs[A, B any](ctx context.Context, xs []A) <-chan B {
	return MapResultFromSlice(mapSliceFDummy[A, B], ctx, xs)
}

func mapSliceWithXs[A, B any](xs []A) <-chan B {
	return mapSliceWithContextAndXs[A, B](context.Background(), xs)
}

func mapSliceWithFAndXs[A, B any](f func(A) rslt.Of[B], xs []A) <-chan B {
	return MapResultFromSlice[A, B](f, context.Background(), xs)
}
