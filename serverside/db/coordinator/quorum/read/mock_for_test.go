package read

import "distributed-kv-db/serverside/db/coordinator/quorum"

func readNodesToChannelCaptureNodes(nodes *[]quorum.Node[int, int]) func([]quorum.Node[int, int]) <-chan int {
	return func(n []quorum.Node[int, int]) <-chan int {
		*nodes = n
		return make(chan int, 2)
	}
}
