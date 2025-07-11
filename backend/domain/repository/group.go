//go:generate mockgen -source=group.go -destination=../../test/mock/repository/group_mock.go -package=mock
package repository

import (
	"backend/domain/model"
	"context"
)

type Group interface {
	GetByID(ctx context.Context, id model.GroupID) (model.Group, error)
	FindByIDs(ctx context.Context, ids []model.GroupID) ([]model.Group, error)
	Store(ctx context.Context, group *model.Group) error
	Invitation(ctx context.Context, invitation model.Invitation) error
}
