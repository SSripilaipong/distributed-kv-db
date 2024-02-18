package temporal

type DataWithHashComparison[D Hashable] struct {
	value D
}

func NewWithHashComparison[D Hashable](data D) DataWithHashComparison[D] {
	return DataWithHashComparison[D]{
		value: data,
	}
}

func (x DataWithHashComparison[D]) Compare(y DataWithHashComparison[D]) int {
	vx, vy := x.value, y.value
	if newnessDiff := vx.Newness() - vy.Newness(); newnessDiff != 0 {
		return newnessDiff
	}
	return vx.Hash() - vy.Hash()
}

func (x DataWithHashComparison[D]) Value() D {
	return x.value
}
