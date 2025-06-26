package repository

import (
	"backend/domain/model"
	"context"
)

type Account interface {
	FindByID(ctx context.Context, id model.AccountID) (model.Account, error)
	Store(ctx context.Context, acc *model.Account) error
}
