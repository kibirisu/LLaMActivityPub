package models

import "database/sql"

type User struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type CreateUserParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
