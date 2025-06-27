package core

import (
	"errors"
	"strings"
)

// 指定されたプレフィックスを入力文字列から削除します。
// プレフィックスが見つからない場合はエラーを返します。
func RemovePrefix(input, prefix string) (string, error) {
	if input == "" {
		return "", errors.New("input string is required")
	}

	if prefix == "" {
		return "", errors.New("prefix is required")
	}

	if !strings.HasPrefix(input, prefix) {
		return "", errors.New("input string must start with '" + prefix + "'")
	}

	result := strings.TrimPrefix(input, prefix)
	if result == "" {
		return "", errors.New("token is required after removing prefix")
	}

	return result, nil
}
