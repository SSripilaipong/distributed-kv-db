package filter

import (
	"errors"
)

func ErrorChannel(n uint) func(<-chan error) error {
	quorum := nToQuorum(int(n))

	return func(ch <-chan error) error {
		y := newSuccessOrErrorCollection()

		for x := range ch {
			y = y.Collect(x)

			if y.ErrCount() >= quorum {
				return y.Error()
			}
			if y.SuccessCount() >= quorum {
				return nil
			}
		}

		return errors.Join(y.Error(), errors.New("channel closed"))
	}
}

type successOrErrorCollection struct {
	successCount int
	errCount     int
	err          error
}

func newSuccessOrErrorCollection() successOrErrorCollection {
	return successOrErrorCollection{}
}

func (c successOrErrorCollection) Collect(err error) successOrErrorCollection {
	if err != nil {
		return c.AddError(err)
	}
	return c.IncreaseSuccessCount()
}

func (c successOrErrorCollection) AddError(err error) successOrErrorCollection {
	return c.
		withErr(errors.Join(c.err, err)).
		withErrCount(c.errCount + 1)
}

func (c successOrErrorCollection) IncreaseSuccessCount() successOrErrorCollection {
	return c.withSuccessCount(c.successCount + 1)
}

func (c successOrErrorCollection) withErr(err error) successOrErrorCollection {
	c.err = err
	return c
}

func (c successOrErrorCollection) withErrCount(i int) successOrErrorCollection {
	c.errCount = i
	return c
}

func (c successOrErrorCollection) withSuccessCount(i int) successOrErrorCollection {
	c.successCount = i
	return c
}

func (c successOrErrorCollection) ErrCount() int {
	return c.errCount
}

func (c successOrErrorCollection) Error() error {
	return c.err
}

func (c successOrErrorCollection) SuccessCount() int {
	return c.successCount
}
