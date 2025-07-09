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
