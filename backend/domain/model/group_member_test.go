package model

import (
	"testing"
	"time"
)

func TestNewGroupMember(t *testing.T) {
	tests := []struct {
		name      string
		id        GroupMemberID
		groupID   GroupID
		accountID AccountID
		wantErr   bool
	}{
		{
			name:      "正常系：有効なグループメンバーを作成",
			id:        NewGroupMemberID(),
			groupID:   NewGroupID(),
			accountID: NewAccountID(),
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			member, err := NewGroupMember(tt.id, tt.groupID, tt.accountID)

			if tt.wantErr {
				if err == nil {
					t.Error("エラーが期待されましたが、エラーが発生しませんでした")
				}
				return
			}

			if err != nil {
				t.Errorf("予期しないエラーが発生しました: %v", err)
				return
			}

			// 正常系の検証
			if member.ID != tt.id {
				t.Errorf("期待されたID: %s, 実際: %s", tt.id, member.ID)
			}
			if member.GroupID != tt.groupID {
				t.Errorf("期待されたGroupID: %s, 実際: %s", tt.groupID, member.GroupID)
			}
			if member.AccountID != tt.accountID {
				t.Errorf("期待されたAccountID: %s, 実際: %s", tt.accountID, member.AccountID)
			}

			// デフォルト値の検証
			if member.Role != MemberRoleAdmin {
				t.Errorf("期待されたRole: %s, 実際: %s", MemberRoleAdmin, member.Role)
			}
			if member.Status != MemberStatusActive {
				t.Errorf("期待されたStatus: %s, 実際: %s", MemberStatusActive, member.Status)
			}

			// JoinedAtがJSTで設定されていることを確認
			now := time.Now()
			if member.JoinedAt.After(now) || member.JoinedAt.Before(now.Add(-time.Second)) {
				t.Errorf("JoinedAtが現在時刻付近に設定されていません: %v", member.JoinedAt)
			}

			// JSTタイムゾーンであることを確認
			zone, offset := member.JoinedAt.Zone()
			if zone != "Asia/Tokyo" || offset != 9*60*60 {
				t.Errorf("JoinedAtがJSTタイムゾーンではありません. zone=%s, offset=%d", zone, offset)
			}
		})
	}
}