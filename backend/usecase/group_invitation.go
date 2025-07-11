//go:generate mockgen -source=group_invitation.go -destination=../test/mock/usecase/group_invitation_mock.go -package=mock
package usecase

import (
	"backend/core"
	"backend/domain/model"
	"backend/domain/repository"
	"context"
	"errors"
)

type (
	GroupInvitation interface {
		Execute(ctx context.Context, input GroupInvitationInput) (GroupInvitationOutput, error)
	}

	GroupInvitationInput struct {
		GroupID string
		UserID  string
	}

	GroupInvitationOutput struct {
		Code      string
		ExpiresAt string
	}

	groupInvitationInteractor struct {
		accountRepo repository.Account
		groupRepo   repository.Group
	}
)

func (g *groupInvitationInteractor) Execute(ctx context.Context, input GroupInvitationInput) (GroupInvitationOutput, error) {
	account, err := g.accountRepo.FindByUserID(ctx, input.UserID)
	if err != nil {
		return GroupInvitationOutput{}, err
	}

	groupID, err := model.ParseGroupID(input.GroupID)
	if err != nil {
		return GroupInvitationOutput{}, err
	}

	group, err := g.groupRepo.GetByID(ctx, groupID)
	if err != nil {
		return GroupInvitationOutput{}, err
	}

	if !group.IsAdmin(account.ID) {
		return GroupInvitationOutput{}, core.NewAppError(core.ErrForbidden, errors.New("you are not admin")).
			WithMessage("管理者権限がありません")
	}

	// 既存の招待コードを確認
	existingInvitation, err := g.groupRepo.GetInvitation(ctx, groupID)
	if err != nil {
		return GroupInvitationOutput{}, err
	}

	// 既存の招待コードがある場合はそれを返す
	if existingInvitation != nil {
		return GroupInvitationOutput{
			Code:      existingInvitation.Code,
			ExpiresAt: core.FormatWithLayout(existingInvitation.ExpiresAt, core.LayoutISO8601),
		}, nil
	}

	// 既存の招待コードがない場合は新規生成
	invitation := group.Invitation()

	err = g.groupRepo.Invitation(ctx, invitation)
	if err != nil {
		return GroupInvitationOutput{}, err
	}

	return GroupInvitationOutput{
		Code:      invitation.Code,
		ExpiresAt: core.FormatWithLayout(invitation.ExpiresAt, core.LayoutISO8601),
	}, nil
}

func NewGroupInvitation(a repository.Account, g repository.Group) GroupInvitation {
	return &groupInvitationInteractor{
		accountRepo: a,
		groupRepo:   g,
	}
}
