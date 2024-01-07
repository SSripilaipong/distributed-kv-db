package read

import (
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/serverside/db/coordinator/quorum"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_readNodesDataToChannel(t *testing.T) {
	type Key = int
	type Data = int
	type Node = quorum.Node[Key, Data]
	type NodeMock = nodeWithId[Key, Data]
	ReadWithNodes := readNodesDataToChannelWithNodes[Key, Data]
	allInChannel := fn.Bind(cntx.WithTimeout(100*time.Millisecond), chn.AllWithCtx[Data])

	t.Run("should close channel after all nodes read", func(tt *testing.T) {
		ch := ReadWithNodes([]Node{NodeMock{}})
		assert.True(tt, allInChannel(ch).IsOk())
	})
}
