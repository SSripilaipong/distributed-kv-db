package prll

import (
	"distributed-kv-db/common/grt"
	"distributed-kv-db/common/slc"
	"distributed-kv-db/common/wgrp"
	"sync"
)

func ApplySlice[A any](wg *sync.WaitGroup, f func(A), xs []A) {
	wg.Add(len(xs))
	slc.Do(grt.Do(wgrp.MustDoneAfterDo(wg, f)), xs)
}
