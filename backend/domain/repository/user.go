package repository

import (
	"backend/domain/model"
	"context"
)

type User interface {
	FindByID(ctx context.Context, id string) (model.User, error)
	Store(ctx context.Context, user *model.User) error
}
