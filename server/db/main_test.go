package db

import (
	"distributed-kv-db/server/db/server"
	"errors"
	"testing"
)

import "github.com/stretchr/testify/assert"

func Test_build(t *testing.T) {
	t.Run("should run serverFunc with dbPort", func(tt *testing.T) {
		var calledPort int
		var serverFunc server.Func = func(port int) error {
			calledPort = port
			return nil
		}
		f := build(serverFunc)

		_ = f(1234, 5678, "", []string{})

		assert.Equal(tt, 1234, calledPort)
	})

	t.Run("should return no error", func(tt *testing.T) {
		var serverFunc server.Func = func(port int) error {
			return nil
		}
		f := build(serverFunc)

		err := f(1234, 5678, "", []string{})

		assert.Nil(tt, err)
	})

	t.Run("should return error if serverFunc returns one", func(tt *testing.T) {
		var serverFunc server.Func = func(port int) error {
			return errors.New("boom")
		}
		f := build(serverFunc)

		err := f(1234, 5678, "", []string{})

		assert.Equal(tt, errors.New("boom"), err)
	})
}
