package data

import (
	"context"

	"borg/pkg/db"
)

type UserRepository interface {
	AddUser(context.Context, db.AddUserQueryParams) error
	GetUsers(context.Context) ([]db.User, error)
	GetUser(context.Context, int32) (db.User, error)
	DeleteUser(context.Context, int32) error
	UpdateUser(context.Context, db.UpdateUserQueryParams) error
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

func (r *userRepository) GetUser(ctx context.Context, id int32) (db.User, error) {
	return r.GetUserQuery(ctx, id)
}

func (r *userRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.DeleteUserQuery(ctx, id)
}

func (r *userRepository) UpdateUser(ctx context.Context, user db.UpdateUserQueryParams) error {
	return r.UpdateUserQuery(ctx, user)
}
