package models

import (
	"database/sql"

	"borg/internal/api"
	"borg/internal/db"
)

func UserFromDBType(u *db.User) *api.User {
	return &api.User{
		Bio:            u.Bio.String,
		CreatedAt:      u.CreatedAt.Time,
		FollowersCount: int(u.FollowersCount.Int32),
		FollowingCount: int(u.FollowingCount.Int32),
		Id:             int(u.ID),
		IsAdmin:        u.IsAdmin.Bool,
		Origin:         u.Origin.String,
		UpdatedAt:      u.UpdatedAt.Time,
		Username:       u.Username,
	}
}

func AddUserToDBType(u *api.NewUser) *db.AddUserParams {
	return &db.AddUserParams{
		Username:       u.Username,
		PasswordHash:   "",
		Bio:            sql.NullString{},
		FollowersCount: sql.NullInt32{},
		FollowingCount: sql.NullInt32{},
		IsAdmin:        sql.NullBool{},
	}
}
