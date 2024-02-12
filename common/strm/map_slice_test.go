package strm

import (
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/slc"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_MapSlice(t *testing.T) {
	type A = string
	type B = int
	defaultTimeout := 100 * time.Millisecond
	allInChannel := fn.WithArg(cntx.WithTimeout(defaultTimeout), chn.AllWithCtx[B])
	waitUtilChannelClosed := rslt.OkFunc(fn.WithArg(cntx.WithTimeout(defaultTimeout), chn.AllWithCtx[B]))
	MapSliceWithXs := mapSliceWithXs[A, B]
	MapSliceWithFAndXs := mapSliceWithFAndXs[A, B]
	FCaptureAllX := mapSliceFCaptureAllX[A, B]
	FWithResults := mapSliceFWithAllResult[A, B]

	t.Run("should close channel after all x in xs read", func(tt *testing.T) {
		assert.True(tt, waitUtilChannelClosed(MapSliceWithXs([]A{"x"})))
	})

	t.Run("should read all x in xs with key", func(tt *testing.T) {
		var allX []A
		waitUtilChannelClosed(MapSliceWithFAndXs(
			FCaptureAllX(&allX), []A{"x1", "x2"},
		))
		assert.Equal(tt, []A{"x1", "x2"}, allX)
	})

	t.Run("should send all read data to channel", func(tt *testing.T) {
		ch := MapSliceWithFAndXs(FWithResults([]rslt.Of[B]{
			rslt.Value(123), rslt.Value(456),
		}), []A{"a", "b"})
		assert.Equal(tt, []B{123, 456}, slc.Sorted(allInChannel(ch).Value()))
	})
}
