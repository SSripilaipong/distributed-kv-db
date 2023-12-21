package coordinator

import "distributed-kv-db/common/result"

func NewGetValue() GetValueFunc {
	return func(request GetValueRequest) result.Of[GetValueResponse] {
		return result.Value(GetValueResponse{Value: "tmp"})
	}
}

func NewSetValue() SetValueFunc {
	return func(request SetValueRequest) result.Of[SetValueResponse] {
		return result.Value(SetValueResponse{})
	}
}
