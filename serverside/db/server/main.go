package server

import "distributed-kv-db/serverside/db/contract"

func New() Func {
	return func(port int) contract.Controller {
		return controller{}
	}
}
