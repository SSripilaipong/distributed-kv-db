package readrepair

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/tstexc"
	"distributed-kv-db/common/tstmck"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FromNodes(t *testing.T) {
	type Key = int
	type Data = int
	type Node = string
	e := tstexc.NewFunc2Dep3In1Out(composeFromNodes[Key, Data, Node])
	qMergeRead := tstmck.NewFunc3In1OutLike(e.D1)
	qBlindWrite := tstmck.NewFunc4In1OutLike(e.D2)
	e = e.
		DefaultDep1(qMergeRead.ReturnO1(rslt.Value(0))).
		DefaultDep2(qBlindWrite.ReturnO1(nil))
	WithQMergeRead := e.WithDep1
	WithQBlindWrite := e.WithDep2
	WithContext := e.WithIn1
	WithNodes := e.WithIn2
	WithKey := e.WithIn3
	QMergeReadCaptureContext := qMergeRead.CaptureI1
	QMergeReadCaptureNodes := qMergeRead.CaptureI2
	QMergeReadCaptureKey := qMergeRead.CaptureI3
	QMergeReadWithResult := qMergeRead.ReturnO1
	QBlindWriteCaptureContext := qBlindWrite.CaptureI1
	QBlindWriteCaptureNodes := qBlindWrite.CaptureI2
	QBlindWriteCaptureKey := qBlindWrite.CaptureI3
	QBlindWriteCaptureData := qBlindWrite.CaptureI4
	QBlindWriteCheckIsCalled := qBlindWrite.CheckIsCalled
	QBlindWriteCheckWithResult := qBlindWrite.ReturnO1

	t.Run("should read nodes", func(t *testing.T) {
		var capturedNodes []Node
		e.Execute(
			WithNodes([]Node{"aaa", "bbb"}),
			WithQMergeRead(QMergeReadCaptureNodes(&capturedNodes)),
		)
		assert.ElementsMatch(t, []Node{"aaa", "bbb"}, capturedNodes)
	})

	t.Run("should read with same context", func(t *testing.T) {
		var capturedCtx context.Context
		ctx, isSame := cntx.WithVerifier()
		e.Execute(
			WithContext(ctx),
			WithQMergeRead(QMergeReadCaptureContext(&capturedCtx)),
		)
		assert.True(t, isSame(capturedCtx))
	})

	t.Run("should read with key", func(t *testing.T) {
		var capturedKey Key
		e.Execute(
			WithKey(999),
			WithQMergeRead(QMergeReadCaptureKey(&capturedKey)),
		)
		assert.Equal(t, 999, capturedKey)
	})

	t.Run("should write with same context", func(t *testing.T) {
		var capturedContext context.Context
		ctx, isSame := cntx.WithVerifier()
		e.Execute(
			WithContext(ctx),
			WithQBlindWrite(QBlindWriteCaptureContext(&capturedContext)),
		)
		assert.True(t, isSame(capturedContext))
	})

	t.Run("should write with nodes", func(t *testing.T) {
		var capturedNodes []Node
		e.Execute(
			WithNodes([]Node{"a", "b"}),
			WithQBlindWrite(QBlindWriteCaptureNodes(&capturedNodes)),
		)
		assert.Equal(t, []Node{"a", "b"}, capturedNodes)
	})

	t.Run("should write with key", func(t *testing.T) {
		var capturedKey Key
		e.Execute(
			WithKey(1234),
			WithQBlindWrite(QBlindWriteCaptureKey(&capturedKey)),
		)
		assert.Equal(t, 1234, capturedKey)
	})

	t.Run("should write with read data", func(t *testing.T) {
		var capturedData Data
		e.Execute(
			WithQMergeRead(QMergeReadWithResult(rslt.Value(555))),
			WithQBlindWrite(QBlindWriteCaptureData(&capturedData)),
		)
		assert.Equal(t, 555, capturedData)
	})

	t.Run("should not write when read fails", func(t *testing.T) {
		var isCalled bool
		e.Execute(
			WithQMergeRead(QMergeReadWithResult(rslt.Error[Data](errors.New("yuck")))),
			WithQBlindWrite(QBlindWriteCheckIsCalled(&isCalled)),
		)
		assert.False(t, isCalled)
	})

	t.Run("should return read result", func(t *testing.T) {
		result := e.Execute(
			WithQMergeRead(QMergeReadWithResult(rslt.Value(999))),
		)

		assert.Equal(t, rslt.Value(999), result)
	})

	t.Run("should return error from write if write fails", func(t *testing.T) {
		result := e.Execute(
			WithQBlindWrite(QBlindWriteCheckWithResult(errors.New("fail ja"))),
		)

		assert.Equal(t, rslt.Error[Data](errors.New("fail ja")), result)
	})
}
