package error

import (
	"errors"
	"fmt"
)

type stringValueError struct {
	msg string
}

func (e stringValueError) Error() string {
	return e.msg
}

func (e stringValueError) is(err error) bool {
	return errors.Is(e, err)
}

func (e stringValueError) as(err error) bool {
	var es stringValueError
	return errors.As(err, &es)
}

type code struct {
	code uint
}

type value struct {
	str string
}

type structValueError struct {
	code  code
	value value
}

func (e structValueError) is(err error) bool {
	return errors.Is(e, err)
}

func (e structValueError) Error() string {
	return fmt.Sprintf("%d: %s", e.code.code, e.value.str)
}

type interfaceValueError struct {
	err error
}

func (e interfaceValueError) Error() string {
	return e.err.Error()
}

func (e interfaceValueError) is(err error) bool {
	return errors.Is(e, err)
}
