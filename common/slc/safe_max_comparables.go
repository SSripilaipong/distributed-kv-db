package slc

import (
	"distributed-kv-db/common/typ"
)

func SafeMaxComparables[T myComparable[T]](xs []T) T {
	return Fold(myComparableMax[T], typ.Zero[T](), xs)
}
