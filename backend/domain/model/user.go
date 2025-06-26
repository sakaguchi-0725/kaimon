package model

import "errors"

type User struct {
	ID    string
	Email string
}

func NewUser(id string, email string) (User, error) {
	if id == "" {
		return User{}, errors.New("id is required")
	}
	if email == "" {
		return User{}, errors.New("email is required")
	}

	return User{ID: id, Email: email}, nil
}
