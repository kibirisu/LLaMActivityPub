package db

import (
	"embed"
	"io/fs"
)

//go:embed migrations/*/*.sql
var migrations embed.FS

func GetMigrations() (res fs.FS, err error) {
	res, err = fs.Sub(migrations, "migrations")
	return
}
