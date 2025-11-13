package server

import (
	"net/http"

	"borg/internal/api"
	"borg/internal/domain"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	ds domain.DataStore
}

func NewServer(listenPort string, ds domain.DataStore) *http.Server {
	server := &Server{
		ds: ds,
	}
	h := api.Handler(server)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:" + listenPort,
	}
	return s
}

// DeleteApiUsersId implements api.ServerInterface.
func (s *Server) DeleteApiUsersId(w http.ResponseWriter, r *http.Request, id int) {
	panic("unimplemented")
}

// GetApiUsersId implements api.ServerInterface.
func (s *Server) GetApiUsersId(w http.ResponseWriter, r *http.Request, id int) {
	panic("unimplemented")
}

// PostApiUsers implements api.ServerInterface.
func (s *Server) PostApiUsers(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PutApiUsersId implements api.ServerInterface.
func (s *Server) PutApiUsersId(w http.ResponseWriter, r *http.Request, id int) {
	panic("unimplemented")
}
