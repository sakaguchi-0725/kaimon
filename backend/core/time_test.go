package core

import (
	"testing"
	"time"
)

func TestNowJST(t *testing.T) {
	now := NowJST()

	// JSTタイムゾーンであることを確認
	zone, offset := now.Zone()
	if zone != "Asia/Tokyo" || offset != 9*60*60 {
		t.Errorf("Expected JST timezone, got zone=%s, offset=%d", zone, offset)
	}
}

func TestToJST(t *testing.T) {
	// UTC時刻を作成
	utc := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)

	// JSTに変換
	jst := ToJST(utc)

	// JSTタイムゾーンであることを確認
	zone, offset := jst.Zone()
	if zone != "Asia/Tokyo" || offset != 9*60*60 {
		t.Errorf("Expected JST timezone, got zone=%s, offset=%d", zone, offset)
	}

	// 時刻が正しく変換されていることを確認（UTC 12:00 = JST 21:00）
	expected := time.Date(2024, 1, 1, 21, 0, 0, 0, JST)
	if !jst.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, jst)
	}
}

func TestParseJST(t *testing.T) {
	layout := LayoutDateTime
	value := "2024-01-01 12:00:00"

	parsed, err := ParseJST(layout, value)
	if err != nil {
		t.Fatalf("ParseJST failed: %v", err)
	}

	// JSTタイムゾーンであることを確認
	zone, offset := parsed.Zone()
	if zone != "Asia/Tokyo" || offset != 9*60*60 {
		t.Errorf("Expected JST timezone, got zone=%s, offset=%d", zone, offset)
	}

	// パースされた時刻が正しいことを確認
	expected := time.Date(2024, 1, 1, 12, 0, 0, 0, JST)
	if !parsed.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, parsed)
	}
}

func TestParseJST_InvalidFormat(t *testing.T) {
	layout := LayoutDateTime
	value := "invalid-date"

	_, err := ParseJST(layout, value)
	if err == nil {
		t.Error("Expected error for invalid date format, got nil")
	}
}

func TestFormatJST(t *testing.T) {
	// UTC時刻を作成
	utc := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	layout := LayoutDateTime

	// JSTでフォーマット
	formatted := FormatJST(utc, layout)

	// JSTに変換された時刻でフォーマットされていることを確認（UTC 12:00 = JST 21:00）
	expected := "2024-01-01 21:00:00"
	if formatted != expected {
		t.Errorf("Expected %s, got %s", expected, formatted)
	}
}

func TestJSTConstant(t *testing.T) {
	// オフセットが9時間（32400秒）であることを確認
	_, jstOffset := time.Now().In(JST).Zone()
	if jstOffset != 9*60*60 {
		t.Errorf("Expected JST offset 32400 seconds, got %d", jstOffset)
	}
}

func TestFormatWithLayout(t *testing.T) {
	// テスト用の時刻を作成
	testTime := time.Date(2024, 1, 1, 12, 30, 45, 0, JST)

	tests := []struct {
		name     string
		layout   string
		expected string
	}{
		{
			name:     "ISO 8601フォーマット",
			layout:   LayoutISO8601,
			expected: "2024-01-01T12:30:45+09:00",
		},
		{
			name:     "短い日付フォーマット",
			layout:   LayoutDateSlash,
			expected: "2024/01/01",
		},
		{
			name:     "時刻のみフォーマット",
			layout:   LayoutTimeOnly,
			expected: "12:30:45",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatWithLayout(testTime, tt.layout)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
