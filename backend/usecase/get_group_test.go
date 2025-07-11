package usecase_test

import (
	"backend/core"
	"backend/domain/model"
	mock "backend/test/mock/repository"
	"backend/usecase"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetGroup_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// テスト用データ
	testTime := core.NowJST()
	testGroupID := model.NewGroupID()
	testUserID := "test-user-id"
	testAccountID := model.NewAccountID()

	testGroup := model.Group{
		ID:          testGroupID,
		Name:        "テストグループ",
		Description: "テストグループの説明",
		CreatedAt:   testTime,
		Members: []model.GroupMember{
			{
				ID:        model.NewGroupMemberID(),
				GroupID:   testGroupID,
				AccountID: testAccountID,
				Role:      model.MemberRoleAdmin,
				Status:    model.MemberStatusActive,
			},
			{
				ID:        model.NewGroupMemberID(),
				GroupID:   testGroupID,
				AccountID: model.NewAccountID(),
				Role:      model.MemberRoleMember,
				Status:    model.MemberStatusActive,
			},
			{
				ID:        model.NewGroupMemberID(),
				GroupID:   testGroupID,
				AccountID: model.NewAccountID(),
				Role:      model.MemberRoleMember,
				Status:    model.MemberStatusPending,
			},
		},
	}

	testAccount := model.Account{
		ID:     testAccountID,
		UserID: testUserID,
		Name:   "テストユーザー",
	}

	testAccounts := []model.Account{
		testAccount,
		{
			ID:     testGroup.Members[1].AccountID,
			UserID: "member-user-id",
			Name:   "メンバーユーザー",
		},
		{
			ID:     testGroup.Members[2].AccountID,
			UserID: "pending-user-id",
			Name:   "保留ユーザー",
		},
	}

	tests := []struct {
		name    string
		input   usecase.GetGroupInput
		setup   func(mockAccountRepo *mock.MockAccount, mockGroupRepo *mock.MockGroup)
		want    usecase.GetGroupOutput
		wantErr error
	}{
		{
			name: "管理者がグループ詳細とメンバー一覧を取得",
			input: usecase.GetGroupInput{
				UserID:  testUserID,
				GroupID: testGroupID.String(),
			},
			setup: func(mockAccountRepo *mock.MockAccount, mockGroupRepo *mock.MockGroup) {
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), testUserID).
					Return(testAccount, nil)
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), testGroupID).
					Return(testGroup, nil)
				mockAccountRepo.EXPECT().
					FindByIDs(gomock.Any(), gomock.Any()).
					Return(testAccounts, nil)
			},
			want: usecase.GetGroupOutput{
				ID:          testGroupID.String(),
				Name:        "テストグループ",
				Description: "テストグループの説明",
				CreatedAt:   testTime.Format("2006-01-02T15:04:05Z"),
				Members: []usecase.Member{
					{
						ID:        testGroup.Members[0].ID.String(),
						AccountID: testGroup.Members[0].AccountID.String(),
						Name:      "テストユーザー",
						Role:      "admin",
						Status:    "active",
					},
					{
						ID:        testGroup.Members[1].ID.String(),
						AccountID: testGroup.Members[1].AccountID.String(),
						Name:      "メンバーユーザー",
						Role:      "member",
						Status:    "active",
					},
					{
						ID:        testGroup.Members[2].ID.String(),
						AccountID: testGroup.Members[2].AccountID.String(),
						Name:      "保留ユーザー",
						Role:      "member",
						Status:    "pending",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "一般メンバーがグループ詳細とメンバー一覧を取得",
			input: usecase.GetGroupInput{
				UserID:  "member-user-id",
				GroupID: testGroupID.String(),
			},
			setup: func(mockAccountRepo *mock.MockAccount, mockGroupRepo *mock.MockGroup) {
				memberAccount := model.Account{
					ID:     testGroup.Members[1].AccountID,
					UserID: "member-user-id",
					Name:   "メンバーユーザー",
				}
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), "member-user-id").
					Return(memberAccount, nil)
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), testGroupID).
					Return(testGroup, nil)
				mockAccountRepo.EXPECT().
					FindByIDs(gomock.Any(), gomock.Any()).
					Return(testAccounts, nil)
			},
			want: usecase.GetGroupOutput{
				ID:          testGroupID.String(),
				Name:        "テストグループ",
				Description: "テストグループの説明",
				CreatedAt:   testTime.Format("2006-01-02T15:04:05Z"),
				Members: []usecase.Member{
					{
						ID:        testGroup.Members[0].ID.String(),
						AccountID: testGroup.Members[0].AccountID.String(),
						Name:      "テストユーザー",
						Role:      "admin",
						Status:    "active",
					},
					{
						ID:        testGroup.Members[1].ID.String(),
						AccountID: testGroup.Members[1].AccountID.String(),
						Name:      "メンバーユーザー",
						Role:      "member",
						Status:    "active",
					},
					{
						ID:        testGroup.Members[2].ID.String(),
						AccountID: testGroup.Members[2].AccountID.String(),
						Name:      "保留ユーザー",
						Role:      "member",
						Status:    "pending",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "メンバーが1人のグループ（作成者のみ）",
			input: usecase.GetGroupInput{
				UserID:  testUserID,
				GroupID: testGroupID.String(),
			},
			setup: func(mockAccountRepo *mock.MockAccount, mockGroupRepo *mock.MockGroup) {
				singleMemberID := model.NewGroupMemberID()
				singleMemberGroup := model.Group{
					ID:          testGroupID,
					Name:        "単一メンバーグループ",
					Description: "作成者のみのグループ",
					CreatedAt:   testTime,
					Members: []model.GroupMember{
						{
							ID:        singleMemberID,
							GroupID:   testGroupID,
							AccountID: testAccountID,
							Role:      model.MemberRoleAdmin,
							Status:    model.MemberStatusActive,
						},
					},
				}
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), testUserID).
					Return(testAccount, nil)
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), testGroupID).
					Return(singleMemberGroup, nil)
				mockAccountRepo.EXPECT().
					FindByIDs(gomock.Any(), gomock.Any()).
					Return([]model.Account{testAccount}, nil)
			},
			want: usecase.GetGroupOutput{
				ID:          testGroupID.String(),
				Name:        "単一メンバーグループ",
				Description: "作成者のみのグループ",
				CreatedAt:   testTime.Format("2006-01-02T15:04:05Z"),
				Members: []usecase.Member{},
			},
			wantErr: nil,
		},
		{
			name: "存在しないグループID",
			input: usecase.GetGroupInput{
				UserID:  testUserID,
				GroupID: testGroupID.String(),
			},
			setup: func(mockAccountRepo *mock.MockAccount, mockGroupRepo *mock.MockGroup) {
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), testUserID).
					Return(testAccount, nil)
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), testGroupID).
					Return(model.Group{}, core.NewAppError(core.ErrNotFound, errors.New("group not found")))
			},
			want:    usecase.GetGroupOutput{},
			wantErr: errors.New("group not found"),
		},
		{
			name: "グループメンバーでないユーザーがアクセス",
			input: usecase.GetGroupInput{
				UserID:  "non-member-user-id",
				GroupID: testGroupID.String(),
			},
			setup: func(mockAccountRepo *mock.MockAccount, mockGroupRepo *mock.MockGroup) {
				nonMemberAccount := model.Account{
					ID:     model.NewAccountID(),
					UserID: "non-member-user-id",
					Name:   "非メンバーユーザー",
				}
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), "non-member-user-id").
					Return(nonMemberAccount, nil)
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), testGroupID).
					Return(testGroup, nil)
			},
			want:    usecase.GetGroupOutput{},
			wantErr: errors.New("not a member of the group"),
		},
		{
			name: "不正なGroupID（空文字）",
			input: usecase.GetGroupInput{
				UserID:  testUserID,
				GroupID: "",
			},
			setup: func(mockAccountRepo *mock.MockAccount, mockGroupRepo *mock.MockGroup) {
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), testUserID).
					Return(testAccount, nil)
			},
			want:    usecase.GetGroupOutput{},
			wantErr: core.NewInvalidError(errors.New("")),
		},
		{
			name: "GroupRepository.GetByIDでエラー",
			input: usecase.GetGroupInput{
				UserID:  testUserID,
				GroupID: testGroupID.String(),
			},
			setup: func(mockAccountRepo *mock.MockAccount, mockGroupRepo *mock.MockGroup) {
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), testUserID).
					Return(testAccount, nil)
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), testGroupID).
					Return(model.Group{}, errors.New("database error"))
			},
			want:    usecase.GetGroupOutput{},
			wantErr: errors.New("database error"),
		},
		{
			name: "AccountRepository.FindByIDsでエラー",
			input: usecase.GetGroupInput{
				UserID:  testUserID,
				GroupID: testGroupID.String(),
			},
			setup: func(mockAccountRepo *mock.MockAccount, mockGroupRepo *mock.MockGroup) {
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), testUserID).
					Return(testAccount, nil)
				mockGroupRepo.EXPECT().
					GetByID(gomock.Any(), testGroupID).
					Return(testGroup, nil)
				mockAccountRepo.EXPECT().
					FindByIDs(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("account fetch error"))
			},
			want:    usecase.GetGroupOutput{},
			wantErr: errors.New("account fetch error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAccountRepo := mock.NewMockAccount(ctrl)
			mockGroupRepo := mock.NewMockGroup(ctrl)

			tt.setup(mockAccountRepo, mockGroupRepo)

			useCase := usecase.NewGetGroup(mockAccountRepo, mockGroupRepo)
			got, err := useCase.Execute(context.Background(), tt.input)

			if tt.wantErr != nil {
				assert.Error(t, err)
				if tt.wantErr.Error() != "" {
					assert.Contains(t, err.Error(), tt.wantErr.Error())
				}
			} else {
				assert.NoError(t, err)
				if tt.name == "メンバーが1人のグループ（作成者のみ）" {
					assert.Equal(t, tt.want.ID, got.ID)
					assert.Equal(t, tt.want.Name, got.Name)
					assert.Equal(t, tt.want.Description, got.Description)
					assert.Equal(t, tt.want.CreatedAt, got.CreatedAt)
					assert.Len(t, got.Members, 1)
					assert.Equal(t, "テストユーザー", got.Members[0].Name)
					assert.Equal(t, "admin", got.Members[0].Role)
					assert.Equal(t, "active", got.Members[0].Status)
				} else {
					assert.Equal(t, tt.want, got)
				}
			}
		})
	}
}
