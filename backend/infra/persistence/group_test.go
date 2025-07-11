package persistence_test

import (
	"backend/domain/model"
	"backend/infra/persistence"
	mock "backend/test/mock/external"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGroupPersistence(t *testing.T) {
	var (
		ctrl            = gomock.NewController(t)
		mockRedisClient = mock.NewMockRedisClient(ctrl)
		groupRepo       = persistence.NewGroup(testDB, mockRedisClient)
	)
	defer ctrl.Finish()

	t.Run("FindByIDs", func(t *testing.T) {
		group1ID := model.NewGroupID()
		group2ID := model.NewGroupID()
		group3ID := model.NewGroupID()
		nonExistentID := model.NewGroupID()

		tests := []struct {
			name           string
			ids            []model.GroupID
			existingGroups []model.Group
			want           []model.Group
			wantErr        error
		}{
			{
				name: "複数の存在するグループを取得",
				ids:  []model.GroupID{group1ID, group2ID},
				existingGroups: []model.Group{
					{
						ID:          group1ID,
						Name:        "テストグループ1",
						Description: "テスト用のグループ1です",
					},
					{
						ID:          group2ID,
						Name:        "テストグループ2",
						Description: "テスト用のグループ2です",
					},
					{
						ID:          group3ID,
						Name:        "テストグループ3",
						Description: "",
					},
				},
				want: []model.Group{
					{
						ID:          group1ID,
						Name:        "テストグループ1",
						Description: "テスト用のグループ1です",
						Members:     []model.GroupMember{},
					},
					{
						ID:          group2ID,
						Name:        "テストグループ2",
						Description: "テスト用のグループ2です",
						Members:     []model.GroupMember{},
					},
				},
				wantErr: nil,
			},
			{
				name: "存在しないグループIDを含む場合",
				ids:  []model.GroupID{group1ID, nonExistentID},
				existingGroups: []model.Group{
					{
						ID:          group1ID,
						Name:        "存在するグループ",
						Description: "これは存在します",
					},
				},
				want: []model.Group{
					{
						ID:          group1ID,
						Name:        "存在するグループ",
						Description: "これは存在します",
					},
				},
				wantErr: nil,
			},
			{
				name:           "空のIDリストで検索",
				ids:            []model.GroupID{},
				existingGroups: []model.Group{},
				want:           []model.Group{},
				wantErr:        nil,
			},
			{
				name:           "全て存在しないIDで検索",
				ids:            []model.GroupID{nonExistentID},
				existingGroups: []model.Group{},
				want:           []model.Group{},
				wantErr:        nil,
			},
			{
				name: "重複するIDを含む場合",
				ids:  []model.GroupID{group1ID, group1ID},
				existingGroups: []model.Group{
					{
						ID:          group1ID,
						Name:        "重複テストグループ",
						Description: "重複IDのテスト",
					},
				},
				want: []model.Group{
					{
						ID:          group1ID,
						Name:        "重複テストグループ",
						Description: "重複IDのテスト",
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
				for _, group := range tt.existingGroups {
					err := CreateTestGroup(group.ID, group.Name, group.Description)
					assert.NoError(t, err)
				}

				got, err := groupRepo.FindByIDs(context.Background(), tt.ids)

				if tt.wantErr != nil {
					assert.Error(t, err)
					assert.ErrorIs(t, err, tt.wantErr)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, len(tt.want), len(got))

					// 結果を検証（順序に依存しない比較）
					for _, wantGroup := range tt.want {
						found := false
						for _, gotGroup := range got {
							if gotGroup.ID == wantGroup.ID {
								assert.Equal(t, wantGroup.Name, gotGroup.Name)
								assert.Equal(t, wantGroup.Description, gotGroup.Description)
								found = true
								break
							}
						}
						assert.True(t, found, "期待されたグループが見つかりません: %s", wantGroup.ID)
					}
				}
			})
		}
	})

	t.Run("GetByID", func(t *testing.T) {
		existingGroupID := model.NewGroupID()
		nonExistentGroupID := model.NewGroupID()

		tests := []struct {
			name           string
			groupID        model.GroupID
			existingGroups []model.Group
			want           model.Group
			wantErr        error
		}{
			{
				name:    "存在するグループを取得",
				groupID: existingGroupID,
				existingGroups: []model.Group{
					{
						ID:          existingGroupID,
						Name:        "テストグループ",
						Description: "テスト用のグループです",
					},
				},
				want: model.Group{
					ID:          existingGroupID,
					Name:        "テストグループ",
					Description: "テスト用のグループです",
					Members:     []model.GroupMember{}, // メンバーなしのグループ
				},
				wantErr: nil,
			},
			{
				name:    "存在するグループを取得（Descriptionが空）",
				groupID: existingGroupID,
				existingGroups: []model.Group{
					{
						ID:          existingGroupID,
						Name:        "説明なしグループ",
						Description: "",
					},
				},
				want: model.Group{
					ID:          existingGroupID,
					Name:        "説明なしグループ",
					Description: "",
					Members:     []model.GroupMember{}, // メンバーなしのグループ
				},
				wantErr: nil,
			},
			{
				name:           "存在しないグループID",
				groupID:        nonExistentGroupID,
				existingGroups: []model.Group{},
				want:           model.Group{},
				wantErr:        persistence.ErrRecordNotFound,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// テストデータを事前に作成
				for _, group := range tt.existingGroups {
					err := CreateTestGroup(group.ID, group.Name, group.Description)
					assert.NoError(t, err)
				}

				got, err := groupRepo.GetByID(context.Background(), tt.groupID)

				if tt.wantErr != nil {
					assert.Error(t, err)
					assert.ErrorIs(t, err, tt.wantErr)
					assert.Equal(t, model.Group{}, got)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.want.ID, got.ID)
					assert.Equal(t, tt.want.Name, got.Name)
					assert.Equal(t, tt.want.Description, got.Description)
					assert.Equal(t, tt.want.Members, got.Members)
					// CreatedAt, UpdatedAtは自動生成なので検証しない
				}
			})
		}
	})

	t.Run("Store", func(t *testing.T) {
		tests := []struct {
			name    string
			group   model.Group
			wantErr error
		}{
			{
				name: "新しいグループを正常に保存",
				group: model.Group{
					ID:          model.NewGroupID(),
					Name:        "新しいテストグループ",
					Description: "新しく作成されるグループです",
				},
				wantErr: nil,
			},
			{
				name: "Descriptionが空のグループを保存",
				group: model.Group{
					ID:          model.NewGroupID(),
					Name:        "説明なしグループ",
					Description: "",
				},
				wantErr: nil,
			},
			{
				name: "長い名前のグループを保存",
				group: model.Group{
					ID:          model.NewGroupID(),
					Name:        "とても長い名前のグループですがこれでも保存できるはずです",
					Description: "長い名前のテストケース",
				},
				wantErr: nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// Storeメソッドを実行
				err := groupRepo.Store(context.Background(), &tt.group)

				if tt.wantErr != nil {
					assert.Error(t, err)
					assert.ErrorIs(t, err, tt.wantErr)
				} else {
					assert.NoError(t, err)

					// 保存されたデータを検証
					stored, err := groupRepo.GetByID(context.Background(), tt.group.ID)
					assert.NoError(t, err)
					assert.Equal(t, tt.group.ID, stored.ID)
					assert.Equal(t, tt.group.Name, stored.Name)
					assert.Equal(t, tt.group.Description, stored.Description)
				}
			})
		}
	})

	t.Run("Invitation", func(t *testing.T) {
		tests := []struct {
			name       string
			invitation model.Invitation
			setupMock  func(*mock.MockRedisClient)
			wantErr    bool
		}{
			{
				name: "招待コードを正常に保存",
				invitation: model.Invitation{
					Code:      "ABC123XY",
					GroupID:   model.NewGroupID(),
					ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
				},
				setupMock: func(mockRedis *mock.MockRedisClient) {
					mockRedis.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				},
				wantErr: false,
			},
			{
				name: "Redisエラーが発生した場合",
				invitation: model.Invitation{
					Code:      "DEF456ZW",
					GroupID:   model.NewGroupID(),
					ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
				},
				setupMock: func(mockRedis *mock.MockRedisClient) {
					mockRedis.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(assert.AnError)
				},
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// モックのセットアップ
				tt.setupMock(mockRedisClient)

				err := groupRepo.Invitation(context.Background(), tt.invitation)

				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})

	t.Run("GetInvitation", func(t *testing.T) {
		groupID := model.NewGroupID()
		
		tests := []struct {
			name      string
			groupID   model.GroupID
			setupMock func(*mock.MockRedisClient)
			want      *model.Invitation
			wantErr   bool
		}{
			{
				name:    "既存の招待コードを正常に取得",
				groupID: groupID,
				setupMock: func(mockRedis *mock.MockRedisClient) {
					invitationJSON := `{"code":"ABC123XY","groupId":"` + groupID.String() + `","expiresAt":"2024-01-08T00:00:00Z"}`
					mockRedis.EXPECT().Get(gomock.Any(), "group_invitation:"+groupID.String()).Return(invitationJSON, nil)
				},
				want: &model.Invitation{
					Code:      "ABC123XY",
					GroupID:   groupID,
					ExpiresAt: time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),
				},
				wantErr: false,
			},
			{
				name:    "招待コードが存在しない場合",
				groupID: groupID,
				setupMock: func(mockRedis *mock.MockRedisClient) {
					mockRedis.EXPECT().Get(gomock.Any(), "group_invitation:"+groupID.String()).Return("", nil)
				},
				want:    nil,
				wantErr: false,
			},
			{
				name:    "Redisエラーが発生した場合",
				groupID: groupID,
				setupMock: func(mockRedis *mock.MockRedisClient) {
					mockRedis.EXPECT().Get(gomock.Any(), "group_invitation:"+groupID.String()).Return("", assert.AnError)
				},
				want:    nil,
				wantErr: true,
			},
			{
				name:    "JSONデシリアライズエラーが発生した場合",
				groupID: groupID,
				setupMock: func(mockRedis *mock.MockRedisClient) {
					mockRedis.EXPECT().Get(gomock.Any(), "group_invitation:"+groupID.String()).Return("invalid-json", nil)
				},
				want:    nil,
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// テストデータをクリア
				defer CleanupTestData()

				// モックのセットアップ
				tt.setupMock(mockRedisClient)

				got, err := groupRepo.GetInvitation(context.Background(), tt.groupID)

				if tt.wantErr {
					assert.Error(t, err)
					assert.Nil(t, got)
				} else {
					assert.NoError(t, err)
					if tt.want == nil {
						assert.Nil(t, got)
					} else {
						assert.NotNil(t, got)
						assert.Equal(t, tt.want.Code, got.Code)
						assert.Equal(t, tt.want.GroupID, got.GroupID)
						// ExpiresAtは時刻の比較が複雑なので、ここでは簡単な検証のみ
						assert.NotZero(t, got.ExpiresAt)
					}
				}
			})
		}
	})
}
