package core

import "context"

type ContextKey string

const (
	UserIDKey ContextKey = "userID"
	EmailKey  ContextKey = "email"
	TxKey     ContextKey = "tx"
)

func GetUserID(ctx context.Context) string {
	return ctx.Value(UserIDKey).(string)
}

func GetEmail(ctx context.Context) string {
	return ctx.Value(EmailKey).(string)
}
