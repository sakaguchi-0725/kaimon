package persistence_test

import (
	"backend/domain/model"
	"backend/infra/persistence"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserPersistence(t *testing.T) {
	// テストDB接続を取得
	userRepo := persistence.NewUser(testDB)

	t.Run("FindByID", func(t *testing.T) {
		tests := []struct {
			name    string
			id      string
			want    model.User
			wantErr bool
		}{
			{
				name:    "存在するユーザーを取得",
				id:      "test-user-1",
				want:    model.User{ID: "test-user-1", Email: "test1@example.com"},
				wantErr: false,
			},
			{
				name:    "存在しないユーザーを取得",
				id:      "non-existent-user",
				want:    model.User{},
				wantErr: true,
			},
			{
				name:    "空のIDでユーザーを取得",
				id:      "",
				want:    model.User{},
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// テストデータを事前に作成（必要な場合）
				if tt.name == "存在するユーザーを取得" {
					testUser := &model.User{
						ID:    "test-user-1",
						Email: "test1@example.com",
					}
					err := userRepo.Store(context.Background(), testUser)
					assert.NoError(t, err)
				}

				got, err := userRepo.FindByID(context.Background(), tt.id)

				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.want, got)
				}
			})
		}
	})

	t.Run("Store", func(t *testing.T) {
		tests := []struct {
			name         string
			user         *model.User
			existingUser *model.User // 事前に作成するユーザー
			wantErr      bool
		}{
			{
				name: "新規ユーザーを作成",
				user: &model.User{
					ID:    "new-user-1",
					Email: "new1@example.com",
				},
				existingUser: nil,
				wantErr:      false,
			},
			{
				name: "既存ユーザーを更新（Emailを変更）",
				user: &model.User{
					ID:    "existing-user-1",
					Email: "updated@example.com",
				},
				existingUser: &model.User{
					ID:    "existing-user-1",
					Email: "original@example.com",
				},
				wantErr: false,
			},
			{
				name: "重複するEmailで新規ユーザーを作成",
				user: &model.User{
					ID:    "new-user-2",
					Email: "duplicate@example.com",
				},
				existingUser: &model.User{
					ID:    "existing-user-2",
					Email: "duplicate@example.com",
				},
				wantErr: true,
			},
			{
				name:         "nilユーザーを保存",
				user:         nil,
				existingUser: nil,
				wantErr:      true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// 事前データの作成
				if tt.existingUser != nil {
					err := userRepo.Store(context.Background(), tt.existingUser)
					assert.NoError(t, err)
				}

				err := userRepo.Store(context.Background(), tt.user)

				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)

					// 保存後の検証
					if tt.user != nil {
						stored, err := userRepo.FindByID(context.Background(), tt.user.ID)
						assert.NoError(t, err)
						assert.Equal(t, *tt.user, stored)
					}
				}
			})
		}
	})

	t.Run("Store_FindByID_Integration", func(t *testing.T) {
		// 統合テスト：保存したデータが正しく取得できることを確認
		defer CleanupTestData()

		user := &model.User{
			ID:    "integration-test-user",
			Email: "integration@example.com",
		}

		// ユーザーを保存
		err := userRepo.Store(context.Background(), user)
		assert.NoError(t, err)

		// 保存したユーザーを取得
		found, err := userRepo.FindByID(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Equal(t, *user, found)
	})
}
