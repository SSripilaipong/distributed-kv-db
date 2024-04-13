package mergeread

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/tstexc"
	"distributed-kv-db/common/tstmck"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FromNodes(t *testing.T) {
	type Key = int
	type Data = int
	type Node = string
	Target := tstexc.NewFunc2Dep3In1Out(composeFromNodes[Key, Data, Node])
	Merge := tstmck.NewFunc2In1OutLike(Target.D1)
	QRead := tstmck.NewFunc3In1OutLike(Target.D2)
	Target = Target.
		DefaultDep1(Merge.ReturnO1(0)).
		DefaultDep2(QRead.ReturnO1(rslt.Value([]Data{})))
	WithMerge := Target.WithDep1
	WithQRead := Target.WithDep2
	WithContext := Target.WithIn1
	WithNodes := Target.WithIn2
	WithKey := Target.WithIn3
	QReadCaptureContext := QRead.CaptureI1
	QReadCaptureNodes := QRead.CaptureI2
	QReadCaptureKey := QRead.CaptureI3
	QReadWithResult := QRead.ReturnO1
	MergeCaptureXAndY := func(x *Data, y *Data) func(x Data, y Data) Data {
		return Merge.WithCaptureI1(x).WithCaptureI2(y).Build()
	}
	MergeCaptureAllXAndAllY := func(x *[]Data, y *[]Data) func(x Data, y Data) Data {
		return Merge.WithCaptureAllI1(x).WithCaptureAllI2(y).Build()
	}

	t.Run("should read nodes with context", func(t *testing.T) {
		var capturedContext context.Context
		ctx, isSame := cntx.WithVerifier()
		Target.Execute(
			WithContext(ctx),
			WithQRead(QReadCaptureContext(&capturedContext)),
		)
		assert.True(t, isSame(capturedContext))
	})

	t.Run("should read from nodes", func(t *testing.T) {
		var nodes []Node
		Target.Execute(
			WithNodes([]Node{"a", "b"}),
			WithQRead(QReadCaptureNodes(&nodes)),
		)
		assert.Equal(t, []Node{"a", "b"}, nodes)
	})

	t.Run("should read nodes with key", func(t *testing.T) {
		var key Key
		Target.Execute(
			WithKey(123),
			WithQRead(QReadCaptureKey(&key)),
		)
		assert.Equal(t, 123, key)
	})

	t.Run("should merge read result", func(t *testing.T) {
		var x, y Data
		Target.Execute(
			WithMerge(MergeCaptureXAndY(&x, &y)),
			WithQRead(QReadWithResult(rslt.Value([]Data{111, 222}))),
		)
		assert.Equal(t, 111, x)
		assert.Equal(t, 222, y)
	})

	t.Run("should reduce merge all read result", func(t *testing.T) {
		var xs, ys []Data
		Target.Execute(
			WithMerge(MergeCaptureAllXAndAllY(&xs, &ys)),
			WithQRead(QReadWithResult(rslt.Value([]Data{111, 222, 333}))),
		)
		assert.Equal(t, []Data{111, 0}, xs)
		assert.Equal(t, []Data{222, 333}, ys)
	})

	t.Run("should return merge result (*)", func(t *testing.T) {
		result := Target.Execute(
			WithMerge(func(x, y Data) Data { return x * y }),
			WithQRead(QReadWithResult(rslt.Value([]Data{2, 3, 4, 5}))),
		)

		assert.Equal(t, rslt.Value(120), result)
	})
}
