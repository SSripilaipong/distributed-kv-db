package fn

func Const[A, B any](v B) func(A) B {
	return func(A) B {
		return v
	}
}
