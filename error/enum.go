package error

import "errors"

// Code .
type Code uint

// Error Code
const (
	Zero Code = iota
	One
)

func (code Code) Error() string {
	return [...]string{
		"Error: Zero",
		"Error: One",
	}[code]
}

// Is .
func (code Code) is(err error) bool {
	return errors.Is(code, err)
}

// As .
func (code Code) as(err error) bool {
	var c Code
	return errors.As(err, &c)
}
