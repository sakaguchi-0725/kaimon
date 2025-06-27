package persistence

import (
	"backend/domain/repository"
	"backend/infra/firebase"
	"context"

	"firebase.google.com/go/v4/errorutils"
)

type authenticator struct {
	firebaseClient firebase.Client
}

func (a *authenticator) VerifyToken(token string) (uid string, email string, err error) {
	ctx := context.Background()

	authToken, err := a.firebaseClient.VerifyIDToken(ctx, token)
	if err != nil {
		switch {
		case errorutils.IsInvalidArgument(err):
			return "", "", ErrInvalidToken
		case errorutils.IsUnauthenticated(err):
			return "", "", ErrTokenExpired
		case errorutils.IsPermissionDenied(err):
			return "", "", ErrAuthenticationFail
		default:
			return "", "", ErrAuthenticationFail
		}
	}

	uid = authToken.UID

	// emailが存在する場合のみ設定
	if emailValue, ok := authToken.Claims["email"].(string); ok {
		email = emailValue
	}

	return uid, email, nil
}

func (a *authenticator) SignInWithEmailAndPassword(email, password string) (uid string, err error) {
	// TODO: Implement Firebase Auth Email/Password sign-in
	// Note: Firebase Admin SDK doesn't provide direct email/password sign-in
	// This would typically be done on the client side and then the resulting ID token
	// would be verified server-side using VerifyToken
	return "", ErrInvalidCredentials
}

func NewAuthenticator(firebaseClient firebase.Client) repository.Authenticator {
	return &authenticator{firebaseClient: firebaseClient}
}
