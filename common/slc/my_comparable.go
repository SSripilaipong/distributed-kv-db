package slc

type myComparable[T any] interface {
	Compare(other T) int
}

func myComparableGreater[A myComparable[B], B any](x A, y B) bool {
	return x.Compare(y) > 0
}

func myComparableMax[T myComparable[T]](x T, y T) T {
	if myComparableGreater(x, y) {
		return x
	}
	return y
}
