package datastore

import (
	"context"

	"borg/pkg/db"
)

type UserRepository interface {
	AddUser(context.Context, db.AddUserQueryParams) error
	GetUsers(context.Context) ([]db.User, error)
}

type userRepository struct {
	*db.Queries
}

func newUserRepository(q *db.Queries) UserRepository {
	return &userRepository{q}
}

func (r *userRepository) AddUser(ctx context.Context, user db.AddUserQueryParams) error {
	return r.AddUserQuery(ctx, user)
}

func (r *userRepository) GetUsers(ctx context.Context) ([]db.User, error) {
	return r.GetUsersQuery(ctx)
}
