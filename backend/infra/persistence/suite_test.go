package persistence_test

import (
	"backend/core"
	"backend/infra/db"
	"fmt"
	"log"
	"os"
	"testing"
)

var testDB *db.Conn

func TestMain(m *testing.M) {
	testDBConfig := core.LoadTestDBConfig()

	// テストDBに接続
	var err error
	testDB, err = db.New(testDBConfig)
	if err != nil {
		log.Fatalf("テストDB接続に失敗: %v", err)
	}

	// マイグレーションを実行
	if err := db.RunMigrations(testDB, testDBConfig); err != nil {
		log.Fatalf("テスト用マイグレーション実行に失敗: %v", err)
	}

	// テスト実行
	code := m.Run()

	// テスト終了後のクリーンアップ
	if err := cleanupTestDB(); err != nil {
		log.Printf("テストDB クリーンアップに失敗: %v", err)
	}

	// DB接続をクローズ
	sqlDB, err := testDB.DB.DB()
	if err == nil {
		sqlDB.Close()
	}

	os.Exit(code)
}

// CleanupTestData はテストデータをクリアする
func CleanupTestData() error {
	return cleanupTestDB()
}

// cleanupTestDB はテスト用データベースの全データを削除する
func cleanupTestDB() error {
	if testDB == nil {
		return fmt.Errorf("テストDB接続が初期化されていません")
	}

	// 全テーブルのデータを削除（外部キーの依存関係を考慮した順序）
	tables := []string{"accounts", "users"}

	for _, table := range tables {
		if err := testDB.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table)).Error; err != nil {
			return fmt.Errorf("テーブル %s のクリアに失敗: %w", table, err)
		}
	}

	return nil
}
