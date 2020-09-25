package error

import "errors"

// errors
var (
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("not found error")
)

func compare(err, target error) bool {
	return errors.Is(err, target)
}
