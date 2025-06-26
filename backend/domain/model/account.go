package model

import "errors"

type Account struct {
	ID     AccountID
	UserID string
	Name   string
}

func NewAccount(id AccountID, userID string, name string) (Account, error) {
	if userID == "" {
		return Account{}, errors.New("userID is required")
	}
	if name == "" {
		return Account{}, errors.New("name is required")
	}

	return Account{ID: id, UserID: userID, Name: name}, nil
}
