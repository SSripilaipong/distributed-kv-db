package temporal

import (
	"distributed-kv-db/common/rslt"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LatestInSlice(t *testing.T) {
	type Data = hashableDataMock
	latestInSlice := LatestInSlice[Data]

	t.Run("should return error if slice is empty", func(tt *testing.T) {
		result := latestInSlice([]Data{})
		assert.Equal(tt, rslt.Error[Data](errors.New("no data")), result)
	})

	t.Run("should return error if slice is nil", func(tt *testing.T) {
		result := latestInSlice(nil)
		assert.Equal(tt, rslt.Error[Data](errors.New("no data")), result)
	})

	t.Run("should return newest data", func(tt *testing.T) {
		result := latestInSlice([]Data{
			orderableDataWithNewness(1),
			orderableDataWithNewness(2),
			orderableDataWithNewness(0),
		})
		assert.Equal(tt, rslt.Value(orderableDataWithNewness(2)), result)
	})

	t.Run("should return one with higher hash if newness is the same", func(tt *testing.T) {
		result := latestInSlice([]Data{
			orderableDataWithNewnessAndHash(1, 999),
			orderableDataWithNewnessAndHash(2, 111),
			orderableDataWithNewnessAndHash(2, 222),
		})
		assert.Equal(tt, rslt.Value(orderableDataWithNewnessAndHash(2, 222)), result)
	})
}
