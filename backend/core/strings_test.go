package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// RemovePrefixのテスト関数
func TestRemovePrefix(t *testing.T) {
	// テストケース定義
	tests := []struct {
		name     string
		input    string
		prefix   string
		expected string
		wantErr  bool
	}{
		{
			name:     "正常なBearerトークン",
			input:    "Bearer token123",
			prefix:   "Bearer ",
			expected: "token123",
			wantErr:  false,
		},
		{
			name:     "Basic認証の場合",
			input:    "Basic dXNlcjpwYXNz",
			prefix:   "Basic ",
			expected: "dXNlcjpwYXNz",
			wantErr:  false,
		},
		{
			name:    "プレフィックスが一致しない場合",
			input:   "Basic token123",
			prefix:  "Bearer ",
			wantErr: true,
		},
		{
			name:    "入力文字列が空の場合",
			input:   "",
			prefix:  "Bearer ",
			wantErr: true,
		},
		{
			name:    "プレフィックスが空の場合",
			input:   "Bearer token123",
			prefix:  "",
			wantErr: true,
		},
		{
			name:    "プレフィックス除去後が空文字になる場合",
			input:   "Bearer ",
			prefix:  "Bearer ",
			wantErr: true,
		},
		{
			name:     "プレフィックス除去後に空白のみの場合",
			input:    "Bearer   ",
			prefix:   "Bearer ",
			expected: "  ",
			wantErr:  false,
		},
		{
			name:     "カスタムプレフィックスの場合",
			input:    "X-API-Key: abc123",
			prefix:   "X-API-Key: ",
			expected: "abc123",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := RemovePrefix(tt.input, tt.prefix)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
