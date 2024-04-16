package tstexc

// Code generated by tstexcgen DO NOT EDIT.

type Func1Dep1In1Out[D1, I1, O1 any] struct {
	execute func(D1) func(I1) O1
	D1      D1
	i1      I1
}

func NewFunc1Dep1In1Out[D1, I1, O1 any](f func(D1) func(I1) O1) Func1Dep1In1Out[D1, I1, O1] {
	return Func1Dep1In1Out[D1, I1, O1]{execute: f}
}

func (f Func1Dep1In1Out[D1, I1, O1]) Execute(options ...func(Func1Dep1In1Out[D1, I1, O1]) Func1Dep1In1Out[D1, I1, O1]) O1 {
	g := f
	for _, option := range options {
		g = option(g)
	}
	return g.execute(g.D1)(g.i1)
}

func (f Func1Dep1In1Out[D1, I1, O1]) DefaultDep1(x D1) Func1Dep1In1Out[D1, I1, O1] {
	g := f
	g.D1 = x
	return g
}

func (f Func1Dep1In1Out[D1, I1, O1]) DefaultIn1(x I1) Func1Dep1In1Out[D1, I1, O1] {
	g := f
	g.i1 = x
	return g
}

func (Func1Dep1In1Out[D1, I1, O1]) WithDep1(x D1) func(Func1Dep1In1Out[D1, I1, O1]) Func1Dep1In1Out[D1, I1, O1] {
	return func(f Func1Dep1In1Out[D1, I1, O1]) Func1Dep1In1Out[D1, I1, O1] {
		return f.DefaultDep1(x)
	}
}

func (Func1Dep1In1Out[D1, I1, O1]) WithIn1(x I1) func(Func1Dep1In1Out[D1, I1, O1]) Func1Dep1In1Out[D1, I1, O1] {
	return func(f Func1Dep1In1Out[D1, I1, O1]) Func1Dep1In1Out[D1, I1, O1] {
		return f.DefaultIn1(x)
	}
}