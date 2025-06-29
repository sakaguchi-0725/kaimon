package model

import "testing"

func TestMemberRole_String(t *testing.T) {
	tests := []struct {
		name     string
		role     MemberRole
		expected string
	}{
		{
			name:     "管理者ロール",
			role:     MemberRoleAdmin,
			expected: "admin",
		},
		{
			name:     "メンバーロール",
			role:     MemberRoleMember,
			expected: "member",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.role.String()
			if result != tt.expected {
				t.Errorf("期待された文字列: %s, 実際: %s", tt.expected, result)
			}
		})
	}
}

func TestMemberRoleConstants(t *testing.T) {
	// 定数の値が期待通りであることを確認
	if MemberRoleAdmin != "admin" {
		t.Errorf("MemberRoleAdminの値が正しくありません: %s", MemberRoleAdmin)
	}
	
	if MemberRoleMember != "member" {
		t.Errorf("MemberRoleMemberの値が正しくありません: %s", MemberRoleMember)
	}
}