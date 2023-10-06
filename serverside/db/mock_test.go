package db

import "distributed-kv-db/serverside/db/contract"

type ControllerMock struct {
	err                error
	ForceStop_IsCalled bool
}

func (c *ControllerMock) Error() <-chan error {
	ch := make(chan error, 1)
	if c.err != nil {
		ch <- c.err
	}
	return ch
}

func (c *ControllerMock) ForceStop() {
	c.ForceStop_IsCalled = true
}

var _ contract.Controller = &ControllerMock{}
