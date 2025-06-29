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
	layout := "2006-01-02 15:04:05"
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
	layout := "2006-01-02 15:04:05"
	value := "invalid-date"
	
	_, err := ParseJST(layout, value)
	if err == nil {
		t.Error("Expected error for invalid date format, got nil")
	}
}

func TestFormatJST(t *testing.T) {
	// UTC時刻を作成
	utc := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	layout := "2006-01-02 15:04:05"
	
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