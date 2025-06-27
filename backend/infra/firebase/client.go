//go:generate mockgen -source=client.go -destination=../../test/mock/firebase_client_mock.go -package=mock
package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type (
	Client interface {
		VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
	}

	client struct {
		*auth.Client
	}
)

func (c *client) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := c.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func NewClient(ctx context.Context) (Client, error) {
	// GOOGLE_APPLICATION_CREDENTIALS を使用して認証情報を設定
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase app: %w", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase auth client: %w", err)
	}

	return &client{authClient}, nil
}
