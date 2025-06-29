package model

import (
	"strings"
	"testing"
	"time"
)

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func TestNewGroup(t *testing.T) {
	tests := []struct {
		name        string
		id          GroupID
		groupName   string
		description string
		wantErr     bool
		errMsg      string
	}{
		{
			name:        "正常系：有効なグループを作成",
			id:          NewGroupID(),
			groupName:   "テストグループ",
			description: "テスト用のグループです",
			wantErr:     false,
		},
		{
			name:        "正常系：説明なしでグループを作成",
			id:          NewGroupID(),
			groupName:   "説明なしグループ",
			description: "",
			wantErr:     false,
		},
		{
			name:        "異常系：グループ名が空文字",
			id:          NewGroupID(),
			groupName:   "",
			description: "説明があってもグループ名が空",
			wantErr:     true,
			errMsg:      "name is required at",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group, err := NewGroup(tt.id, tt.groupName, tt.description)

			if tt.wantErr {
				if err == nil {
					t.Error("エラーが期待されましたが、エラーが発生しませんでした")
					return
				}
				if !contains(err.Error(), tt.errMsg) {
					t.Errorf("エラーメッセージに期待される文字列が含まれていません. 期待: %s, 実際: %s", tt.errMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("予期しないエラーが発生しました: %v", err)
				return
			}

			// 正常系の検証
			if group.ID != tt.id {
				t.Errorf("期待されたID: %s, 実際: %s", tt.id, group.ID)
			}
			if group.Name != tt.groupName {
				t.Errorf("期待されたグループ名: %s, 実際: %s", tt.groupName, group.Name)
			}
			if group.Description != tt.description {
				t.Errorf("期待された説明: %s, 実際: %s", tt.description, group.Description)
			}

			// CreatedAtがJSTで設定されていることを確認
			now := time.Now()
			if group.CreatedAt.After(now) || group.CreatedAt.Before(now.Add(-time.Second)) {
				t.Errorf("CreatedAtが現在時刻付近に設定されていません: %v", group.CreatedAt)
			}

			// JSTタイムゾーンであることを確認
			zone, offset := group.CreatedAt.Zone()
			if zone != "Asia/Tokyo" || offset != 9*60*60 {
				t.Errorf("CreatedAtがJSTタイムゾーンではありません. zone=%s, offset=%d", zone, offset)
			}
		})
	}
}