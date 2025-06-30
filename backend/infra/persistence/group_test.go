package persistence_test

import (
	"backend/domain/model"
	"backend/infra/persistence"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupPersistence(t *testing.T) {
	groupRepo := persistence.NewGroup(testDB)

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
					},
					{
						ID:          group2ID,
						Name:        "テストグループ2",
						Description: "テスト用のグループ2です",
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
					// CreatedAt, UpdatedAtは自動生成なので検証しない
				}
			})
		}
	})
}