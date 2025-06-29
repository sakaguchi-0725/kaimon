package registry

import (
	"backend/domain/repository"
	"backend/infra/db"
	"backend/infra/firebase"
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

func NewRepository(db *db.Conn) (*Repository, error) {
	firebase, err := firebase.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	return &Repository{
		Authenticator: persistence.NewAuthenticator(firebase),
		Transaction:   persistence.NewTransaction(db),
		Account:       persistence.NewAccount(db),
		User:          persistence.NewUser(db),
		Group:         persistence.NewGroup(db),
		GroupMember:   persistence.NewGroupMember(db),
	}, nil
}
