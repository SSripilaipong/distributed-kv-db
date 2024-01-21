package read

import "context"

func readNodesDummy[Key, Data, Node any](_ context.Context, _ Key, _ []Node) <-chan Data {
	return make(chan Data, 2)
}

func readNodesCaptureNodes[Key, Data, Node any](nodes *[]Node) func(ctx context.Context, key Key, nodes []Node) <-chan Data {
	return func(c context.Context, k Key, n []Node) <-chan Data {
		*nodes = n
		return readNodesDummy[Key, Data, Node](c, k, n)
	}
}
