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

type KeyNotFoundError2 struct {
	key string
}

var _ error = KeyNotFoundError2{}

func NewKeyNotFoundError2(key string) KeyNotFoundError2 {
	return KeyNotFoundError2{key: key}
}

func (e KeyNotFoundError2) Error() string {
	return fmt.Sprintf("key %#v not found", e.key)
}
