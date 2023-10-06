package server

import "distributed-kv-db/serverside/db/contract"

type Func func(port int) contract.Controller
