package fnopts

func Apply[Data any](d Data, options []func(*Data)) Data {
	for _, option := range options {
		option(&d)
	}
	return d
}
