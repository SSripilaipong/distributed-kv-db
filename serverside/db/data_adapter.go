package db

import (
	"distributed-kv-db/serverside/db/data/temporal"
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

var _ temporal.Data = orderableDataAdapter{}
