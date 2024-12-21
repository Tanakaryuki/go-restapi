package errors

import (
	"runtime/debug"
)

const (
	ErrEmailInUse      = "email already in use"
	ErrUsernameInUse   = "username already in use"
	ErrInvalidPassword = "invalid password"
	ErrIDInUse         = "id already in use"
)

type Error struct {
	Status     int
	Err        error
	StackTrace string
}

func New(status int, err error) *Error {
	return &Error{
		Err:        err,
		Status:     status,
		StackTrace: string(debug.Stack()),
	}
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func (e *Error) As(target interface{}) bool {
	switch v := target.(type) {
	case **Error:
		*v = e
		return true
	default:
		return false
	}
}

func As(err error, target interface{}) bool {
	return err.(*Error).As(target)
}
