package model

import (
	"backend/core"
	"errors"
)

type Account struct {
	ID              AccountID
	UserID          string
	Name            string
	ProfileImageURL string
}

func NewAccount(id AccountID, userID string, name string, profileImageURL string) (Account, error) {
	if userID == "" {
		return Account{}, core.NewInvalidError(errors.New("userID is required"))
	}
	if name == "" {
		return Account{}, core.NewInvalidError(errors.New("name is required"))
	}

	return Account{ID: id, UserID: userID, Name: name, ProfileImageURL: profileImageURL}, nil
}
