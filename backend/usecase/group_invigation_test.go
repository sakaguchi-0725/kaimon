package usecase_test

import (
	"backend/core"
	"backend/domain/model"
	mock "backend/test/mock/repository"
	"backend/usecase"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGroupInvitation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountRepo := mock.NewMockAccount(ctrl)
	groupRepo := mock.NewMockGroup(ctrl)

	tests := []struct {
		name    string
		input   usecase.GroupInvitationInput
		setup   func()
		wantErr bool
		errMsg  string
	}{
		{
			name: "正常系：管理者が招待コードを生成",
			input: usecase.GroupInvitationInput{
				GroupID: "01234567-89ab-cdef-0123-456789abcdef",
				UserID:  "test-user-id",
			},
			setup: func() {
				accountID := model.NewAccountID()
				account := model.Account{
					ID:     accountID,
					UserID: "test-user-id",
					Name:   "テスト管理者",
				}
				
				groupID, _ := model.ParseGroupID("01234567-89ab-cdef-0123-456789abcdef")
				group := model.Group{
					ID:          groupID,
					Name:        "テストグループ",
					Description: "テスト用グループ",
					Members: []model.GroupMember{
						{
							ID:        model.NewGroupMemberID(),
							AccountID: accountID,
							Role:      model.MemberRoleAdmin,
							Status:    model.MemberStatusActive,
						},
					},
					CreatedAt: core.NowJST(),
				}

				accountRepo.EXPECT().FindByUserID(gomock.Any(), "test-user-id").Return(account, nil)
				groupRepo.EXPECT().GetByID(gomock.Any(), groupID).Return(group, nil)
				groupRepo.EXPECT().Invitation(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "異常系：存在しないユーザー",
			input: usecase.GroupInvitationInput{
				GroupID: "01234567-89ab-cdef-0123-456789abcdef",
				UserID:  "non-existent-user",
			},
			setup: func() {
				accountRepo.EXPECT().FindByUserID(gomock.Any(), "non-existent-user").Return(model.Account{}, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "異常系：無効なグループID",
			input: usecase.GroupInvitationInput{
				GroupID: "invalid-group-id",
				UserID:  "test-user-id",
			},
			setup: func() {
				account := model.Account{
					ID:     model.NewAccountID(),
					UserID: "test-user-id",
					Name:   "テストユーザー",
				}
				accountRepo.EXPECT().FindByUserID(gomock.Any(), "test-user-id").Return(account, nil)
			},
			wantErr: true,
		},
		{
			name: "異常系：存在しないグループ",
			input: usecase.GroupInvitationInput{
				GroupID: "01234567-89ab-cdef-0123-456789abcdef",
				UserID:  "test-user-id",
			},
			setup: func() {
				account := model.Account{
					ID:     model.NewAccountID(),
					UserID: "test-user-id",
					Name:   "テストユーザー",
				}
				groupID, _ := model.ParseGroupID("01234567-89ab-cdef-0123-456789abcdef")
				
				accountRepo.EXPECT().FindByUserID(gomock.Any(), "test-user-id").Return(account, nil)
				groupRepo.EXPECT().GetByID(gomock.Any(), groupID).Return(model.Group{}, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "異常系：管理者権限がない",
			input: usecase.GroupInvitationInput{
				GroupID: "01234567-89ab-cdef-0123-456789abcdef",
				UserID:  "test-user-id",
			},
			setup: func() {
				accountID := model.NewAccountID()
				account := model.Account{
					ID:     accountID,
					UserID: "test-user-id",
					Name:   "テスト一般ユーザー",
				}
				
				groupID, _ := model.ParseGroupID("01234567-89ab-cdef-0123-456789abcdef")
				group := model.Group{
					ID:          groupID,
					Name:        "テストグループ",
					Description: "テスト用グループ",
					Members: []model.GroupMember{
						{
							ID:        model.NewGroupMemberID(),
							AccountID: accountID,
							Role:      model.MemberRoleMember, // 一般メンバー
							Status:    model.MemberStatusActive,
						},
					},
					CreatedAt: core.NowJST(),
				}

				accountRepo.EXPECT().FindByUserID(gomock.Any(), "test-user-id").Return(account, nil)
				groupRepo.EXPECT().GetByID(gomock.Any(), groupID).Return(group, nil)
			},
			wantErr: true,
			errMsg:  "you are not admin",
		},
		{
			name: "異常系：招待コード保存に失敗",
			input: usecase.GroupInvitationInput{
				GroupID: "01234567-89ab-cdef-0123-456789abcdef",
				UserID:  "test-user-id",
			},
			setup: func() {
				accountID := model.NewAccountID()
				account := model.Account{
					ID:     accountID,
					UserID: "test-user-id",
					Name:   "テスト管理者",
				}
				
				groupID, _ := model.ParseGroupID("01234567-89ab-cdef-0123-456789abcdef")
				group := model.Group{
					ID:          groupID,
					Name:        "テストグループ",
					Description: "テスト用グループ",
					Members: []model.GroupMember{
						{
							ID:        model.NewGroupMemberID(),
							AccountID: accountID,
							Role:      model.MemberRoleAdmin,
							Status:    model.MemberStatusActive,
						},
					},
					CreatedAt: core.NowJST(),
				}

				accountRepo.EXPECT().FindByUserID(gomock.Any(), "test-user-id").Return(account, nil)
				groupRepo.EXPECT().GetByID(gomock.Any(), groupID).Return(group, nil)
				groupRepo.EXPECT().Invitation(gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			groupInvitation := usecase.NewGroupInvitation(accountRepo, groupRepo)
			output, err := groupInvitation.Execute(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
				assert.Empty(t, output.Code)
				assert.Empty(t, output.ExpiresAt)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, output.Code)
				assert.NotEmpty(t, output.ExpiresAt)
				
				// ExpiresAtが有効なISO8601フォーマットであることを確認
				_, parseErr := time.Parse(core.LayoutISO8601, output.ExpiresAt)
				assert.NoError(t, parseErr)
			}
		})
	}
}