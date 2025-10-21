package main

import (
	"context"
	"database/sql"
	"embed"
	"encoding/json"
	"fmt"
	db "llamap/server/db/sqlc"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

var (
	//go:embed db/migrations/*.sql
	embedMigrations embed.FS
	queries         *db.Queries
	ctx             context.Context
)

func main() {
	// Read the database connection string
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Fallback for local dev (adjust credentials as needed)
		dsn = "postgres://dev:password@localhost:5432/devdb?sslmode=disable"
	}

	// Connect to PostgreSQL
	var err error
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
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(pool, "db/migrations"); err != nil {
		log.Fatal(err)
	}

	// Define routes
	http.Handle("/", http.FileServer(http.Dir("../web/dist")))
	http.HandleFunc("/api/ping", pingHandler)
	http.HandleFunc("/api/users", usersHandler)

	// Start the HTTP server
	addr := ":8080"
	log.Printf("🚀 Server running at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

// pingHandler — simple health check endpoint
func pingHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"message": "pong"})
}

// usersHandler — fetch some users from the DB
func usersHandler(w http.ResponseWriter, r *http.Request) {
	res, err := queries.GetUsers(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("DB error: %v", err), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"users": res})
}

// Helper to write JSON responses
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to write JSON: %v", err)
	}
}
