package router

import (
	"encoding/json/v2" // experimental features
	"io/fs"
	"log"
	"net/http"
	"strconv"

	"borg/pkg/data"
	"borg/pkg/db"
	"borg/web"
)

type Router struct {
	http.Handler
	ds     data.DataStore
	assets fs.FS
}

func NewRouter(appEnv string, q *db.Queries) *Router {
	assets, err := web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	ds := data.NewDataStore(q)
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
	h.HandleFunc("GET /api/{id}", r.handleGetUser)
	h.HandleFunc("DELETE /api/{id}", r.handleDeleteUser)
	h.HandleFunc("PUT /api/{id}", r.handleUpdateUser)
	h.HandleFunc("/foo/", r.handleFoo)

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.MarshalWrite(w, &users); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Router) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idVal := r.PathValue("id")
	if idVal == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idVal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.ds.GetUser(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.MarshalWrite(w, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Router) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	idVal := r.PathValue("id")
	if idVal == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idVal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = h.ds.DeleteUser(r.Context(), int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Router) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	idVal := r.PathValue("id")
	if idVal == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idVal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user db.UpdateUserQueryParams
	if err = json.UnmarshalRead(r.Body, &user); err != nil || user.ID != int32(id) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = h.ds.UpdateUser(r.Context(), user); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Router) handleFoo(w http.ResponseWriter, r *http.Request) {
	var payload db.AddUserQueryParams
	if err := json.UnmarshalRead(r.Body, &payload, data.WithUnmarshalers()); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.MarshalWrite(w, payload, data.WithMarshalers())
}
