package model

import "testing"

func TestMemberStatus_String(t *testing.T) {
	tests := []struct {
		name     string
		status   MemberStatus
		expected string
	}{
		{
			name:     "アクティブステータス",
			status:   MemberStatusActive,
			expected: "active",
		},
		{
			name:     "ペンディングステータス",
			status:   MemberStatusPending,
			expected: "pending",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.status.String()
			if result != tt.expected {
				t.Errorf("期待された文字列: %s, 実際: %s", tt.expected, result)
			}
		})
	}
}

func TestMemberStatusConstants(t *testing.T) {
	// 定数の値が期待通りであることを確認
	if MemberStatusActive != "active" {
		t.Errorf("MemberStatusActiveの値が正しくありません: %s", MemberStatusActive)
	}
	
	if MemberStatusPending != "pending" {
		t.Errorf("MemberStatusPendingの値が正しくありません: %s", MemberStatusPending)
	}
}