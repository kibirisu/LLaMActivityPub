package db

import "embed"

//go:embed migrations/*.sql
var migrations embed.FS

func getMigrations() embed.FS {
	return migrations
}
