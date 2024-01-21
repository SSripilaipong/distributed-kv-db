package server

import (
	"distributed-kv-db/serverside/db/contract"
	"distributed-kv-db/serverside/db/coordinator/usecase/getvalue"
	"distributed-kv-db/serverside/db/coordinator/usecase/setvalue"
	"distributed-kv-db/serverside/db/server/grpc"
)

func New(getValue getvalue.Func, setValue setvalue.Func) Func {
	runServer := grpc.New(getValue, setValue)

	return func(port int) contract.Controller {
		runningServer := runServer(port)
		if runningServer.IsOk() {
			// TODO: implement this
		}
		return nil
	}
}
