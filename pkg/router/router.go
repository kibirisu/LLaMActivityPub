package router

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"net/http"

	"borg/pkg/db"
	"borg/web"
)

type Router struct {
	http.Handler
	db     *db.Queries
	assets fs.FS
}

func New(appEnv string, q *db.Queries) *Router {
	assets, err := web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	r := &Router{db: q, assets: assets}

	h := http.NewServeMux()

	if appEnv == "prod" {
		h.HandleFunc("/", r.handleRoot)
		h.HandleFunc("/static/", func(w http.ResponseWriter, req *http.Request) {
			http.StripPrefix("/", http.HandlerFunc(r.handleAssets)).ServeHTTP(w, req)
		})
	}
	h.HandleFunc("/api/", func(w http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(body))
		var user db.User
		json.Unmarshal(body, &user)
		log.Println(user)
		newUser := db.AddUserParams{
			Username:       user.Username,
			PasswordHash:   user.PasswordHash,
			Bio:            sql.NullString{},
			FollowersCount: sql.NullInt32{},
			FollowingCount: sql.NullInt32{},
			IsAdmin:        sql.NullBool{},
		}
		if err = r.db.AddUser(req.Context(), newUser); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	r.Handler = h

	return r
}

func (h *Router) handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, h.assets, "index.html")
}

func (h *Router) handleAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServerFS(h.assets).ServeHTTP(w, r)
}
