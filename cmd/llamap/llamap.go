package main

import (
	"context"
	"database/sql"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"llamap/pkg/config"
	"llamap/pkg/db"
	"llamap/web"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

var (
	assets fs.FS
	ctx    context.Context //nolint
)

func main() {
	conf := config.GetConfig()

	// Connect to PostgreSQL
	pool, err := sql.Open("pgx", conf.DatabaseUrl)
	if err != nil {
		log.Fatalf("❌ Failed to open DB: %v", err)
	}
	defer pool.Close() //nolint

	// Verify connection
	if err := pool.Ping(); err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}
	log.Println("✅ Connected to PostgreSQL")

	ctx = context.Background()

	// Run database migrations
	migrations, err := db.GetMigrations()
	if err != nil {
		log.Fatal(err)
	}
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(pool, "postgres"); err != nil {
		log.Fatal(err)
	}

	assets, err = web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// Define routes
	if conf.AppEnv == "prod" {
		mux.HandleFunc("/", handleApp)
	}

	// Start the HTTP server
	addr := ":8080"
	log.Printf("🚀 Server running at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

// Can be done more effectively
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
