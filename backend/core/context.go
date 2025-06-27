package core

import "context"

type ContextKey string

const (
	UserIDKey ContextKey = "userID"
	TxKey     ContextKey = "tx"
)

func GetUserID(ctx context.Context) string {
	return ctx.Value(UserIDKey).(string)
}
