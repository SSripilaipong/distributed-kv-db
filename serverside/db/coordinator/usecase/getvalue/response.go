package getvalue

func responseFromValue(value string) Response {
	return Response{Value: value}
}

type Response struct {
	Value string
}
