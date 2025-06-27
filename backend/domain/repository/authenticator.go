//go:generate mockgen -source=authenticator.go -destination=../../test/mock/repository/authenticator_mock.go -package=mock
package repository

type Authenticator interface {
	VerifyToken(token string) (uid, email string, err error)
}
