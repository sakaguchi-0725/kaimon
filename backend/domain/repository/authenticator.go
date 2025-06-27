package repository

type Authenticator interface {
	VerifyToken(token string) (uid, email string, err error)
}
