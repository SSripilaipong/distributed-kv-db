package grpc

import "distributed-kv-db/serverside/db/coordinator"

func newWithSetValueFunc(setValue coordinator.SetValueFunc) Func {
	return New(nil, setValue)
}

func newWithGetValueFunc(getValue coordinator.GetValueFunc) Func {
	return New(getValue, nil)
}
