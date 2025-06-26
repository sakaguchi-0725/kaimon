package core

import (
	"fmt"
	"os"
	"strings"
)

type (
	Config struct {
		Port string
		Env  string

		Redis RedisConfig
		DB    DBConfig
	}

	DBConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	RedisConfig struct {
		Host string
		Port string
	}
)

// DSN はデータベース接続文字列を返します
func (c DBConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		c.Host, c.Port, c.Name, c.User, c.Password)
}

// 環境変数からConfigを読み込む
// 必須の環境変数が設定されていない場合はエラーを返す
func LoadConfig() (*Config, error) {
	config := &Config{
		Port: getEnvOrDefault("PORT", "8080"),
		Env:  getEnvOrDefault("APP_ENV", "development"),

		Redis: RedisConfig{
			Host: mustGetEnv("REDIS_HOST"),
			Port: getEnvOrDefault("REDIS_PORT", "6379"),
		},

		DB: DBConfig{
			Host:     mustGetEnv("DB_HOST"),
			Port:     getEnvOrDefault("DB_PORT", "5432"),
			User:     mustGetEnv("DB_USER"),
			Password: mustGetEnv("DB_PASSWORD"),
			Name:     mustGetEnv("DB_NAME"),
		},
	}

	var missingEnvVars []string
	for _, err := range validateConfig(config) {
		if err != nil {
			missingEnvVars = append(missingEnvVars, err.Error())
		}
	}

	if len(missingEnvVars) > 0 {
		return nil, fmt.Errorf("required environment variables are not set: %s", strings.Join(missingEnvVars, ", "))
	}

	return config, nil
}

// validateConfig はConfigの値を検証
func validateConfig(config *Config) []error {
	var errors []error

	// 必須の環境変数を検証
	if config.DB.Host == "" {
		errors = append(errors, fmt.Errorf("DB_HOST"))
	}
	if config.DB.User == "" {
		errors = append(errors, fmt.Errorf("DB_USER"))
	}
	if config.DB.Password == "" {
		errors = append(errors, fmt.Errorf("DB_PASSWORD"))
	}
	if config.DB.Name == "" {
		errors = append(errors, fmt.Errorf("DB_NAME"))
	}
	if config.Redis.Host == "" {
		errors = append(errors, fmt.Errorf("REDIS_HOST"))
	}

	return errors
}

// 環境変数を取得し、設定されていない場合はデフォルト値を返す
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// 環境変数を取得する
// 環境変数が設定されていない場合は空文字を返す
func mustGetEnv(key string) string {
	return os.Getenv(key)
}

// LoadTestDBConfig はテスト用のDB設定を環境変数から読み込む
func LoadTestDBConfig() DBConfig {
	return DBConfig{
		Host:     getEnvOrDefault("TEST_DB_HOST", "localhost"),
		Port:     getEnvOrDefault("TEST_DB_PORT", "5433"),
		User:     getEnvOrDefault("TEST_DB_USER", "postgres"),
		Password: getEnvOrDefault("TEST_DB_PASSWORD", "password"),
		Name:     getEnvOrDefault("TEST_DB_NAME", "kaimon_test"),
	}
}
