//go:generate mockgen -source=firebase.go -destination=../../test/mock/external/firebase_mock.go -package=mock
package external

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type (
	FirebaseClient interface {
		VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
	}

	firebaseClient struct {
		*auth.Client
	}
)

func (c *firebaseClient) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := c.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func NewFirebaseClient(ctx context.Context) (FirebaseClient, error) {
	// GOOGLE_APPLICATION_CREDENTIALS を使用して認証情報を設定
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase app: %w", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase auth client: %w", err)
	}

	return &firebaseClient{authClient}, nil
}
