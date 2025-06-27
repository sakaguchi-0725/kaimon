package persistence

import (
	"backend/core"
	"errors"
)

var (
	ErrRecordNotFound     = core.NewAppError(core.ErrNotFound, errors.New("record not found"))
	ErrInvalidInput       = core.NewInvalidError(errors.New("invalid input"))
	ErrDuplicateRecord    = core.NewInvalidError(errors.New("record already exists"))
	ErrInvalidToken       = core.NewInvalidError(errors.New("invalid token"))
	ErrTokenExpired       = core.NewInvalidError(errors.New("token expired"))
	ErrAuthenticationFail = core.NewAppError(core.ErrUnauthorized, errors.New("authentication failed"))
)
