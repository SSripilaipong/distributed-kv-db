package grpc

import (
	"distributed-kv-db/serverside/db/coordinator/usecase/getvalue"
	"distributed-kv-db/serverside/db/coordinator/usecase/setvalue"
)

func newWithSetValueFunc(setValue setvalue.Func) Func {
	return New(nil, setValue)
}

func newWithGetValueFunc(getValue getvalue.Func) Func {
	return New(getValue, nil)
}
