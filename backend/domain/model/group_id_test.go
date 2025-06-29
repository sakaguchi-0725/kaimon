package model

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewGroupID(t *testing.T) {
	id := NewGroupID()
	
	// UUIDとしてパースできることを確認
	_, err := uuid.Parse(string(id))
	if err != nil {
		t.Errorf("NewGroupIDが有効なUUIDを生成していません: %v", err)
	}
	
	// 複数回呼び出しで異なるIDが生成されることを確認
	id2 := NewGroupID()
	if id == id2 {
		t.Error("NewGroupIDが同じIDを生成しています")
	}
}

func TestParseGroupID(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "正常系：有効なUUID",
			input:   "550e8400-e29b-41d4-a716-446655440000",
			wantErr: false,
		},
		{
			name:    "異常系：無効なUUID形式",
			input:   "invalid-uuid",
			wantErr: true,
		},
		{
			name:    "異常系：空文字",
			input:   "",
			wantErr: true,
		},
		{
			name:    "異常系：UUIDに似ているが無効",
			input:   "550e8400-e29b-41d4-a716-44665544000g",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := ParseGroupID(tt.input)

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

			// 正常系の場合、IDが正しく設定されていることを確認
			if id.String() != tt.input {
				t.Errorf("期待されたID: %s, 実際: %s", tt.input, id.String())
			}
		})
	}
}

func TestGroupID_String(t *testing.T) {
	expected := "550e8400-e29b-41d4-a716-446655440000"
	id := GroupID(expected)
	
	if id.String() != expected {
		t.Errorf("期待された文字列: %s, 実際: %s", expected, id.String())
	}
}