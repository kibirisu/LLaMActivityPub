package db

import (
	"context"
	"database/sql"
	"log"

	"borg/pkg/db/postgres"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func GetDB(ctx context.Context, url string) (*postgres.Queries, error) {
	pool, err := sql.Open("pgx", url)
	if err != nil {
		return nil, err
	}
	log.Println("✅ Connected to DB")

	// Run database migrations
	migrations, err := getMigrations()
	if err != nil {
		return nil, err
	}
	goose.SetBaseFS(migrations)

	driver := "postgres"
	if err := goose.SetDialect(driver); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(pool, driver); err != nil {
		log.Fatal(err)
	}

	return postgres.New(pool), nil
}
