package model

import "github.com/google/uuid"

type AccountID string

func NewAccountID() AccountID {
	return AccountID(uuid.New().String())
}

func ParseAccountID(s string) (AccountID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return "", err
	}
	return AccountID(id.String()), nil
}
