package coordinator

import "fmt"

type KeyNotFoundError struct {
	key string
}

var _ error = KeyNotFoundError{}

func NewKeyNotFoundError(key string) KeyNotFoundError {
	return KeyNotFoundError{key: key}
}

func (e KeyNotFoundError) Error() string {
	return fmt.Sprintf("key %#v not found", e.key)
}
