//go:generate mockgen -source=authenticator.go -destination=../../test/mock/repository/authenticator_mock.go -package=mock
package repository

import "context"

type Authenticator interface {
	VerifyToken(ctx context.Context, token string) (uid, email, name string, err error)
}
