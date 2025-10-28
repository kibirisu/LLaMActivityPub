package router

import (
	"context"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"net/http"

	"borg/pkg/db"
	"borg/pkg/db/models"
	"borg/web"
)

type contextKey string

var (
	assets fs.FS
	dbKey  contextKey = "db"
)

func Serve(appEnv, port string, q db.Querier) {
	r := http.NewServeMux()

	var err error
	assets, err = web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	addr := ":" + port

	if appEnv == "prod" {
		r.HandleFunc("/", handleRoot)
		r.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
			http.StripPrefix("/", http.HandlerFunc(handleAssets)).ServeHTTP(w, r)
		})
	}
	r.HandleFunc("POST /api/", provideQuerier(handleFoo, q))

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, assets, "index.html")
}

func handleAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServerFS(assets).ServeHTTP(w, r)
}

func handleFoo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	q, ok := ctx.Value(dbKey).(db.Querier)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user models.CreateUserParams
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(user.Name)
	q.CreateUserQuery(ctx, user)
	if users, err := q.GetUsersQuery(ctx); err != nil {
		log.Println(err)
	} else {
		log.Println(users[0].Email)
	}
	w.WriteHeader(http.StatusOK)
}

func provideQuerier(h http.HandlerFunc, db db.Querier) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, dbKey, db)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
