package syncval

import "sync"

func Func[A, B any](f func(A) B) func(A) B {
	var lock sync.Mutex
	return func(x A) B {
		lock.Lock()
		defer lock.Unlock()
		return f(x)
	}
}
