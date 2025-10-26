package db

import (
	"embed"
	"io/fs"
)

//go:embed migrations/*/*.sql
var migrations embed.FS

func getMigrations() (res fs.FS, err error) {
	res, err = fs.Sub(migrations, "migrations")
	return
}
