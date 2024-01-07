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
	ReadNodesDataToChannelWithResult := readNodesDataToChannelWithResult[Key, Data]
	NewFuncWithLatestData := newFuncWithLatestData[Key, Data]
	ReadNodesDataToChannelCaptureContext := readNodesDataToChannelCaptureContext[Key, Data]
	ReadNodesDataToChannelCaptureKey := readNodesDataToChannelCaptureKey[Key, Data]

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
		Read(newFuncWithDiscoveryAndReadNodesDataToChannels(
			discoveryWithNodesFunc(nodesFuncWithResult(rslt.Value([]Node{NodeMock{1}, NodeMock{2}}))),
			readNodesDataToChannelCaptureNodes(&nodes),
		))
		assert.Equal(tt, []Node{NodeMock{1}, NodeMock{2}}, nodes)
	})

	t.Run("should read nodes data with context", func(tt *testing.T) {
		var ctx context.Context
		ReadWithContext(newFuncWithReadNodesDataToChannels(
			ReadNodesDataToChannelCaptureContext(&ctx),
		), cntx.WithValue("foo", "bar"))
		assert.Equal(tt, "bar", ctx.Value("foo"))
	})

	t.Run("should read nodes data with key", func(tt *testing.T) {
		var key Key
		ReadWithKey(newFuncWithReadNodesDataToChannels(
			ReadNodesDataToChannelCaptureKey(&key),
		), 123)
		assert.Equal(tt, 123, key)
	})

	t.Run("should call read only a quorum of nodes from channels", func(tt *testing.T) {
		dataChan := chn.NewFromSlice([]Data{11, 12, 13})
		Read(newFuncWithDiscoveryAndReadNodesDataToChannels(
			discoveryWithNodesFunc(nodesFuncWithResult(rslt.Value([]Node{NodeMock{}, NodeMock{}, NodeMock{}}))),
			ReadNodesDataToChannelWithResult(dataChan),
		))
		assert.Equal(tt, rslt.Value(13), chn.ReceiveNoWait(dataChan)) // remaining data
	})

	t.Run("should find latest data with a quorum of data", func(tt *testing.T) {
		var xs []Data
		Read(newFuncWithDiscoveryAndReadNodesDataToChannelsAndLatestData(
			discoveryWithNodesFunc(nodesFuncWithResult(rslt.Value([]Node{NodeMock{}, NodeMock{}, NodeMock{}}))),
			ReadNodesDataToChannelWithResult(chn.NewFromSlice([]Data{123, 456, 0})),
			latestDataCaptureXs(&xs),
		))
		assert.Equal(tt, []Data{123, 456}, xs)
	})

	t.Run("should return latest data", func(tt *testing.T) {
		result := Read(NewFuncWithLatestData(latestDataWithResult(555)))
		assert.Equal(tt, rslt.Value(555), result)
	})
}
