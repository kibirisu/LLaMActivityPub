package sqlite

import (
	"context"

	"borg/pkg/db/models"
)

func (u User) toModel() models.User {
	return models.User(u)
}

func toModel(c models.CreateUserParams) CreateUserParams {
	return CreateUserParams(c)
}

func (q *Queries) GetUsersQuery(ctx context.Context) ([]models.User, error) {
	u, err := q.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	var users []models.User
	for _, user := range u {
		users = append(users, user.toModel())
	}
	return users, nil
}

func (q *Queries) CreateUserQuery(ctx context.Context, arg models.CreateUserParams) error {
	return q.CreateUser(ctx, toModel(arg))
}
