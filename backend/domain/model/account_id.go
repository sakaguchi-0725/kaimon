package model

import (
	"backend/core"

	"github.com/google/uuid"
)

type AccountID string

func NewAccountID() AccountID {
	return AccountID(uuid.New().String())
}

func ParseAccountID(s string) (AccountID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return "", core.NewInvalidError(err)
	}
	return AccountID(id.String()), nil
}

func (id AccountID) String() string {
	return string(id)
}
