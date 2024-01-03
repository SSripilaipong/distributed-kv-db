package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New(t *testing.T) {
	type Node = quorum.Node[int, int]
	type NodeMock = nodeWithId[int, int]
	Read := read[int, int]
	ReadWithContext := readWithContext[int, int]
	DiscoveryCaptureCtx := discoveryCaptureCtx[int, int]

	t.Run("should read nodes from discovery with context", func(tt *testing.T) {
		var ctx context.Context
		ReadWithContext(
			newFuncWithDiscovery(DiscoveryCaptureCtx(&ctx)),
			cntx.WithValue("code name", "007"),
		)
		assert.Equal(tt, "007", ctx.Value("code name"))
	})

	t.Run("should call read nodes from discovery to channel", func(tt *testing.T) {
		var nodes []Node
		Read(newFuncWithDiscoveryAndReadNodesToChannels(
			discoveryWithNodes(rslt.Value([]Node{NodeMock{1}, NodeMock{2}})),
			readNodesToChannelCaptureNodes(&nodes),
		))
		assert.Equal(tt, []Node{NodeMock{1}, NodeMock{2}}, nodes)
	})

	t.Run("should call read only a quorum of nodes from channels", func(tt *testing.T) {
		dataChan := chn.NewFromSlice([]int{11, 12, 13})
		Read(newFuncWithDiscoveryAndReadNodesToChannels(
			discoveryWithNodes(rslt.Value([]Node{NodeMock{}, NodeMock{}, NodeMock{}})),
			readNodesToChannelWithResult[int](dataChan),
		))
		assert.Equal(tt, rslt.Value(13), chn.ReceiveNoWait(dataChan)) // remaining data
	})

}
