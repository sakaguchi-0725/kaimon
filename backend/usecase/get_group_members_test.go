package usecase_test

import (
	"backend/domain/model"
	"backend/infra/persistence"
	mock "backend/test/mock/repository"
	"backend/usecase"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetGroupMembers_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGroupRepo := mock.NewMockGroup(ctrl)
	mockAccountRepo := mock.NewMockAccount(ctrl)

	// テストデータ
	groupID := model.NewGroupID()
	userID := "user-1"
	account1ID := model.NewAccountID()
	account2ID := model.NewAccountID()
	account3ID := model.NewAccountID()

	testMembers := []model.GroupMember{
		{
			ID:        model.NewGroupMemberID(),
			GroupID:   groupID,
			AccountID: account1ID,
			Role:      model.MemberRoleAdmin,
			Status:    model.MemberStatusActive,
		},
		{
			ID:        model.NewGroupMemberID(),
			GroupID:   groupID,
			AccountID: account2ID,
			Role:      model.MemberRoleMember,
			Status:    model.MemberStatusActive,
		},
		{
			ID:        model.NewGroupMemberID(),
			GroupID:   groupID,
			AccountID: account3ID,
			Role:      model.MemberRoleMember,
			Status:    model.MemberStatusPending,
		},
	}

	testAccounts := []model.Account{
		{ID: account1ID, UserID: userID, Name: "管理者"},
		{ID: account2ID, UserID: "user-2", Name: "メンバー1"},
		{ID: account3ID, UserID: "user-3", Name: "メンバー2"},
	}

	tests := []struct {
		name      string
		input     usecase.GetGroupMembersInput
		setupMock func()
		want      usecase.GetGroupMembersOutput
		wantErr   error
	}{
		{
			name: "正常ケース: 管理者がメンバー一覧を取得",
			input: usecase.GetGroupMembersInput{
				UserID:  userID,
				GroupID: groupID.String(),
			},
			setupMock: func() {
				// ユーザーのアカウント取得
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), userID).
					Return(testAccounts[0], nil)

				// グループ取得（メンバー情報含む）
				testGroup := model.Group{
					ID:      groupID,
					Name:    "テストグループ",
					Members: testMembers,
				}
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), groupID).
					Return(testGroup, nil)

				// 全メンバーのアカウント情報を一括取得
				accountIDs := []model.AccountID{account1ID, account2ID, account3ID}
				mockAccountRepo.EXPECT().
					FindByIDs(gomock.Any(), accountIDs).
					Return(testAccounts, nil)
			},
			want: usecase.GetGroupMembersOutput{
				Members: []usecase.Member{
					{
						ID:     testMembers[0].ID.String(),
						Name:   "管理者",
						Role:   "admin",
						Status: "active",
					},
					{
						ID:     testMembers[1].ID.String(),
						Name:   "メンバー1",
						Role:   "member",
						Status: "active",
					},
					{
						ID:     testMembers[2].ID.String(),
						Name:   "メンバー2",
						Role:   "member",
						Status: "pending",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "正常ケース: 一般メンバーがメンバー一覧を取得",
			input: usecase.GetGroupMembersInput{
				UserID:  "user-2",
				GroupID: groupID.String(),
			},
			setupMock: func() {
				// ユーザーのアカウント取得
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), "user-2").
					Return(testAccounts[1], nil)

				// グループ取得（メンバー情報含む）
				testGroup := model.Group{
					ID:      groupID,
					Name:    "テストグループ",
					Members: testMembers,
				}
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), groupID).
					Return(testGroup, nil)

				// 全メンバーのアカウント情報を一括取得
				accountIDs := []model.AccountID{account1ID, account2ID, account3ID}
				mockAccountRepo.EXPECT().
					FindByIDs(gomock.Any(), accountIDs).
					Return(testAccounts, nil)
			},
			want: usecase.GetGroupMembersOutput{
				Members: []usecase.Member{
					{
						ID:     testMembers[0].ID.String(),
						Name:   "管理者",
						Role:   "admin",
						Status: "active",
					},
					{
						ID:     testMembers[1].ID.String(),
						Name:   "メンバー1",
						Role:   "member",
						Status: "active",
					},
					{
						ID:     testMembers[2].ID.String(),
						Name:   "メンバー2",
						Role:   "member",
						Status: "pending",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "正常ケース: メンバーが1人のグループ（作成者のみ）",
			input: usecase.GetGroupMembersInput{
				UserID:  userID,
				GroupID: groupID.String(),
			},
			setupMock: func() {
				singleMember := []model.GroupMember{testMembers[0]}

				// ユーザーのアカウント取得
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), userID).
					Return(testAccounts[0], nil)

				// グループ取得（1人のメンバーのみ）
				testGroup := model.Group{
					ID:      groupID,
					Name:    "テストグループ",
					Members: singleMember,
				}
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), groupID).
					Return(testGroup, nil)

				// アカウント情報取得
				accountIDs := []model.AccountID{account1ID}
				mockAccountRepo.EXPECT().
					FindByIDs(gomock.Any(), accountIDs).
					Return([]model.Account{testAccounts[0]}, nil)
			},
			want: usecase.GetGroupMembersOutput{
				Members: []usecase.Member{
					{
						ID:     testMembers[0].ID.String(),
						Name:   "管理者",
						Role:   "admin",
						Status: "active",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常ケース: 存在しないグループID",
			input: usecase.GetGroupMembersInput{
				UserID:  userID,
				GroupID: groupID.String(),
			},
			setupMock: func() {
				// ユーザーのアカウント取得
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), userID).
					Return(testAccounts[0], nil)

				// 存在しないグループID
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), groupID).
					Return(model.Group{}, persistence.ErrRecordNotFound)
			},
			want:    usecase.GetGroupMembersOutput{},
			wantErr: persistence.ErrRecordNotFound,
		},
		{
			name: "異常ケース: グループメンバーでないユーザーがアクセス",
			input: usecase.GetGroupMembersInput{
				UserID:  "non-member-user",
				GroupID: groupID.String(),
			},
			setupMock: func() {
				nonMemberAccountID := model.NewAccountID()
				nonMemberAccount := model.Account{
					ID:     nonMemberAccountID,
					UserID: "non-member-user",
					Name:   "非メンバー",
				}

				// ユーザーのアカウント取得
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), "non-member-user").
					Return(nonMemberAccount, nil)

				// グループ取得（非メンバーユーザーは含まれない）
				testGroup := model.Group{
					ID:      groupID,
					Name:    "テストグループ",
					Members: testMembers, // 非メンバーユーザーは含まれない
				}
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), groupID).
					Return(testGroup, nil)
			},
			want:    usecase.GetGroupMembersOutput{},
			wantErr: errors.New("not a member of the group"),
		},
		{
			name: "異常ケース: 不正なGroupID（空文字）",
			input: usecase.GetGroupMembersInput{
				UserID:  userID,
				GroupID: "",
			},
			setupMock: func() {
				// ユーザーのアカウント取得
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), userID).
					Return(testAccounts[0], nil)
			},
			want:    usecase.GetGroupMembersOutput{},
			wantErr: errors.New("invalid group id"),
		},
		{
			name: "異常ケース: GroupMemberRepository.FindByGroupIDでエラー",
			input: usecase.GetGroupMembersInput{
				UserID:  userID,
				GroupID: groupID.String(),
			},
			setupMock: func() {
				// ユーザーのアカウント取得
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), userID).
					Return(testAccounts[0], nil)

				// グループ取得でエラー
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), groupID).
					Return(model.Group{}, errors.New("repository error"))
			},
			want:    usecase.GetGroupMembersOutput{},
			wantErr: errors.New("repository error"),
		},
		{
			name: "異常ケース: AccountRepository.FindByIDsでエラー",
			input: usecase.GetGroupMembersInput{
				UserID:  userID,
				GroupID: groupID.String(),
			},
			setupMock: func() {
				// ユーザーのアカウント取得
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), userID).
					Return(testAccounts[0], nil)

				// グループ取得（メンバー情報含む）
				testGroup := model.Group{
					ID:      groupID,
					Name:    "テストグループ",
					Members: testMembers,
				}
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), groupID).
					Return(testGroup, nil)

				// アカウント情報一括取得でエラー
				accountIDs := []model.AccountID{account1ID, account2ID, account3ID}
				mockAccountRepo.EXPECT().
					FindByIDs(gomock.Any(), accountIDs).
					Return([]model.Account{}, errors.New("account repository error"))
			},
			want:    usecase.GetGroupMembersOutput{},
			wantErr: errors.New("account repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			uc := usecase.NewGetGroupMembers(mockAccountRepo, mockGroupRepo)
			got, err := uc.Execute(context.Background(), tt.input)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
