package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/slc"
	"distributed-kv-db/common/typ"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_readNodesDataToChannel(t *testing.T) {
	type Key = int
	type Data = int
	type Node = ReadableNode[Key, Data]
	defaultTimeout := 100 * time.Millisecond
	NodeMock := nodeDummy[Key, Data]
	ReadWithNodes := readNodesDataToChannelWithNodes[Key, Data, Node]
	allInChannel := fn.WithArg(cntx.WithTimeout(defaultTimeout), chn.AllWithCtx[Data])
	waitUtilChannelClosed := rslt.OkFunc(fn.WithArg(cntx.WithTimeout(defaultTimeout), chn.AllWithCtx[Data]))
	ReadWithContextAndNodes := readNodesDataToChannelWithContextAndNodes[Key, Data, Node]
	ReadWithKeyAndNodes := readWithKeyAndNodes[Key, Data, Node]
	NodeWithReadFunc := nodeWithReadFunc[Key, Data]
	ReadFuncCaptureContext := readFuncCaptureContext[Key, Data]
	ReadFuncCaptureKey := readFuncCaptureKey[Key, Data]
	ReadFuncWithResult := readFuncWithResult[Key, Data]

	t.Run("should close channel after all nodes read", func(tt *testing.T) {
		assert.True(tt, waitUtilChannelClosed(ReadWithNodes([]Node{NodeMock()})))
	})

	t.Run("should read all nodes with context", func(tt *testing.T) {
		var ctx1, ctx2 context.Context
		waitUtilChannelClosed(ReadWithContextAndNodes(
			cntx.WithValue("aaa", "bbb"),
			[]Node{
				NodeWithReadFunc(ReadFuncCaptureContext(&ctx1)),
				NodeWithReadFunc(ReadFuncCaptureContext(&ctx2)),
			},
		))
		assert.Equal(tt, "bbb", ctx1.Value("aaa"))
		assert.Equal(tt, "bbb", ctx2.Value("aaa"))
	})

	t.Run("should read all nodes with key", func(tt *testing.T) {
		var key1, key2 Key
		waitUtilChannelClosed(ReadWithKeyAndNodes(
			555,
			[]Node{
				NodeWithReadFunc(ReadFuncCaptureKey(&key1)),
				NodeWithReadFunc(ReadFuncCaptureKey(&key2)),
			},
		))
		assert.Equal(tt, 555, key1)
		assert.Equal(tt, 555, key2)
	})

	t.Run("should send all read data to channel", func(tt *testing.T) {
		ch := ReadWithNodes([]Node{
			NodeWithReadFunc(ReadFuncWithResult(rslt.Value(123))),
			NodeWithReadFunc(ReadFuncWithResult(rslt.Value(456))),
		})
		assert.Equal(tt, []Data{123, 456}, slc.Sorted(allInChannel(ch).Value()))
	})
}

func readNodesDataToChannelWithNodes[Key, Data any, Node ReadableNode[Key, Data]](nodes []Node) <-chan Data {
	return readNodesDataToChannelWithContextAndNodes[Key, Data](context.Background(), nodes)
}

func readNodesDataToChannelWithContextAndNodes[Key, Data any, Node ReadableNode[Key, Data]](ctx context.Context, nodes []Node) <-chan Data {
	return NodesDataToChannel[Key, Data](ctx, typ.Zero[Key](), nodes)
}

func readWithKeyAndNodes[Key, Data any, Node ReadableNode[Key, Data]](key Key, nodes []Node) <-chan Data {
	return NodesDataToChannel[Key, Data](context.Background(), key, nodes)
}
