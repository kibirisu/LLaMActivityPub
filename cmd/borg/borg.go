package main

import (
	"context"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"borg/pkg/config"
	"borg/pkg/db"
	"borg/web"
)

type contextKey string

var (
	assets fs.FS
	dbKey  contextKey = "db"
)

func main() {
	ctx := context.Background()
	conf := config.GetConfig()

	db, err := db.GetDB(ctx, conf.DatabaseDriver, conf.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	assets, err = web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// Define routes
	if conf.AppEnv == "prod" {
		mux.HandleFunc("/", withMiddleware(handleApp, db))
	}

	// Start the HTTP server
	addr := ":8080"
	log.Printf("🚀 Server running at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

func withMiddleware(h http.HandlerFunc, db db.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, dbKey, db)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

// Can be done more effectively
func handleApp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries, ok := ctx.Value(dbKey).(db.Queries)
	if !ok {
		log.Fatal("Could not get DB pool")
	}
	if _, err := queries.GetUsersQuery(ctx); err != nil {
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
