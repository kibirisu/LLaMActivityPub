package router

import (
	"encoding/json/v2" // experimental features
	"io/fs"
	"log"
	"net/http"

	"borg/pkg/datastore"
	"borg/pkg/db"
	"borg/web"
)

type Router struct {
	http.Handler
	ds     datastore.DataStore
	assets fs.FS
}

func NewRouter(appEnv string, q *db.Queries) *Router {
	assets, err := web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	ds := datastore.NewDataStore(q)
	r := &Router{ds: ds, assets: assets}

	h := http.NewServeMux()

	if appEnv == "prod" {
		h.HandleFunc("/", r.handleRoot)
		h.HandleFunc("/static/", func(w http.ResponseWriter, req *http.Request) {
			http.StripPrefix("/", http.HandlerFunc(r.handleAssets)).ServeHTTP(w, req)
		})
	}
	h.HandleFunc("GET /api/", r.handleGetUsers)
	h.HandleFunc("POST /api/", r.handleCreateUser)

	r.Handler = h

	return r
}

func (h *Router) handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, h.assets, "index.html")
}

func (h *Router) handleAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServerFS(h.assets).ServeHTTP(w, r)
}

func (h *Router) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user db.AddUserQueryParams
	if err := json.UnmarshalRead(r.Body, &user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.ds.AddUser(r.Context(), user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Router) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.ds.GetUsers(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = json.MarshalWrite(w, &users); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
