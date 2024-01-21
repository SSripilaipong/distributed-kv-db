package error

import "fmt"

type KeyNotFound struct {
	key string
}

var _ error = KeyNotFound{}

func NewKeyNotFound(key string) KeyNotFound {
	return KeyNotFound{key: key}
}

func (e KeyNotFound) Error() string {
	return fmt.Sprintf("key %#v not found", e.key)
}
