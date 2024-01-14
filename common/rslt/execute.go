package rslt

func MapOfFunc[A, B any](f Of[func(A) B]) func(A) Of[B] {
	if f.IsOk() {
		return func(x A) Of[B] {
			return Value(f.Value()(x))
		}
	}

	return func(x A) Of[B] {
		return Error[B](f.Error())
	}
}

func MapOfFuncPartial[A, B any](f Of[func(A) Of[B]]) func(A) Of[B] {
	if f.IsOk() {
		return func(x A) Of[B] {
			return f.Value()(x)
		}
	}

	return func(x A) Of[B] {
		return Error[B](f.Error())
	}
}
