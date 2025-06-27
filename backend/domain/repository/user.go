//go:generate mockgen -source=user.go -destination=../../test/mock/repository/user_mock.go -package=mock
package repository

import (
	"backend/domain/model"
	"context"
)

type User interface {
	FindByID(ctx context.Context, id string) (model.User, error)
	FindByUID(ctx context.Context, uid string) (*model.User, error)
	Store(ctx context.Context, user *model.User) error
}
