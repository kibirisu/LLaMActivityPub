package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver for database/sql
)

var db *sql.DB

func main() {
	// Read the database connection string
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Fallback for local dev (adjust credentials as needed)
		dsn = "postgres://dev:password@localhost:5432/devdb?sslmode=disable"
	}

	// Connect to PostgreSQL
	var err error
	db, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to open DB: %v", err)
	}
	defer db.Close()

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}
	log.Println("✅ Connected to PostgreSQL")

	// Define routes
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/users", usersHandler)

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
	rows, err := db.Query("SELECT id, name FROM users LIMIT 10")
	if err != nil {
		http.Error(w, fmt.Sprintf("DB error: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			http.Error(w, fmt.Sprintf("Row error: %v", err), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	 writeJSON(w, http.StatusOK, map[string]interface{}{"users": users})
}

// Helper to write JSON responses
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to write JSON: %v", err)
	}
}

