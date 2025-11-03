package router

import (
	"context"
	"encoding/json/v2"
	"log"
	"net/http"
	"strconv"

	"borg/pkg/data"
	"borg/pkg/db"

	"github.com/go-chi/chi/v5"
)

func (h *Router) addUserRoute(r chi.Router) {
	r.Use(optionCtx)
	r.Get("/", getUsers(h.ds))
	r.Post("/", createUser(h.ds))
	r.Route("/{id}", func(r chi.Router) {
		r.Use(userCtx)
		r.Get("/", getUser(h.ds))
		r.Delete("/", deleteUser(h.ds))
		r.Put("/", updateUser(h.ds))
	})
}

func userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idVal := chi.URLParam(r, "id")
		if idVal == "" {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idVal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		ctx := context.WithValue(r.Context(), "id", int32(id))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func optionCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "opts", data.GetOptions())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createUser(repo data.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user db.AddUserQueryParams
		opts := r.Context().Value("opts").(json.Options)
		if err := json.UnmarshalRead(r.Body, &user, opts); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := repo.AddUser(r.Context(), user); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func getUsers(repo data.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repo.GetUsers(r.Context())
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opts := r.Context().Value("opts").(json.Options)
		if err = json.MarshalWrite(w, &users, opts); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func getUser(repo data.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("id").(int32)
		opts := r.Context().Value("opts").(json.Options)
		user, err := repo.GetUser(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = json.MarshalWrite(w, &user, opts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func deleteUser(repo data.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("id").(int32)
		if err := repo.DeleteUser(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func updateUser(repo data.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user db.UpdateUserQueryParams
		id := r.Context().Value("id").(int32)
		opts := r.Context().Value("opts").(json.Options)
		if err := json.UnmarshalRead(r.Body, &user, opts); err != nil || user.ID != int32(id) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := repo.UpdateUser(r.Context(), user); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
