package model

import (
	"backend/core"
	"errors"
)

type User struct {
	ID    string
	Email string
}

func NewUser(id string, email string) (User, error) {
	if id == "" {
		return User{}, core.NewInvalidError(errors.New("id is required"))
	}
	if email == "" {
		return User{}, core.NewInvalidError(errors.New("email is required"))
	}

	return User{ID: id, Email: email}, nil
}
