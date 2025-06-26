package db

import (
	"backend/core"
	"fmt"
	"path/filepath"
	"runtime"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

// RunMigrations はsql-migrateを使用してデータベースマイグレーションを実行する
func RunMigrations(conn *Conn, cfg core.DBConfig) error {
	// GORMの*sql.DBインスタンスを取得
	sqlDB, err := conn.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB instance from GORM: %w", err)
	}

	// データベース接続をテスト
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("database connection test failed: %w", err)
	}

	// プロジェクトルートからmigrationsディレクトリの絶対パスを取得
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "..", "..")
	migrationsPath := filepath.Join(projectRoot, "migrations")

	// マイグレーション設定
	migrations := &migrate.FileMigrationSource{
		Dir: migrationsPath,
	}

	// マイグレーションを実行
	n, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	fmt.Printf("Applied %d migrations\n", n)
	return nil
}

// RollbackMigrations はsql-migrateを使用してデータベースロールバックを実行する
func RollbackMigrations(conn *Conn, cfg core.DBConfig, steps int) error {
	// GORMの*sql.DBインスタンスを取得
	sqlDB, err := conn.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB instance from GORM: %w", err)
	}

	// データベース接続をテスト
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("database connection test failed: %w", err)
	}

	// プロジェクトルートからmigrationsディレクトリの絶対パスを取得
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "..", "..")
	migrationsPath := filepath.Join(projectRoot, "migrations")

	// マイグレーション設定
	migrations := &migrate.FileMigrationSource{
		Dir: migrationsPath,
	}

	// ロールバックを実行
	n, err := migrate.ExecMax(sqlDB, "postgres", migrations, migrate.Down, steps)
	if err != nil {
		return fmt.Errorf("failed to rollback migrations: %w", err)
	}

	fmt.Printf("Rolled back %d migrations\n", n)
	return nil
}
