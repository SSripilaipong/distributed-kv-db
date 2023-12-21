package result

func Fmap[A, B any](f func(A) B) func(Of[A]) Of[B] {
	return func(x Of[A]) Of[B] {
		if x.IsOk() {
			return Value(f(x.Value()))
		}
		return Error[B](x.Error())
	}
}

func FmapPartial[A, B any](f func(A) Of[B]) func(Of[A]) Of[B] {
	return func(x Of[A]) Of[B] {
		if x.IsOk() {
			return f(x.Value())
		}
		return Error[B](x.Error())
	}
}

func FmapError[A any](f func(A) error) func(Of[A]) error {
	return func(x Of[A]) error {
		if x.IsError() {
			return x.Error()
		}
		if err := f(x.Value()); err != nil {
			return err
		}
		return nil
	}
}
