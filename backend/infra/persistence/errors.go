package persistence

import (
	"backend/core"
	"errors"
)

var (
	ErrRecordNotFound  = core.NewAppError(core.ErrNotFound, errors.New("record not found"))
	ErrInvalidInput    = core.NewInvalidError(errors.New("invalid input"))
	ErrDuplicateRecord = core.NewInvalidError(errors.New("record already exists"))
)
