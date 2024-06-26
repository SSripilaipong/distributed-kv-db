package tstexc

// Code generated by tstexcgen DO NOT EDIT.

type Func2Dep3In1Out[D1, D2, I1, I2, I3, O1 any] struct {
	execute func(D1, D2) func(I1, I2, I3) O1
	D1      D1
	D2      D2
	i1      I1
	i2      I2
	i3      I3
}

func NewFunc2Dep3In1Out[D1, D2, I1, I2, I3, O1 any](f func(D1, D2) func(I1, I2, I3) O1) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
	return Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]{execute: f}
}

func (f Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) Execute(options ...func(Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) O1 {
	g := f
	for _, option := range options {
		g = option(g)
	}
	return g.execute(g.D1, g.D2)(g.i1, g.i2, g.i3)
}

func (f Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) DefaultDep1(x D1) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
	g := f
	g.D1 = x
	return g
}

func (f Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) DefaultDep2(x D2) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
	g := f
	g.D2 = x
	return g
}

func (f Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) DefaultIn1(x I1) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
	g := f
	g.i1 = x
	return g
}

func (f Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) DefaultIn2(x I2) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
	g := f
	g.i2 = x
	return g
}

func (f Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) DefaultIn3(x I3) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
	g := f
	g.i3 = x
	return g
}

func (Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) WithDep1(x D1) func(Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
	return func(f Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
		return f.DefaultDep1(x)
	}
}

func (Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) WithDep2(x D2) func(Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
	return func(f Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
		return f.DefaultDep2(x)
	}
}

func (Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) WithIn1(x I1) func(Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
	return func(f Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
		return f.DefaultIn1(x)
	}
}

func (Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) WithIn2(x I2) func(Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
	return func(f Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
		return f.DefaultIn2(x)
	}
}

func (Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) WithIn3(x I3) func(Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
	return func(f Func2Dep3In1Out[D1, D2, I1, I2, I3, O1]) Func2Dep3In1Out[D1, D2, I1, I2, I3, O1] {
		return f.DefaultIn3(x)
	}
}
