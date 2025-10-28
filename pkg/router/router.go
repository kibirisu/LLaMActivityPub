package router

import (
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"net/http"

	"borg/pkg/db"
	"borg/pkg/db/models"
	"borg/web"
)

type Router struct {
	http.Handler
	db     db.Querier
	assets fs.FS
}

func New(appEnv string, q db.Querier) *Router {
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
	h.HandleFunc("POST /api/", r.handleFoo)

	r.Handler = h

	return r
}

func (h *Router) handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, h.assets, "index.html")
}

func (h *Router) handleAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServerFS(h.assets).ServeHTTP(w, r)
}

func (h *Router) handleFoo(w http.ResponseWriter, r *http.Request) {
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
	err = h.db.CreateUserQuery(r.Context(), user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if users, err := h.db.GetUsersQuery(r.Context()); err != nil {
		log.Println(err)
	} else {
		log.Println(users[0].Email)
	}
	w.WriteHeader(http.StatusOK)
}
