package slc

import (
	"cmp"
	"slices"
)

func Sorted[T cmp.Ordered](xs []T) []T {
	ys := slices.Clone(xs)
	slices.Sort(ys)
	return ys
}
