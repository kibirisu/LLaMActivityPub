package main

import (
	"context"
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	db "llamap/db/sqlc"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

var (
	//go:embed web/dist
	appDist embed.FS
	//go:embed db/migrations/*.sql
	migrations embed.FS
	assets     fs.FS
	queries    *db.Queries
	ctx        context.Context
)

func main() {
	// Read the database connection string
	dsn := os.Getenv("DATABASE_URL")
	env := os.Getenv("APP_ENV")
	if dsn == "" {
		// Fallback for local dev (adjust credentials as needed)
		dsn = "postgres://dev:password@localhost:5432/devdb?sslmode=disable"
	}
	if env == "" {
		env = "dev"
	}

	// Connect to PostgreSQL
	pool, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to open DB: %v", err)
	}
	defer pool.Close()

	// Verify connection
	if err := pool.Ping(); err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}
	log.Println("✅ Connected to PostgreSQL")

	ctx = context.Background()

	// Run database migrations
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(pool, "db/migrations"); err != nil {
		log.Fatal(err)
	}

	if err := getAssets(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// Define routes
	if env == "prod" {
		mux.HandleFunc("/", handleApp)
	}

	// Start the HTTP server
	addr := ":8080"
	log.Printf("🚀 Server running at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

func getAssets() error {
	var err error
	assets, err = fs.Sub(appDist, "web/dist")
	if err != nil {
		return err
	}
	return nil
}

func handleApp(w http.ResponseWriter, r *http.Request) {
	file := strings.TrimPrefix(r.URL.Path, "/")
	info, err := fs.Stat(assets, file)
	if err != nil {
		log.Println(err)
		file = "index.html"
	} else if !info.Mode().IsRegular() {
		file = "index.html"
	}
	http.ServeFileFS(w, r, assets, file)
}
