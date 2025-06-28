package core

import "context"

type ContextKey string

const (
	UserIDKey ContextKey = "userID"
	EmailKey  ContextKey = "email"
	TxKey     ContextKey = "tx"
)

func GetUserID(ctx context.Context) string {
	if value := ctx.Value(UserIDKey); value != nil {
		return value.(string)
	}
	return ""
}

func GetEmail(ctx context.Context) string {
	if value := ctx.Value(EmailKey); value != nil {
		return value.(string)
	}
	return ""
}
