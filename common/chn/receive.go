package chn

import (
	"distributed-kv-db/common/rslt"
	"errors"
)

func ReceiveNoWait[T any](ch <-chan T) rslt.Of[T] {
	select {
	case x, ok := <-ch:
		if !ok {
			return rslt.Error[T](errors.New("channel closed"))
		}
		return rslt.Value(x)
	default:
		return rslt.Error[T](errors.New("no data"))
	}
}

func FirstNFunc[T any](n int) func(ch <-chan T) rslt.Of[[]T] {
	return func(ch <-chan T) rslt.Of[[]T] {
		var result []T
		for x := range ch {
			result = append(result, x)

			if len(result) >= n {
				break
			}
		}
		if len(result) < n {
			return rslt.Error[[]T](errors.New("channel closed"))
		}
		return rslt.Value(result)
	}
}
