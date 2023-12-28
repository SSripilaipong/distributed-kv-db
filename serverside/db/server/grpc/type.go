package grpc

import "distributed-kv-db/common/rslt"

type Func func(port int) rslt.Of[RunningServer]
