package request

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

func Test_NodesToChannel(t *testing.T) {
	type Node = string
	type Data = int
	defaultTimeout := 100 * time.Millisecond
	NewTarget := nodesToChannelForTest[Node, Data]
	Execute := runNodesToChannelForTest[Node, Data]
	WithRequest := withRequest[Node, Data]
	WithNodes := withNodes[Node]
	waitUtilChannelClosed := rslt.OkFunc(fn.Bind(cntx.WithTimeout(defaultTimeout), chn.AllWithCtx[Data]))
	allInChannel := fn.Bind(cntx.WithTimeout(defaultTimeout), chn.AllWithCtx[Data])

	t.Run("should return channel", func(tt *testing.T) {
		assert.NotNil(tt, Execute(NewTarget()))
	})

	t.Run("should apply request to all nodes", func(tt *testing.T) {
		var allNode []Node
		waitUtilChannelClosed(Execute(
			NewTarget(WithRequest(requestCaptureAllNode[Node, Data](&allNode))),
			WithNodes([]Node{"node1", "node2"}),
		))
		assert.Equal(tt, []Node{"node1", "node2"}, slc.Sorted(allNode))
	})

	t.Run("should return all result", func(tt *testing.T) {
		result := allInChannel(Execute(
			NewTarget(WithRequest(requestWithAllResult[Node, Data]([]Data{2, 1}))),
			WithNodes([]Node{"node1", "node2"}),
		))
		assert.Equal(tt, []Data{1, 2}, slc.Sorted(result.Value()))
	})
}
