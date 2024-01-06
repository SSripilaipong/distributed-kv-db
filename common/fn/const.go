package fn

func Const[A, B any](v B) func(A) B {
	return func(A) B {
		return v
	}
}

func Const2[A1, A2, B any](v B) func(A1, A2) B {
	return func(A1, A2) B {
		return v
	}
}
