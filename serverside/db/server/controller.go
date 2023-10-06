package server

type controller struct {
}

func (c controller) Error() <-chan error {
	return nil
}

func (c controller) ForceStop() {
}
