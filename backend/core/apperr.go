package core

import (
	"fmt"
	"runtime"
)

type AppError struct {
	code    ErrorCode
	file    string
	line    int
	message string
	error
}

func NewAppError(code ErrorCode, err error) *AppError {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	return &AppError{
		code:    code,
		file:    file,
		line:    line,
		message: code.DefaultMessage(),
		error:   err,
	}
}

func NewInvalidError(err error) *AppError {
	return NewAppError(ErrBadRequest, err)
}

func (e *AppError) WithMessage(msg string) *AppError {
	e.message = msg
	return e
}

func (e *AppError) Unwrap() error {
	return e.error
}

func (e *AppError) Is(target error) bool {
	if t, ok := target.(*AppError); ok {
		return e.code == t.code
	}
	return false
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s at %s:%d", e.error.Error(), e.file, e.line)
}

func (e *AppError) Code() ErrorCode {
	return e.code
}

func (e *AppError) Message() string {
	return e.message
}

func (e *AppError) StackTrace() string {
	return fmt.Sprintf("%s:%d", e.file, e.line)
}
