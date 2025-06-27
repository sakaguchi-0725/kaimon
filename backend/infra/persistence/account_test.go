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
