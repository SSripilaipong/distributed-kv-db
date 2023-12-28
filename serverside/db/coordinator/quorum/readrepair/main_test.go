package readrepair

import (
	"distributed-kv-db/common/rslt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New(t *testing.T) {

	t.Run("should quorum read with key", func(tt *testing.T) {
		var key int
		readRepairWithKey(newFuncWithQuorumRead(readQuorumCaptureKey[int, int](&key)), 123)
		assert.Equal(tt, 123, key)
	})

	t.Run("should quorum write with key and result from quorum read", func(tt *testing.T) {
		var key, data int

		readRepairWithKey(newFuncWithQuorumReadAndQuorumWrite(
			readQuorumWithResult[int, int](rslt.Value(555)),
			writeQuorumCaptureKeyAndData[int, int](&key, &data),
		), 456)

		assert.Equal(tt, 456, key)
		assert.Equal(tt, 555, data)
	})

}
