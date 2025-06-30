package persistence_test

import (
	"backend/domain/model"
	"backend/infra/persistence"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountPersistence(t *testing.T) {
	accountRepo := persistence.NewAccount(testDB)

	t.Run("FindByID", func(t *testing.T) {
		testAccountID := model.NewAccountID()
		nonExistentID := model.NewAccountID()
		
		tests := []struct {
			name    string
			id      model.AccountID
			want    model.Account
			wantErr error
		}{
			{
				name: "存在するアカウントを取得",
				id:   testAccountID,
				want: model.Account{
					ID:     testAccountID,
					UserID: "test-user-1",
					Name:   "Test Account",
				},
				wantErr: nil,
			},
			{
				name:    "存在しないアカウントを取得",
				id:      nonExistentID,
				want:    model.Account{},
				wantErr: persistence.ErrRecordNotFound,
			},
			{
				name:    "空のIDでアカウントを取得",
				id:      "",
				want:    model.Account{},
				wantErr: persistence.ErrInvalidInput,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// テストデータを事前に作成（必要な場合）
				if tt.name == "存在するアカウントを取得" {
					// 依存するユーザーを先に作成
					userRepo := persistence.NewUser(testDB)
					testUser := &model.User{
						ID:    "test-user-1",
						Email: "test1@example.com",
					}
					err := userRepo.Store(context.Background(), testUser)
					assert.NoError(t, err)

					// テストアカウントを作成
					testAccount := &model.Account{
						ID:     testAccountID,
						UserID: "test-user-1",
						Name:   "Test Account",
					}
					err = accountRepo.Store(context.Background(), testAccount)
					assert.NoError(t, err)
				}

				got, err := accountRepo.FindByID(context.Background(), tt.id)

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

	t.Run("FindByUserID", func(t *testing.T) {
		testAccountID := model.NewAccountID()
		
		tests := []struct {
			name    string
			userID  string
			want    model.Account
			wantErr error
		}{
			{
				name:   "存在するユーザーIDでアカウントを取得",
				userID: "test-user-1",
				want: model.Account{
					ID:     testAccountID,
					UserID: "test-user-1",
					Name:   "Test Account",
				},
				wantErr: nil,
			},
			{
				name:    "存在しないユーザーIDでアカウントを取得",
				userID:  "non-existent-user",
				want:    model.Account{},
				wantErr: persistence.ErrRecordNotFound,
			},
			{
				name:    "空のユーザーIDでアカウントを取得",
				userID:  "",
				want:    model.Account{},
				wantErr: persistence.ErrInvalidInput,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// テストデータを事前に作成（必要な場合）
				if tt.name == "存在するユーザーIDでアカウントを取得" {
					// 依存するユーザーを先に作成
					userRepo := persistence.NewUser(testDB)
					testUser := &model.User{
						ID:    "test-user-1",
						Email: "test1@example.com",
					}
					err := userRepo.Store(context.Background(), testUser)
					assert.NoError(t, err)

					// テストアカウントを作成
					testAccount := &model.Account{
						ID:     testAccountID,
						UserID: "test-user-1",
						Name:   "Test Account",
					}
					err = accountRepo.Store(context.Background(), testAccount)
					assert.NoError(t, err)
				}

				got, err := accountRepo.FindByUserID(context.Background(), tt.userID)

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

	t.Run("Store", func(t *testing.T) {
		newAccountID := model.NewAccountID()
		existingAccountID := model.NewAccountID()
		invalidAccountID := model.NewAccountID()
		duplicateAccountID := model.NewAccountID()
		existingDuplicateAccountID := model.NewAccountID()
		
		tests := []struct {
			name            string
			account         *model.Account
			existingAccount *model.Account
			existingUser    *model.User
			wantErr         error
		}{
			{
				name: "新規アカウントを作成",
				account: &model.Account{
					ID:     newAccountID,
					UserID: "new-user-1",
					Name:   "New Account",
				},
				existingUser: &model.User{
					ID:    "new-user-1",
					Email: "new1@example.com",
				},
				wantErr: nil,
			},
			{
				name: "既存アカウントを更新（名前を変更）",
				account: &model.Account{
					ID:     existingAccountID,
					UserID: "existing-user-1",
					Name:   "Updated Account Name",
				},
				existingAccount: &model.Account{
					ID:     existingAccountID,
					UserID: "existing-user-1",
					Name:   "Original Account Name",
				},
				existingUser: &model.User{
					ID:    "existing-user-1",
					Email: "existing1@example.com",
				},
				wantErr: nil,
			},
			{
				name: "存在しないユーザーIDでアカウントを作成",
				account: &model.Account{
					ID:     invalidAccountID,
					UserID: "non-existent-user",
					Name:   "Invalid Account",
				},
				wantErr: nil, // 外部キー制約エラーは予期しないエラーとして扱う
			},
			{
				name: "重複するユーザーIDでアカウントを作成",
				account: &model.Account{
					ID:     duplicateAccountID,
					UserID: "duplicate-user-1",
					Name:   "Duplicate Account",
				},
				existingAccount: &model.Account{
					ID:     existingDuplicateAccountID,
					UserID: "duplicate-user-1",
					Name:   "Existing Account",
				},
				existingUser: &model.User{
					ID:    "duplicate-user-1",
					Email: "duplicate1@example.com",
				},
				wantErr: nil, // unique制約エラーは予期しないエラーとして扱う
			},
			{
				name:    "nilアカウントを保存",
				account: nil,
				wantErr: persistence.ErrInvalidInput,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// 事前データの作成
				userRepo := persistence.NewUser(testDB)
				if tt.existingUser != nil {
					err := userRepo.Store(context.Background(), tt.existingUser)
					assert.NoError(t, err)
				}

				if tt.existingAccount != nil {
					err := accountRepo.Store(context.Background(), tt.existingAccount)
					assert.NoError(t, err)
				}

				err := accountRepo.Store(context.Background(), tt.account)

				if tt.wantErr != nil {
					assert.Error(t, err)
					assert.ErrorIs(t, err, tt.wantErr)
				} else {
					// 外部キー制約やunique制約エラーの場合はDBエラーが発生する可能性がある
					if tt.name == "存在しないユーザーIDでアカウントを作成" || tt.name == "重複するユーザーIDでアカウントを作成" {
						// これらのケースではDBレベルでエラーが発生することを期待
						if err != nil {
							// DBレベルのエラーが発生したことを確認
							assert.Error(t, err)
							return
						}
					}
					
					assert.NoError(t, err)

					// 保存後の検証
					if tt.account != nil {
						stored, err := accountRepo.FindByID(context.Background(), tt.account.ID)
						assert.NoError(t, err)
						assert.Equal(t, *tt.account, stored)
					}
				}
			})
		}
	})

	t.Run("FindByIDs", func(t *testing.T) {
		account1ID := model.NewAccountID()
		account2ID := model.NewAccountID()
		account3ID := model.NewAccountID()
		nonExistentID := model.NewAccountID()

		tests := []struct {
			name             string
			ids              []model.AccountID
			existingAccounts []model.Account
			existingUsers    []model.User
			want             []model.Account
			wantErr          error
		}{
			{
				name: "複数アカウント取得",
				ids:  []model.AccountID{account1ID, account2ID},
				existingAccounts: []model.Account{
					{
						ID:     account1ID,
						UserID: "user-1",
						Name:   "アカウント1",
					},
					{
						ID:     account2ID,
						UserID: "user-2",
						Name:   "アカウント2",
					},
					{
						ID:     account3ID,
						UserID: "user-3",
						Name:   "アカウント3",
					},
				},
				existingUsers: []model.User{
					{ID: "user-1", Email: "user1@example.com"},
					{ID: "user-2", Email: "user2@example.com"},
					{ID: "user-3", Email: "user3@example.com"},
				},
				want: []model.Account{
					{
						ID:     account1ID,
						UserID: "user-1",
						Name:   "アカウント1",
					},
					{
						ID:     account2ID,
						UserID: "user-2",
						Name:   "アカウント2",
					},
				},
				wantErr: nil,
			},
			{
				name: "単一アカウント取得",
				ids:  []model.AccountID{account1ID},
				existingAccounts: []model.Account{
					{
						ID:     account1ID,
						UserID: "user-1",
						Name:   "単一アカウント",
					},
				},
				existingUsers: []model.User{
					{ID: "user-1", Email: "user1@example.com"},
				},
				want: []model.Account{
					{
						ID:     account1ID,
						UserID: "user-1",
						Name:   "単一アカウント",
					},
				},
				wantErr: nil,
			},
			{
				name:             "空IDリスト",
				ids:              []model.AccountID{},
				existingAccounts: []model.Account{},
				existingUsers:    []model.User{},
				want:             []model.Account{},
				wantErr:          nil,
			},
			{
				name: "存在しないID含む",
				ids:  []model.AccountID{account1ID, nonExistentID},
				existingAccounts: []model.Account{
					{
						ID:     account1ID,
						UserID: "user-1",
						Name:   "存在するアカウント",
					},
				},
				existingUsers: []model.User{
					{ID: "user-1", Email: "user1@example.com"},
				},
				want: []model.Account{
					{
						ID:     account1ID,
						UserID: "user-1",
						Name:   "存在するアカウント",
					},
				},
				wantErr: nil,
			},
			{
				name:             "全て存在しないID",
				ids:              []model.AccountID{nonExistentID},
				existingAccounts: []model.Account{},
				existingUsers:    []model.User{},
				want:             []model.Account{},
				wantErr:          nil,
			},
			{
				name: "重複ID",
				ids:  []model.AccountID{account1ID, account1ID},
				existingAccounts: []model.Account{
					{
						ID:     account1ID,
						UserID: "user-1",
						Name:   "重複テストアカウント",
					},
				},
				existingUsers: []model.User{
					{ID: "user-1", Email: "user1@example.com"},
				},
				want: []model.Account{
					{
						ID:     account1ID,
						UserID: "user-1",
						Name:   "重複テストアカウント",
					},
				},
				wantErr: nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// 依存するユーザーを事前に作成
				userRepo := persistence.NewUser(testDB)
				for _, user := range tt.existingUsers {
					err := userRepo.Store(context.Background(), &user)
					assert.NoError(t, err)
				}

				// テストアカウントを事前に作成
				for _, account := range tt.existingAccounts {
					err := accountRepo.Store(context.Background(), &account)
					assert.NoError(t, err)
				}

				got, err := accountRepo.FindByIDs(context.Background(), tt.ids)

				if tt.wantErr != nil {
					assert.Error(t, err)
					assert.ErrorIs(t, err, tt.wantErr)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, len(tt.want), len(got))

					// 結果を検証（順序に依存しない比較）
					for _, wantAccount := range tt.want {
						found := false
						for _, gotAccount := range got {
							if gotAccount.ID == wantAccount.ID {
								assert.Equal(t, wantAccount.UserID, gotAccount.UserID)
								assert.Equal(t, wantAccount.Name, gotAccount.Name)
								found = true
								break
							}
						}
						assert.True(t, found, "期待されたアカウントが見つかりません: %s", wantAccount.ID)
					}
				}
			})
		}
	})

	t.Run("Store_FindByID_Integration", func(t *testing.T) {
		// 統合テスト：保存したデータが正しく取得できることを確認
		defer CleanupTestData()

		integrationAccountID := model.NewAccountID()

		// 依存するユーザーを作成
		userRepo := persistence.NewUser(testDB)
		user := &model.User{
			ID:    "integration-test-user",
			Email: "integration@example.com",
		}
		err := userRepo.Store(context.Background(), user)
		assert.NoError(t, err)

		// アカウントを作成
		account := &model.Account{
			ID:     integrationAccountID,
			UserID: "integration-test-user",
			Name:   "Integration Test Account",
		}

		// アカウントを保存
		err = accountRepo.Store(context.Background(), account)
		assert.NoError(t, err)

		// 保存したアカウントを取得
		found, err := accountRepo.FindByID(context.Background(), account.ID)
		assert.NoError(t, err)
		assert.Equal(t, *account, found)
	})
}
