package persistence_test

import (
	"backend/core"
	"backend/domain/model"
	"backend/infra/db"
	"backend/infra/dto"
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

	// マイグレーション履歴テーブルをリセット（エラーは無視）
	_ = testDB.Exec("DROP TABLE IF EXISTS gorp_migrations")

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
	tables := []string{"members", "groups", "accounts", "users"}

	for _, table := range tables {
		if err := testDB.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table)).Error; err != nil {
			return fmt.Errorf("テーブル %s のクリアに失敗: %w", table, err)
		}
	}

	return nil
}

// CreateTestUser はテスト用のユーザーを作成する
func CreateTestUser(id, email string) error {
	user := model.User{
		ID:    id,
		Email: email,
	}
	userDto := dto.ToUserDto(user)
	return testDB.Create(&userDto).Error
}

// CreateTestAccount はテスト用のアカウントを作成する
func CreateTestAccount(id model.AccountID, userID, name string) error {
	account := model.Account{
		ID:     id,
		UserID: userID,
		Name:   name,
	}
	accountDto := dto.ToAccountDto(account)
	return testDB.Create(&accountDto).Error
}

// CreateTestGroup はテスト用のグループを作成する
func CreateTestGroup(id model.GroupID, name, description string) error {
	group := model.Group{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   core.NowJST(),
	}
	groupDto := dto.ToGroupDto(group)
	return testDB.Create(&groupDto).Error
}

// CreateTestGroupMember はテスト用のグループメンバーを作成する
func CreateTestGroupMember(id model.GroupMemberID, groupID model.GroupID, accountID model.AccountID, role model.MemberRole, status model.MemberStatus) error {
	member := model.GroupMember{
		ID:        id,
		GroupID:   groupID,
		AccountID: accountID,
		Role:      role,
		Status:    status,
		JoinedAt:  core.NowJST(),
	}
	memberDto := dto.ToGroupMemberDto(member)
	return testDB.Create(&memberDto).Error
}

// CreateTestAccountWithUser はユーザーとアカウントをセットで作成する便利メソッド
func CreateTestAccountWithUser(userID, email string, accountID model.AccountID, accountName string) error {
	// 先にユーザーを作成
	if err := CreateTestUser(userID, email); err != nil {
		return fmt.Errorf("ユーザー作成に失敗: %w", err)
	}

	// 次にアカウントを作成
	if err := CreateTestAccount(accountID, userID, accountName); err != nil {
		return fmt.Errorf("アカウント作成に失敗: %w", err)
	}

	return nil
}
