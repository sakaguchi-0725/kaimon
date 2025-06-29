//go:generate mockgen -source=group.go -destination=../../test/mock/repository/group_mock.go -package=mock
package repository

import (
	"backend/domain/model"
	"context"
)

type Group interface {
	FindByIDs(ctx context.Context, ids []model.GroupID) ([]model.Group, error)
}
