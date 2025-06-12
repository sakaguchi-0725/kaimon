package core

import "context"

type ContextKey string

const (
	UserIDKey ContextKey = "userID"
)

func GetUserID(ctx context.Context) string {
	return ctx.Value(UserIDKey).(string)
}
