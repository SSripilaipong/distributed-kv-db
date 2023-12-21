package server

import (
	"distributed-kv-db/serverside/db/contract"
	"distributed-kv-db/serverside/db/coordinator"
	"distributed-kv-db/serverside/db/server/grpc"
)

func New(getValue coordinator.GetValueFunc, setValue coordinator.SetValueFunc) Func {
	runServer := grpc.New(getValue, setValue)

	return func(port int) contract.Controller {
		runningServer := runServer(port)
		if runningServer.IsOk() {
			// TODO: implement this
		}
		return nil
	}
}
