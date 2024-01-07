package read

import (
	"distributed-kv-db/common/rslt"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_latestData(t *testing.T) {
	type Data = orderableDataMock
	LatestData := latestData[Data]
	t.Run("should return error if slice is empty", func(tt *testing.T) {
		result := LatestData([]Data{})
		assert.Equal(tt, rslt.Error[Data](errors.New("no data")), result)
	})

	t.Run("should return error if slice is nil", func(tt *testing.T) {
		result := LatestData(nil)
		assert.Equal(tt, rslt.Error[Data](errors.New("no data")), result)
	})

	t.Run("should return newest data", func(tt *testing.T) {
		result := LatestData([]Data{
			orderableDataWithNewness(1),
			orderableDataWithNewness(2),
			orderableDataWithNewness(0),
		})
		assert.Equal(tt, rslt.Value(orderableDataWithNewness(2)), result)
	})
}
