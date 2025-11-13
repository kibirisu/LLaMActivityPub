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

func UpdateUserToDBType(u *api.UpdateUser) *db.UpdateUserParams {
	var bio sql.NullString
	var isAdmin sql.NullBool
	if u.Bio != nil {
		bio = sql.NullString{
			String: *u.Bio,
			Valid:  true,
		}
	}
	if u.IsAdmin != nil {
		isAdmin = sql.NullBool{
			Bool:  *u.IsAdmin,
			Valid: true,
		}
	}
	return &db.UpdateUserParams{
		ID:             0,
		PasswordHash:   "",
		Bio:            bio,
		FollowersCount: sql.NullInt32{},
		FollowingCount: sql.NullInt32{},
		IsAdmin:        isAdmin,
	}
}
