package persistence

import (
	"backend/domain/repository"
	"backend/infra/firebase"
	"context"
	"fmt"
)

type authenticator struct {
	firebaseClient firebase.Client
}

func (a *authenticator) VerifyToken(token string) (uid string, email string, err error) {
	ctx := context.Background()

	authToken, err := a.firebaseClient.VerifyIDToken(ctx, token)
	if err != nil {
		return "", "", fmt.Errorf("トークンの検証に失敗しました: %w", err)
	}

	uid = authToken.UID

	// emailが存在する場合のみ設定
	if emailValue, ok := authToken.Claims["email"].(string); ok {
		email = emailValue
	}

	return uid, email, nil
}

func NewAuthenticator(firebaseClient firebase.Client) repository.Authenticator {
	return &authenticator{firebaseClient: firebaseClient}
}
