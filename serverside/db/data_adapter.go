package db

import "distributed-kv-db/serverside/db/coordinator/quorum"

type orderableDataAdapter struct{}

func (d orderableDataAdapter) Hash() string {
	//TODO implement me
	panic("implement me")
}

func (d orderableDataAdapter) Newness() int {
	//TODO implement me
	panic("implement me")
}

var _ quorum.Orderable = orderableDataAdapter{}
var _ quorum.Hashable = orderableDataAdapter{}
