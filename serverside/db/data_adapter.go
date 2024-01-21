package db

import (
	"distributed-kv-db/serverside/db/coordinator/quorum/readlatest"
)

type orderableDataAdapter struct{}

func (d orderableDataAdapter) Hash() int {
	//TODO implement me
	panic("implement me")
}

func (d orderableDataAdapter) Newness() int {
	//TODO implement me
	panic("implement me")
}

var _ readlatest.Orderable = orderableDataAdapter{}
var _ readlatest.Hashable = orderableDataAdapter{}
