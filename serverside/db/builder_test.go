package db

import (
	"distributed-kv-db/serverside/db/contract"
	"distributed-kv-db/serverside/db/server"
	"errors"
	"testing"
)

import "github.com/stretchr/testify/assert"

func Test_build(t *testing.T) {
	t.Run("should run serverFunc with dbPort", func(t *testing.T) {
		var calledPort int
		var serverFunc server.Func = func(port int) contract.Controller {
			calledPort = port
			return &ControllerMock{}
		}

		_ = execute(buildInterruptedWithServerFunc(serverFunc))

		assert.Equal(t, 1234, calledPort)
	})

	t.Run("should return no error", func(t *testing.T) {
		var serverFunc server.Func = func(port int) contract.Controller {
			return &ControllerMock{}
		}

		err := execute(buildInterruptedWithServerFunc(serverFunc))

		assert.Nil(t, err)
	})

	t.Run("should force stop server if interrupted", func(t *testing.T) {
		serverCtrl := &ControllerMock{}
		var serverFunc server.Func = func(port int) contract.Controller {
			return serverCtrl
		}

		_ = execute(buildInterruptedWithServerFunc(serverFunc))

		assert.True(t, serverCtrl.ForceStop_IsCalled)
	})

	t.Run("should return error if serverFunc returns one", func(t *testing.T) {
		var serverFunc server.Func = func(port int) contract.Controller {
			return &ControllerMock{err: errors.New("boom")}
		}

		err := execute(buildWithServerFunc(serverFunc))

		assert.Equal(t, errors.New("boom"), err)
	})
}

func buildInterruptedWithServerFunc(serverFunc server.Func) Func {
	return builder(interruptedSignal, serverFunc)
}

func buildWithServerFunc(serverFunc server.Func) Func {
	return builder(uninterruptedSignal, serverFunc)
}

func execute(f Func) error {
	return f(1234, 5678, "", []string{})
}
