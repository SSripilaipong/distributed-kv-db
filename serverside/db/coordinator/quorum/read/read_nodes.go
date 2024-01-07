package read

import (
	"context"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func readNodesDataToChannel[Key, Data any](ctx context.Context, key Key, nodes []quorum.Node[Key, Data]) <-chan Data {
	ch := make(chan Data)
	close(ch)
	for _, node := range nodes {
		node.Read(ctx, typ.Zero[Key]())
	}
	return ch
}
