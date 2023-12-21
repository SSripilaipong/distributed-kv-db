package grpc

import "distributed-kv-db/serverside/db/coordinator"

type getValueRequest struct {
	key string
}

func (r getValueRequest) Key() string {
	return r.key
}

type setValueRequest struct {
	key   string
	value string
}

func (r setValueRequest) Key() string {
	return r.key
}

func (r setValueRequest) Value() string {
	return r.value
}

var _ coordinator.GetValueRequest = getValueRequest{}
var _ coordinator.SetValueRequest = setValueRequest{}
