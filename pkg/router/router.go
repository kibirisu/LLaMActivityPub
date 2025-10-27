package router

import (
	"context"
	"io/fs"
	"log"
	"net/http"

	"borg/pkg/db"
	"borg/web"
)

type contextKey string

var (
	assets fs.FS
	dbKey  contextKey = "db"
)

func Serve(appEnv, port string, querier db.Querier) {
	r := http.NewServeMux()

	var err error
	assets, err = web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	addr := ":" + port

	if appEnv == "prod" {
		r.HandleFunc("/", provideQuerier(handleRoot, querier))
		r.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
			http.StripPrefix("/", http.HandlerFunc(handleAssets)).ServeHTTP(w, r)
		})
	}

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries, ok := ctx.Value(dbKey).(db.Querier)
	if !ok {
		log.Fatal("Could not get DB pool")
	}
	if _, err := queries.GetUsersQuery(ctx); err != nil {
		log.Println(err)
	}
	http.ServeFileFS(w, r, assets, "index.html")
}

func handleAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServerFS(assets).ServeHTTP(w, r)
}

func provideQuerier(h http.HandlerFunc, db db.Querier) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, dbKey, db)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
