package tstexc

type Func1Dep4In1Out[D1, I1, I2, I3, I4, O1 any] struct {
	execute func(D1) func(I1, I2, I3, I4) O1
	D1      D1
	i1      I1
	i2      I2
	i3      I3
	i4      I4
	o1      O1
}

func NewFunc1Dep4In1Out[D1, I1, I2, I3, I4, O1 any](f func(D1) func(I1, I2, I3, I4) O1) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1] {
	return Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]{execute: f}
}

func (f Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) DefaultI1(x I1) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1] {
	g := f
	g.i1 = x
	return g
}

func (f Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) Execute(options ...func(Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) O1 {
	g := f
	for _, option := range options {
		g = option(g)
	}
	return g.execute(g.D1)(g.i1, g.i2, g.i3, g.i4)
}

func (Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) WithDep1(x D1) func(Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1] {
	return func(f Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1] {
		f.D1 = x
		return f
	}
}

func (Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) WithIn2(x I2) func(Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1] {
	return func(f Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1] {
		f.i2 = x
		return f
	}
}

func (Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) WithIn3(x I3) func(Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1] {
	return func(f Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1] {
		f.i3 = x
		return f
	}
}

func (Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) WithIn1(x I1) func(Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1] {
	return func(f Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1] {
		f.i1 = x
		return f
	}
}

func (Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) WithIn4(x I4) func(Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1] {
	return func(f Func1Dep4In1Out[D1, I1, I2, I3, I4, O1]) Func1Dep4In1Out[D1, I1, I2, I3, I4, O1] {
		f.i4 = x
		return f
	}
}
