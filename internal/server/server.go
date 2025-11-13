package server

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"

	"borg/internal/api"
	"borg/internal/domain"
	"borg/internal/models"
	"borg/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	ds     domain.DataStore
	assets fs.FS
}

func NewServer(listenPort string, ds domain.DataStore) *http.Server {
	assets, err := web.GetAssets()
	if err != nil {
		panic(err)
	}
	server := &Server{
		ds:     ds,
		assets: assets,
	}
	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/", func(r chi.Router) {
		r.Get("/*", server.handleRoot)
		r.Get("/static/*", server.handleAssets)
	})
	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:" + listenPort,
	}
	return s
}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, s.assets, "index.html")
}

func (s *Server) handleAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServerFS(s.assets).ServeHTTP(w, r)
}

// DeleteApiUsersId implements api.ServerInterface.
func (s *Server) DeleteApiUsersId(w http.ResponseWriter, r *http.Request, id int) {
	panic("unimplemented")
}

// GetApiUsersId implements api.ServerInterface.
func (s *Server) GetApiUsersId(w http.ResponseWriter, r *http.Request, id int) {
	user, err := s.ds.UserRepository().GetByID(r.Context(), int32(id))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := models.UserFromDBType(&user)
	_ = json.NewEncoder(w).Encode(res)
}

// PostApiUsers implements api.ServerInterface.
func (s *Server) PostApiUsers(w http.ResponseWriter, r *http.Request) {
	var user api.PostApiUsersJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := s.ds.UserRepository().Create(r.Context(), *models.AddUserToDBType(&user)); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// PutApiUsersId implements api.ServerInterface.
func (s *Server) PutApiUsersId(w http.ResponseWriter, r *http.Request, id int) {
	panic("unimplemented")
}
