package usecase_test

import (
	"backend/domain/model"
	"backend/infra/persistence"
	"backend/test/mock/repository"
	"backend/usecase"
	"context"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetJoinedGroups_Execute(t *testing.T) {
	ctx := context.Background()

	// テストデータ用のID生成
	accountID := model.NewAccountID()
	group1ID := model.NewGroupID()
	group2ID := model.NewGroupID()

	tests := []struct {
		name      string
		userID    string
		setupMock func(*mock.MockAccount, *mock.MockGroupMember, *mock.MockGroup)
		want      []usecase.GetJoinedGroupOutput
		wantErr   error
	}{
		{
			name:   "正常ケース: ユーザーが複数のグループに参加している場合",
			userID: "test-user-id",
			setupMock: func(acc *mock.MockAccount, gm *mock.MockGroupMember, g *mock.MockGroup) {
				account := model.Account{ID: accountID, UserID: "test-user-id", Name: "Test Account"}
				members := []model.GroupMember{
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: accountID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group2ID,
						AccountID: accountID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusActive,
					},
				}
				groups := []model.Group{
					{
						ID:          group1ID,
						Name:        "Group 1",
						Description: "テストグループ1",
					},
					{
						ID:          group2ID,
						Name:        "Group 2",
						Description: "テストグループ2",
					},
				}

				acc.EXPECT().FindByUserID(ctx, "test-user-id").Return(account, nil)
				gm.EXPECT().FindByAccountID(ctx, accountID).Return(members, nil)
				g.EXPECT().FindByIDs(ctx, []model.GroupID{group1ID, group2ID}).Return(groups, nil)
				gm.EXPECT().CountByGroupID(ctx, group1ID).Return(5, nil)
				gm.EXPECT().CountByGroupID(ctx, group2ID).Return(3, nil)
			},
			want: []usecase.GetJoinedGroupOutput{
				{
					ID:          group1ID.String(),
					Name:        "Group 1",
					MemberCount: 5,
				},
				{
					ID:          group2ID.String(),
					Name:        "Group 2",
					MemberCount: 3,
				},
			},
			wantErr: nil,
		},
		{
			name:   "正常ケース: ユーザーがグループに参加していない場合",
			userID: "test-user-id",
			setupMock: func(acc *mock.MockAccount, gm *mock.MockGroupMember, g *mock.MockGroup) {
				account := model.Account{ID: accountID, UserID: "test-user-id", Name: "Test Account"}
				members := []model.GroupMember{}

				acc.EXPECT().FindByUserID(ctx, "test-user-id").Return(account, nil)
				gm.EXPECT().FindByAccountID(ctx, accountID).Return(members, nil)
			},
			want:    []usecase.GetJoinedGroupOutput{},
			wantErr: nil,
		},
		{
			name:   "正常ケース: 1つのグループに参加している場合",
			userID: "test-user-id",
			setupMock: func(acc *mock.MockAccount, gm *mock.MockGroupMember, g *mock.MockGroup) {
				account := model.Account{ID: accountID, UserID: "test-user-id", Name: "Test Account"}
				members := []model.GroupMember{
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: accountID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
				}
				groups := []model.Group{
					{
						ID:          group1ID,
						Name:        "Solo Group",
						Description: "一人のグループ",
					},
				}

				acc.EXPECT().FindByUserID(ctx, "test-user-id").Return(account, nil)
				gm.EXPECT().FindByAccountID(ctx, accountID).Return(members, nil)
				g.EXPECT().FindByIDs(ctx, []model.GroupID{group1ID}).Return(groups, nil)
				gm.EXPECT().CountByGroupID(ctx, group1ID).Return(1, nil)
			},
			want: []usecase.GetJoinedGroupOutput{
				{
					ID:          group1ID.String(),
					Name:        "Solo Group",
					MemberCount: 1,
				},
			},
			wantErr: nil,
		},
		{
			name:   "異常ケース: アカウントが見つからない場合",
			userID: "non-existent-user",
			setupMock: func(acc *mock.MockAccount, gm *mock.MockGroupMember, g *mock.MockGroup) {
				acc.EXPECT().FindByUserID(ctx, "non-existent-user").Return(model.Account{}, persistence.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: persistence.ErrRecordNotFound,
		},
		{
			name:   "異常ケース: グループメンバー取得でエラーが発生した場合",
			userID: "test-user-id",
			setupMock: func(acc *mock.MockAccount, gm *mock.MockGroupMember, g *mock.MockGroup) {
				account := model.Account{ID: accountID, UserID: "test-user-id", Name: "Test Account"}

				acc.EXPECT().FindByUserID(ctx, "test-user-id").Return(account, nil)
				gm.EXPECT().FindByAccountID(ctx, accountID).Return(nil, errors.New("database error"))
			},
			want:    nil,
			wantErr: errors.New("database error"),
		},
		{
			name:   "異常ケース: グループ情報取得でエラーが発生した場合",
			userID: "test-user-id",
			setupMock: func(acc *mock.MockAccount, gm *mock.MockGroupMember, g *mock.MockGroup) {
				account := model.Account{ID: accountID, UserID: "test-user-id", Name: "Test Account"}
				members := []model.GroupMember{
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: accountID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
				}

				acc.EXPECT().FindByUserID(ctx, "test-user-id").Return(account, nil)
				gm.EXPECT().FindByAccountID(ctx, accountID).Return(members, nil)
				g.EXPECT().FindByIDs(ctx, []model.GroupID{group1ID}).Return(nil, errors.New("database error"))
			},
			want:    nil,
			wantErr: errors.New("database error"),
		},
		{
			name:   "異常ケース: メンバー数取得でエラーが発生した場合",
			userID: "test-user-id",
			setupMock: func(acc *mock.MockAccount, gm *mock.MockGroupMember, g *mock.MockGroup) {
				account := model.Account{ID: accountID, UserID: "test-user-id", Name: "Test Account"}
				members := []model.GroupMember{
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: accountID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
				}
				groups := []model.Group{
					{
						ID:          group1ID,
						Name:        "Test Group",
						Description: "テストグループ",
					},
				}

				acc.EXPECT().FindByUserID(ctx, "test-user-id").Return(account, nil)
				gm.EXPECT().FindByAccountID(ctx, accountID).Return(members, nil)
				g.EXPECT().FindByIDs(ctx, []model.GroupID{group1ID}).Return(groups, nil)
				gm.EXPECT().CountByGroupID(ctx, group1ID).Return(0, errors.New("database error"))
			},
			want:    nil,
			wantErr: errors.New("database error"),
		},
		{
			name:   "異常ケース: 空のユーザーIDが渡された場合",
			userID: "",
			setupMock: func(acc *mock.MockAccount, gm *mock.MockGroupMember, g *mock.MockGroup) {
				acc.EXPECT().FindByUserID(ctx, "").Return(model.Account{}, persistence.ErrInvalidInput)
			},
			want:    nil,
			wantErr: persistence.ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAccount := mock.NewMockAccount(ctrl)
			mockGroupMember := mock.NewMockGroupMember(ctrl)
			mockGroup := mock.NewMockGroup(ctrl)

			getJoinedGroups := usecase.NewGetJoinedGroups(mockAccount, mockGroupMember, mockGroup)

			tt.setupMock(mockAccount, mockGroupMember, mockGroup)

			got, err := getJoinedGroups.Execute(ctx, tt.userID)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr.Error())
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}