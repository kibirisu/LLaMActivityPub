package main

import (
	"context"
	"database/sql"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"borg/pkg/config"
	"borg/pkg/db"
	"borg/pkg/db/postgres"
	"borg/web"
)

var assets fs.FS

func main() {
	conf := config.GetConfig()

	db, err := db.GetDB(conf.DatabaseDriver, conf.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	assets, err = web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// Define routes
	ctx := context.Background()
	if conf.AppEnv == "prod" {
		mux.HandleFunc("/", withMiddleware(ctx, handleApp, db))
	}

	// Start the HTTP server
	addr := ":8080"
	log.Printf("🚀 Server running at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

func withMiddleware(ctx context.Context, h http.HandlerFunc, dbtx *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(ctx, "dbtx", dbtx)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

// Can be done more effectively
func handleApp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	bar := ctx.Value("dbtx").(*sql.DB)
	queries := postgres.New(bar)
	if _, err := queries.GetUsers(ctx); err != nil {
		log.Println(err)
	}
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
