package blindwrite

import "context"

type WritableNode[Key, Data any] interface {
	Write(ctx context.Context, key Key, data Data) error
}

func writeToNode[Key, Data any, Node WritableNode[Key, Data]](ctx context.Context, key Key, data Data, node Node) error {
	return node.Write(ctx, key, data)
}
