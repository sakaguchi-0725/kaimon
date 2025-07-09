package registry

import (
	"backend/core"
	"backend/domain/repository"
	"backend/infra/db"
	"backend/infra/external"
	"backend/infra/persistence"
	"context"
)

type Repository struct {
	Authenticator repository.Authenticator
	Transaction   repository.Transaction

	Account     repository.Account
	User        repository.User
	Group       repository.Group
	GroupMember repository.GroupMember
}

func NewRepository(db *db.Conn, cfg core.RedisConfig) (*Repository, error) {
	firebase, err := external.NewFirebaseClient(context.Background())
	if err != nil {
		return nil, err
	}

	redis := external.NewRedisClient(cfg)

	return &Repository{
		Authenticator: persistence.NewAuthenticator(firebase),
		Transaction:   persistence.NewTransaction(db),
		Account:       persistence.NewAccount(db),
		User:          persistence.NewUser(db),
		Group:         persistence.NewGroup(db, redis),
		GroupMember:   persistence.NewGroupMember(db),
	}, nil
}
