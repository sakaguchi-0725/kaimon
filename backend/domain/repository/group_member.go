//go:generate mockgen -source=group_member.go -destination=../../test/mock/repository/group_member_mock.go -package=mock
package repository

import (
	"backend/domain/model"
	"context"
)

type GroupMember interface {
	FindByAccountID(ctx context.Context, accountID model.AccountID) ([]model.GroupMember, error)
	FindByGroupID(ctx context.Context, groupID model.GroupID) ([]model.GroupMember, error)
	CountByGroupID(ctx context.Context, groupID model.GroupID) (int, error)
}
