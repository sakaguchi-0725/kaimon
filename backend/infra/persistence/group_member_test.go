package persistence_test

import (
	"backend/domain/model"
	"backend/infra/persistence"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupMemberPersistence(t *testing.T) {
	groupMemberRepo := persistence.NewGroupMember(testDB)

	t.Run("FindByAccountID", func(t *testing.T) {
		account1ID := model.NewAccountID()
		account2ID := model.NewAccountID()
		group1ID := model.NewGroupID()
		group2ID := model.NewGroupID()
		group3ID := model.NewGroupID()
		nonExistentAccountID := model.NewAccountID()

		tests := []struct {
			name               string
			accountID          model.AccountID
			existingUsers      []model.User
			existingAccounts   []model.Account
			existingGroups     []model.Group
			existingMembers    []model.GroupMember
			want               []model.GroupMember
			wantErr            error
		}{
			{
				name:      "アカウントが参加している全グループメンバーシップを取得",
				accountID: account1ID,
				existingUsers: []model.User{
					{ID: "user1", Email: "user1@example.com"},
					{ID: "user2", Email: "user2@example.com"},
				},
				existingAccounts: []model.Account{
					{ID: account1ID, UserID: "user1", Name: "Account 1"},
					{ID: account2ID, UserID: "user2", Name: "Account 2"},
				},
				existingGroups: []model.Group{
					{ID: group1ID, Name: "Group 1", Description: "グループ1"},
					{ID: group2ID, Name: "Group 2", Description: "グループ2"},
					{ID: group3ID, Name: "Group 3", Description: "グループ3"},
				},
				existingMembers: []model.GroupMember{
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: account1ID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group2ID,
						AccountID: account1ID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusActive,
					},
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group3ID,
						AccountID: account2ID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
				},
				want: []model.GroupMember{
					{
						GroupID:   group1ID,
						AccountID: account1ID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
					{
						GroupID:   group2ID,
						AccountID: account1ID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusActive,
					},
				},
				wantErr: nil,
			},
			{
				name:             "グループに参加していないアカウント",
				accountID:        nonExistentAccountID,
				existingUsers:    []model.User{},
				existingAccounts: []model.Account{},
				existingGroups:   []model.Group{},
				existingMembers:  []model.GroupMember{},
				want:             []model.GroupMember{},
				wantErr:          nil,
			},
			{
				name:      "Pendingステータスのメンバーシップも含む",
				accountID: account1ID,
				existingUsers: []model.User{
					{ID: "user1", Email: "user1@example.com"},
				},
				existingAccounts: []model.Account{
					{ID: account1ID, UserID: "user1", Name: "Account 1"},
				},
				existingGroups: []model.Group{
					{ID: group1ID, Name: "Group 1", Description: "グループ1"},
					{ID: group2ID, Name: "Group 2", Description: "グループ2"},
				},
				existingMembers: []model.GroupMember{
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: account1ID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group2ID,
						AccountID: account1ID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusPending,
					},
				},
				want: []model.GroupMember{
					{
						GroupID:   group1ID,
						AccountID: account1ID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
					{
						GroupID:   group2ID,
						AccountID: account1ID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusPending,
					},
				},
				wantErr: nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// テストデータを事前に作成
				for _, user := range tt.existingUsers {
					err := CreateTestUser(user.ID, user.Email)
					assert.NoError(t, err)
				}

				for _, account := range tt.existingAccounts {
					err := CreateTestAccount(account.ID, account.UserID, account.Name)
					assert.NoError(t, err)
				}

				for _, group := range tt.existingGroups {
					err := CreateTestGroup(group.ID, group.Name, group.Description)
					assert.NoError(t, err)
				}

				for _, member := range tt.existingMembers {
					err := CreateTestGroupMember(member.ID, member.GroupID, member.AccountID, member.Role, member.Status)
					assert.NoError(t, err)
				}

				got, err := groupMemberRepo.FindByAccountID(context.Background(), tt.accountID)

				if tt.wantErr != nil {
					assert.Error(t, err)
					assert.ErrorIs(t, err, tt.wantErr)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, len(tt.want), len(got))

					// 結果を検証（IDとJoinedAtは除外して比較）
					for _, wantMember := range tt.want {
						found := false
						for _, gotMember := range got {
							if gotMember.GroupID == wantMember.GroupID && gotMember.AccountID == wantMember.AccountID {
								assert.Equal(t, wantMember.Role, gotMember.Role)
								assert.Equal(t, wantMember.Status, gotMember.Status)
								found = true
								break
							}
						}
						assert.True(t, found, "期待されたメンバーシップが見つかりません: GroupID=%s, AccountID=%s", wantMember.GroupID, wantMember.AccountID)
					}
				}
			})
		}
	})

	t.Run("FindByGroupID", func(t *testing.T) {
		group1ID := model.NewGroupID()
		group2ID := model.NewGroupID()
		emptyGroupID := model.NewGroupID()
		nonExistentGroupID := model.NewGroupID()
		account1ID := model.NewAccountID()
		account2ID := model.NewAccountID()
		account3ID := model.NewAccountID()

		tests := []struct {
			name             string
			groupID          model.GroupID
			existingUsers    []model.User
			existingAccounts []model.Account
			existingGroups   []model.Group
			existingMembers  []model.GroupMember
			want             []model.GroupMember
			wantErr          error
		}{
			{
				name:    "複数メンバーがいるグループ",
				groupID: group1ID,
				existingUsers: []model.User{
					{ID: "user1", Email: "user1@example.com"},
					{ID: "user2", Email: "user2@example.com"},
					{ID: "user3", Email: "user3@example.com"},
				},
				existingAccounts: []model.Account{
					{ID: account1ID, UserID: "user1", Name: "Account 1"},
					{ID: account2ID, UserID: "user2", Name: "Account 2"},
					{ID: account3ID, UserID: "user3", Name: "Account 3"},
				},
				existingGroups: []model.Group{
					{ID: group1ID, Name: "Group 1", Description: "グループ1"},
					{ID: group2ID, Name: "Group 2", Description: "グループ2"},
				},
				existingMembers: []model.GroupMember{
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: account1ID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: account2ID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusActive,
					},
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: account3ID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusPending,
					},
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group2ID,
						AccountID: account1ID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
				},
				want: []model.GroupMember{
					{
						GroupID:   group1ID,
						AccountID: account1ID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
					{
						GroupID:   group1ID,
						AccountID: account2ID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusActive,
					},
					{
						GroupID:   group1ID,
						AccountID: account3ID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusPending,
					},
				},
				wantErr: nil,
			},
			{
				name:             "メンバーがいないグループ",
				groupID:          emptyGroupID,
				existingUsers:    []model.User{},
				existingAccounts: []model.Account{},
				existingGroups: []model.Group{
					{ID: emptyGroupID, Name: "Empty Group", Description: "空のグループ"},
				},
				existingMembers: []model.GroupMember{},
				want:            []model.GroupMember{},
				wantErr:         nil,
			},
			{
				name:             "存在しないグループ",
				groupID:          nonExistentGroupID,
				existingUsers:    []model.User{},
				existingAccounts: []model.Account{},
				existingGroups:   []model.Group{},
				existingMembers:  []model.GroupMember{},
				want:             []model.GroupMember{},
				wantErr:          nil,
			},
			{
				name:    "異なるステータスのメンバーが混在するグループ",
				groupID: group1ID,
				existingUsers: []model.User{
					{ID: "user1", Email: "user1@example.com"},
					{ID: "user2", Email: "user2@example.com"},
				},
				existingAccounts: []model.Account{
					{ID: account1ID, UserID: "user1", Name: "Account 1"},
					{ID: account2ID, UserID: "user2", Name: "Account 2"},
				},
				existingGroups: []model.Group{
					{ID: group1ID, Name: "Group 1", Description: "グループ1"},
				},
				existingMembers: []model.GroupMember{
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: account1ID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: account2ID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusPending,
					},
				},
				want: []model.GroupMember{
					{
						GroupID:   group1ID,
						AccountID: account1ID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
					{
						GroupID:   group1ID,
						AccountID: account2ID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusPending,
					},
				},
				wantErr: nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// テストデータを事前に作成
				for _, user := range tt.existingUsers {
					err := CreateTestUser(user.ID, user.Email)
					assert.NoError(t, err)
				}

				for _, account := range tt.existingAccounts {
					err := CreateTestAccount(account.ID, account.UserID, account.Name)
					assert.NoError(t, err)
				}

				for _, group := range tt.existingGroups {
					err := CreateTestGroup(group.ID, group.Name, group.Description)
					assert.NoError(t, err)
				}

				for _, member := range tt.existingMembers {
					err := CreateTestGroupMember(member.ID, member.GroupID, member.AccountID, member.Role, member.Status)
					assert.NoError(t, err)
				}

				got, err := groupMemberRepo.FindByGroupID(context.Background(), tt.groupID)

				if tt.wantErr != nil {
					assert.Error(t, err)
					assert.ErrorIs(t, err, tt.wantErr)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, len(tt.want), len(got))

					// 結果を検証（IDとJoinedAtは除外して比較）
					for _, wantMember := range tt.want {
						found := false
						for _, gotMember := range got {
							if gotMember.GroupID == wantMember.GroupID && gotMember.AccountID == wantMember.AccountID {
								assert.Equal(t, wantMember.Role, gotMember.Role)
								assert.Equal(t, wantMember.Status, gotMember.Status)
								found = true
								break
							}
						}
						assert.True(t, found, "期待されたメンバーが見つかりません: GroupID=%s, AccountID=%s", wantMember.GroupID, wantMember.AccountID)
					}
				}
			})
		}
	})

	t.Run("CountByGroupID", func(t *testing.T) {
		group1ID := model.NewGroupID()
		group2ID := model.NewGroupID()
		emptyGroupID := model.NewGroupID()
		nonExistentGroupID := model.NewGroupID()
		account1ID := model.NewAccountID()
		account2ID := model.NewAccountID()
		account3ID := model.NewAccountID()

		tests := []struct {
			name             string
			groupID          model.GroupID
			existingUsers    []model.User
			existingAccounts []model.Account
			existingGroups   []model.Group
			existingMembers  []model.GroupMember
			want             int
			wantErr          error
		}{
			{
				name:    "複数メンバーがいるグループ",
				groupID: group1ID,
				existingUsers: []model.User{
					{ID: "user1", Email: "user1@example.com"},
					{ID: "user2", Email: "user2@example.com"},
					{ID: "user3", Email: "user3@example.com"},
				},
				existingAccounts: []model.Account{
					{ID: account1ID, UserID: "user1", Name: "Account 1"},
					{ID: account2ID, UserID: "user2", Name: "Account 2"},
					{ID: account3ID, UserID: "user3", Name: "Account 3"},
				},
				existingGroups: []model.Group{
					{ID: group1ID, Name: "Group 1", Description: "グループ1"},
					{ID: group2ID, Name: "Group 2", Description: "グループ2"},
				},
				existingMembers: []model.GroupMember{
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: account1ID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: account2ID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusActive,
					},
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group1ID,
						AccountID: account3ID,
						Role:      model.MemberRoleMember,
						Status:    model.MemberStatusPending,
					},
					{
						ID:        model.NewGroupMemberID(),
						GroupID:   group2ID,
						AccountID: account1ID,
						Role:      model.MemberRoleAdmin,
						Status:    model.MemberStatusActive,
					},
				},
				want:    3, // Active + Pending の両方をカウント
				wantErr: nil,
			},
			{
				name:             "メンバーがいないグループ",
				groupID:          emptyGroupID,
				existingUsers:    []model.User{},
				existingAccounts: []model.Account{},
				existingGroups: []model.Group{
					{ID: emptyGroupID, Name: "Empty Group", Description: "空のグループ"},
				},
				existingMembers: []model.GroupMember{},
				want:            0,
				wantErr:         nil,
			},
			{
				name:             "存在しないグループ",
				groupID:          nonExistentGroupID,
				existingUsers:    []model.User{},
				existingAccounts: []model.Account{},
				existingGroups:   []model.Group{},
				existingMembers:  []model.GroupMember{},
				want:             0,
				wantErr:          nil,
			},
			{
				name:             "無効なGroupID（空文字）",
				groupID:          model.GroupID(""),
				existingUsers:    []model.User{},
				existingAccounts: []model.Account{},
				existingGroups:   []model.Group{},
				existingMembers:  []model.GroupMember{},
				want:             0,
				wantErr:          persistence.ErrInvalidInput,
			},
			{
				name:             "無効なGroupID（不正なUUID）",
				groupID:          model.GroupID("invalid-uuid"),
				existingUsers:    []model.User{},
				existingAccounts: []model.Account{},
				existingGroups:   []model.Group{},
				existingMembers:  []model.GroupMember{},
				want:             0,
				wantErr:          persistence.ErrInvalidInput,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// テストデータを事前に作成
				for _, user := range tt.existingUsers {
					err := CreateTestUser(user.ID, user.Email)
					assert.NoError(t, err)
				}

				for _, account := range tt.existingAccounts {
					err := CreateTestAccount(account.ID, account.UserID, account.Name)
					assert.NoError(t, err)
				}

				for _, group := range tt.existingGroups {
					err := CreateTestGroup(group.ID, group.Name, group.Description)
					assert.NoError(t, err)
				}

				for _, member := range tt.existingMembers {
					err := CreateTestGroupMember(member.ID, member.GroupID, member.AccountID, member.Role, member.Status)
					assert.NoError(t, err)
				}

				got, err := groupMemberRepo.CountByGroupID(context.Background(), tt.groupID)

				if tt.wantErr != nil {
					assert.Error(t, err)
					assert.ErrorIs(t, err, tt.wantErr)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.want, got)
				}
			})
		}
	})
}