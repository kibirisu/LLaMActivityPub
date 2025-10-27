package db

import (
	"context"
	"database/sql"
	"log"

	"borg/pkg/db/models"
	"borg/pkg/db/postgres"
	"borg/pkg/db/sqlite"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

type Querier interface {
	GetUsersQuery(context.Context) ([]models.User, error)
	CreateUserQuery(context.Context, models.CreateUserParams) error
}

func GetDB(ctx context.Context, driver, url string) (Querier, error) {
	var driverName string
	switch driver {
	case "sqlite":
		driverName = "sqlite"
		log.Println("Using sqlite driver")
	case "postgres":
		driverName = "pgx"
		log.Println("Using postgres driver")
	default:
		log.Fatal("Uknown db driver name")
	}

	pool, err := sql.Open(driverName, url)
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

	if err := goose.SetDialect(driver); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(pool, driver); err != nil {
		log.Fatal(err)
	}

	var res Querier
	if driver == "sqlite" {
		res = sqlite.New(pool)
	} else {
		res = postgres.New(pool)
	}
	return res, nil
}
