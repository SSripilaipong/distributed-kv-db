package zd

import (
	"distributed-kv-db/common/rslt"
	"fmt"
)

func MustBeMoreThan(m int) func(n int) rslt.Of[int] {
	return func(n int) rslt.Of[int] {
		if n <= m {
			return rslt.Error[int](fmt.Errorf("n must be more than %d, got %d", m, n))
		}
		return rslt.Value(n)
	}
}
