package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func GetDB(driver, url string) (*sql.DB, error) {
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
	return pool, nil
}
