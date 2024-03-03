package tstmck

type Func4In1Out[I1, I2, I3, I4, O1 any] struct {
	o1     func() O1
	i1Hook func(t I1)
	i2Hook func(t I2)
	i3Hook func(t I3)
	i4Hook func(t I4)
}

func NewFunc4In1OutLike[I1, I2, I3, I4, O1 any](_ func(I1, I2, I3, I4) O1) Func4In1Out[I1, I2, I3, I4, O1] {
	return Func4In1Out[I1, I2, I3, I4, O1]{
		i1Hook: func(t I1) {},
		i2Hook: func(t I2) {},
		i3Hook: func(t I3) {},
		i4Hook: func(t I4) {},
		o1:     func() (_ O1) { return },
	}
}

func (f Func4In1Out[I1, I2, I3, I4, O1]) CaptureAllI1(x *[]I1) func(I1, I2, I3, I4) O1 {
	f.i1Hook = func(t I1) {
		*x = append(*x, t)
	}
	return f.Build()
}

func (f Func4In1Out[I1, I2, I3, I4, O1]) CaptureAllI2(x *[]I2) func(I1, I2, I3, I4) O1 {
	f.i2Hook = func(t I2) {
		*x = append(*x, t)
	}
	return f.Build()
}

func (f Func4In1Out[I1, I2, I3, I4, O1]) CaptureAllI3(x *[]I3) func(I1, I2, I3, I4) O1 {
	f.i3Hook = func(t I3) {
		*x = append(*x, t)
	}
	return f.Build()
}

func (f Func4In1Out[I1, I2, I3, I4, O1]) CaptureAllI4(x *[]I4) func(I1, I2, I3, I4) O1 {
	f.i4Hook = func(t I4) {
		*x = append(*x, t)
	}
	return f.Build()
}

func (f Func4In1Out[I1, I2, I3, I4, O1]) ReturnAllO1(xs []O1) func(I1, I2, I3, I4) O1 {
	results := append([]O1{}, xs...)

	f.o1 = func() O1 {
		x := results[0]
		results = results[1:]
		return x
	}
	return f.Build()
}

func (f Func4In1Out[I1, I2, I3, I4, O1]) Build() func(I1, I2, I3, I4) O1 {
	return func(i1 I1, i2 I2, i3 I3, i4 I4) O1 {
		f.i1Hook(i1)
		f.i2Hook(i2)
		f.i3Hook(i3)
		f.i4Hook(i4)
		return f.o1()
	}
}
