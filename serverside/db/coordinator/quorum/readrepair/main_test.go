package readrepair

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New(t *testing.T) {

	t.Run("should quorum read with key", func(tt *testing.T) {
		var key int
		readRepairWithKey(newFuncWithReadRepair(readQuorumCaptureKey[int, int](&key)), 123)
		assert.Equal(tt, 123, key)
	})

}
