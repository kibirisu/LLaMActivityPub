package main

import (
	"context"
	"encoding/json"
	"fmt"
	db "llamap/server/db/sqlc"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

var (
	queries *db.Queries
	ctx     context.Context
)

func main() {
	ctx = context.Background()
	conn, err := pgx.Connect(ctx, "postgres://llamap:changeme@localhost:5432/dev")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries = db.New(conn)

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
