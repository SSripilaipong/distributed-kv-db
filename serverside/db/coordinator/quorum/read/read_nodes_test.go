package read

import (
	"context"
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
	defaultTimeout := 100 * time.Millisecond
	NodeMock := nodeDummy[Key, Data]
	ReadWithNodes := readNodesDataToChannelWithNodes[Key, Data]
	allInChannel := fn.Bind(cntx.WithTimeout(defaultTimeout), chn.AllWithCtx[Data])
	ReadWithContextAndNodes := readNodesDataToChannelWithContextAndNodes[Key, Data]
	NodeWithReadFunc := nodeWithReadFunc[Key, Data]
	ReadFuncCaptureContext := readFuncCaptureContext[Key, Data]

	t.Run("should close channel after all nodes read", func(tt *testing.T) {
		ch := ReadWithNodes([]Node{NodeMock()})
		assert.True(tt, allInChannel(ch).IsOk())
	})

	t.Run("should read all nodes with context", func(tt *testing.T) {
		var ctx1, ctx2 context.Context
		ReadWithContextAndNodes(
			cntx.WithValue("aaa", "bbb"),
			[]Node{
				NodeWithReadFunc(ReadFuncCaptureContext(&ctx1)),
				NodeWithReadFunc(ReadFuncCaptureContext(&ctx2)),
			},
		)
		assert.Equal(tt, "bbb", ctx1.Value("aaa"))
		assert.Equal(tt, "bbb", ctx2.Value("aaa"))
	})
}
