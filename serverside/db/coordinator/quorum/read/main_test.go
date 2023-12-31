package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/rslt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New(t *testing.T) {
	type Key = int
	type Data = int
	type Node = string
	Read := read[Key, Data]
	ReadWithContext := readWithContext[Key, Data]
	ReadWithKey := readWithKey[Key, Data]
	DiscoverNodesCaptureKey := discoverNodesFuncCaptureKey[Key, Node]
	DiscoverNodesCaptureContext := discoverNodesFuncCaptureContext[Key, Node]
	DiscoverNodesFuncWithResult := discoverNodesFuncWithResult[Key, Node]
	ReadNodesDataToChannelWithResult := readNodesDataToChannelWithResult[Key, Data, Node]
	NewFuncWithLatestData := newFuncWithLatestData[Key, Data, Node]
	NewFuncWithDiscoverNodes := newFuncWithDiscoverNodes[Key, Data, Node]
	ReadNodesDataToChannelCaptureContext := readNodesDataToChannelCaptureContext[Key, Data, Node]
	ReadNodesDataToChannelCaptureKey := readNodesDataToChannelCaptureKey[Key, Data, Node]
	ReadNodesDataToChannelCaptureNodes := readNodesDataToChannelCaptureNodes[Key, Data, Node]

	t.Run("should read nodes from discoverNodes with key", func(tt *testing.T) {
		var key Key
		ReadWithKey(
			NewFuncWithDiscoverNodes(DiscoverNodesCaptureKey(&key)),
			123,
		)
		assert.Equal(tt, 123, key)
	})

	t.Run("should read nodes from discoverNodes with context", func(tt *testing.T) {
		var ctx context.Context
		ReadWithContext(
			NewFuncWithDiscoverNodes(DiscoverNodesCaptureContext(&ctx)),
			cntx.WithValue("code name", "007"),
		)
		assert.Equal(tt, "007", ctx.Value("code name"))
	})

	t.Run("should call read nodes from discoverNodes to channel", func(tt *testing.T) {
		var nodes []Node
		Read(newFuncWithDiscoverNodesAndReadNodesDataToChannels(
			DiscoverNodesFuncWithResult(rslt.Value([]Node{"node1", "node2"})),
			ReadNodesDataToChannelCaptureNodes(&nodes),
		))
		assert.Equal(tt, []Node{"node1", "node2"}, nodes)
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
		Read(newFuncWithDiscoverNodesAndReadNodesDataToChannels(
			DiscoverNodesFuncWithResult(rslt.Value([]Node{"", "", ""})),
			ReadNodesDataToChannelWithResult(dataChan),
		))
		assert.Equal(tt, rslt.Value(13), chn.ReceiveNoWait(dataChan)) // remaining data
	})

	t.Run("should find latest data with a quorum of data", func(tt *testing.T) {
		var xs []Data
		Read(newFuncWithDiscoverNodesAndReadNodesDataToChannelsAndLatestData(
			DiscoverNodesFuncWithResult(rslt.Value([]Node{"", "", ""})),
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
