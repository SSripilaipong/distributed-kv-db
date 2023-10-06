package contract

type Controller interface {
	Error() <-chan error
	ForceStop()
}
