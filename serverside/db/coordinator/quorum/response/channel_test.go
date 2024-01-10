package response

import (
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/slc"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_channel(t *testing.T) {
	type Response = int
	type Node = string
	defaultTimeout := 100 * time.Millisecond
	allInChannel := fn.Bind(cntx.WithTimeout(defaultTimeout), chn.AllWithCtx[Response])
	waitUtilChannelClosed := rslt.OkFunc(fn.Bind(cntx.WithTimeout(defaultTimeout), chn.AllWithCtx[Response]))
	ChannelWithNodes := channelWithNodes[Response, Node]
	ChannelWithRequestAndNodes := channelWithRequestAndNodes[Response, Node]
	RequestCaptureAllNodes := requestCaptureAllNodes[Response, Node]
	RequestWithResults := requestWithResults[Response, Node]

	t.Run("should close channel after all nodes read", func(tt *testing.T) {
		assert.True(tt, waitUtilChannelClosed(ChannelWithNodes([]Node{"my node"})))
	})

	t.Run("should read all nodes with key", func(tt *testing.T) {
		var nodes []Node
		waitUtilChannelClosed(ChannelWithRequestAndNodes(
			RequestCaptureAllNodes(&nodes), []Node{"node1", "node2"},
		))
		assert.Equal(tt, []Node{"node1", "node2"}, nodes)
	})

	t.Run("should send all read data to channel", func(tt *testing.T) {
		ch := ChannelWithRequestAndNodes(RequestWithResults([]rslt.Of[Response]{
			rslt.Value(123), rslt.Value(456),
		}), []Node{"a", "b"})
		assert.Equal(tt, []Response{123, 456}, slc.Sorted(allInChannel(ch).Value()))
	})
}
