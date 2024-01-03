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
	type Key = int
	type Data = int
	type Node = quorum.Node[Key, Data]
	type NodeMock = nodeWithId[Key, Data]
	Read := read[Key, Data]
	ReadWithContext := readWithContext[Key, Data]
	ReadWithKey := readWithKey[Key, Data]
	NodesFuncCaptureKey := nodesFuncCaptureKey[Key, Data]
	NodesFuncCaptureContext := nodesFuncCaptureContext[Key, Data]
	ReadNodesToChannelWithResult := readNodesToChannelWithResult[Key, Data]

	t.Run("should read nodes from discovery with key", func(tt *testing.T) {
		var key Key
		ReadWithKey(
			newFuncWithDiscovery(discoveryWithNodesFunc(NodesFuncCaptureKey(&key))),
			123,
		)
		assert.Equal(tt, 123, key)
	})

	t.Run("should read nodes from discovery with context", func(tt *testing.T) {
		var ctx context.Context
		ReadWithContext(
			newFuncWithDiscovery(discoveryWithNodesFunc(NodesFuncCaptureContext(&ctx))),
			cntx.WithValue("code name", "007"),
		)
		assert.Equal(tt, "007", ctx.Value("code name"))
	})

	t.Run("should call read nodes from discovery to channel", func(tt *testing.T) {
		var nodes []Node
		Read(newFuncWithDiscoveryAndReadNodesToChannels(
			discoveryWithNodesFunc(nodesFuncWithResult(rslt.Value([]Node{NodeMock{1}, NodeMock{2}}))),
			readNodesToChannelCaptureNodes(&nodes),
		))
		assert.Equal(tt, []Node{NodeMock{1}, NodeMock{2}}, nodes)
	})

	t.Run("should call read only a quorum of nodes from channels", func(tt *testing.T) {
		dataChan := chn.NewFromSlice([]Key{11, 12, 13})
		Read(newFuncWithDiscoveryAndReadNodesToChannels(
			discoveryWithNodesFunc(nodesFuncWithResult(rslt.Value([]Node{NodeMock{}, NodeMock{}, NodeMock{}}))),
			ReadNodesToChannelWithResult(dataChan),
		))
		assert.Equal(tt, rslt.Value(13), chn.ReceiveNoWait(dataChan)) // remaining data
	})

}
