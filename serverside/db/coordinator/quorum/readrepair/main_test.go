package readrepair

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/rslt"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New(t *testing.T) {

	t.Run("should quorum read with key", func(tt *testing.T) {
		var key int
		readRepairWithKey(newFuncWithQuorumRead(quorumReadCaptureKey[int, int](&key)), 123)
		assert.Equal(tt, 123, key)
	})

	t.Run("should quorum read with context", func(tt *testing.T) {
		var ctx context.Context
		readRepairWithContext(
			newFuncWithQuorumRead(quorumReadCaptureContext[int, int](&ctx)),
			cntx.WithValue("this is", "the same"),
		)
		assert.Equal(tt, "the same", ctx.Value("this is"))
	})

	t.Run("should quorum write with key", func(tt *testing.T) {
		var key int
		readRepairWithKey(newFuncWithQuorumWrite(quorumWriteCaptureKey[int, int](&key)), 456)
		assert.Equal(tt, 456, key)
	})

	t.Run("should quorum write with result from quorum read", func(tt *testing.T) {
		var data int
		readRepair(newFuncWithQuorumReadAndQuorumWrite(
			quorumReadWithResult[int, int](rslt.Value(555)),
			quorumWriteCaptureData[int, int](&data),
		))
		assert.Equal(tt, 555, data)
	})

	t.Run("should quorum write with context", func(tt *testing.T) {
		var ctx context.Context
		readRepairWithContext(newFuncWithQuorumWrite(
			quorumWriteCaptureContext[int, int](&ctx)),
			cntx.WithValue("my", "context ja"),
		)
		assert.Equal(tt, "context ja", ctx.Value("my"))
	})

	t.Run("should not quorum write if quorum read fails", func(tt *testing.T) {
		var isCalled bool
		readRepair(newFuncWithQuorumReadAndQuorumWrite(
			quorumReadWithResult[int, int](rslt.Error[int](errors.New(""))),
			quorumWriteCaptureIsCalled[int, int](&isCalled),
		))
		assert.False(tt, isCalled)
	})
}
