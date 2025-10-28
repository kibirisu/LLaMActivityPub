package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func GetDB(ctx context.Context, url string) (*Queries, error) {
	pool, err := sql.Open("pgx", url)
	if err != nil {
		return nil, err
	}
	log.Println("✅ Connected to DB")

	// Run database migrations
	goose.SetBaseFS(getMigrations())

	driver := "postgres"
	if err := goose.SetDialect(driver); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(pool, "migrations"); err != nil {
		log.Fatal(err)
	}

	return New(pool), nil
}
