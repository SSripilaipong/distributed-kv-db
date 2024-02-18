package readlatest

import (
	"context"
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
	NewFuncWithLatestData := newFuncWithLatestData[Key, Data]
	ReadQuorumOfNodesDataWithResult := readQuorumWithResult[Key, Data]
	ReadQuorumCaptureContext := readQuorumCaptureContext[Key, Data]
	ReadQuorumCaptureKey := readQuorumCaptureKey[Key, Data]

	t.Run("should do quorum read with context", func(tt *testing.T) {
		var ctx context.Context
		ReadWithContext(newFuncWithReadQuorumOfNodesData(
			ReadQuorumCaptureContext(&ctx),
		), cntx.WithValue("foo", "bar"))
		assert.Equal(tt, "bar", ctx.Value("foo"))
	})

	t.Run("should do quorum read with key", func(tt *testing.T) {
		var key Key
		ReadWithKey(newFuncWithReadQuorumOfNodesData(
			ReadQuorumCaptureKey(&key),
		), 123)
		assert.Equal(tt, 123, key)
	})

	t.Run("should find latest data from quorum read", func(tt *testing.T) {
		var xs []Data
		Read(newFuncWithReadQuorumAndLatestData(
			ReadQuorumOfNodesDataWithResult(rslt.Value([]Data{123, 456})),
			latestDataCaptureXs(&xs),
		))
		assert.Equal(tt, []Data{123, 456}, xs)
	})

	t.Run("should return latest data", func(tt *testing.T) {
		result := Read(NewFuncWithLatestData(latestDataWithResult(555)))
		assert.Equal(tt, rslt.Value(555), result)
	})
}
