package filter

import (
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/zd"
)

var nToQuorum = fn.Compose(zd.Successor, zd.Half)
