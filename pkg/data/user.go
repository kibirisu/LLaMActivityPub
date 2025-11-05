package data

import (
	"context"

	"borg/pkg/db"
)

type UserRepository interface {
	Repository[db.User, db.AddUserParams, db.UpdateUserParams]
	GetFollowed(context.Context, int32) ([]db.User, error)
	GetFollowers(context.Context, int32) ([]db.User, error)
}

type userRepository struct {
	*db.Queries
}

func newUserRepository(q *db.Queries) UserRepository {
	return &userRepository{q}
}

func (r *userRepository) Create(ctx context.Context, user db.AddUserParams) error {
	return r.AddUser(ctx, user)
}

func (r *userRepository) GetByID(ctx context.Context, id int32) (db.User, error) {
	return r.GetUser(ctx, id)
}

func (r *userRepository) Update(ctx context.Context, user db.UpdateUserParams) error {
	return r.UpdateUser(ctx, user)
}

func (r *userRepository) Delete(ctx context.Context, id int32) error {
	return r.DeleteUser(ctx, id)
}

func (r *userRepository) GetFollowed(ctx context.Context, id int32) ([]db.User, error) {
	return r.GetFollowedUsers(ctx, id)
}

func (r *userRepository) GetFollowers(ctx context.Context, id int32) ([]db.User, error) {
	return r.GetFollowingUsers(ctx, id)
}
