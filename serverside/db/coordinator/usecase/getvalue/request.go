package getvalue

func keyOfRequest(request Request) string {
	return request.Key
}

type Request struct {
	Key string
}
