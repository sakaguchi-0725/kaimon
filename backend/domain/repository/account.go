//go:generate mockgen -source=account.go -destination=../../test/mock/repository/account_mock.go -package=mock
package repository

import (
	"backend/domain/model"
	"context"
)

type Account interface {
	FindByID(ctx context.Context, id model.AccountID) (model.Account, error)
	FindByUserID(ctx context.Context, userID string) (model.Account, error)
	FindByIDs(ctx context.Context, ids []model.AccountID) ([]model.Account, error)
	Store(ctx context.Context, acc *model.Account) error
}
