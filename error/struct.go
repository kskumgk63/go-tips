package error

import (
	"errors"
)

type err1 struct {
	msg string
}

func (e err1) Error() string {
	return e.msg
}

func (e err1) is(err error) bool {
	return errors.Is(e, err)
}

func (e err1) as(err error) bool {
	var es err1
	return errors.As(err, &es)
}

type err2 struct {
	msg string
}

func (e err2) Error() string {
	return e.msg
}
