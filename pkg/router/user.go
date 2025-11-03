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

type contextKey string

const keyID = contextKey("keyId")

func (h *Router) addUserRoute(r chi.Router) {
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
		ctx := context.WithValue(r.Context(), keyID, int32(id))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createUser(ds data.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user db.AddUserQueryParams
		if err := json.UnmarshalRead(r.Body, &user, ds.GetOpts()); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := ds.AddUser(r.Context(), user); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func getUsers(ds data.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := ds.GetUsers(r.Context())
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = json.MarshalWrite(w, &users, ds.GetOpts()); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func getUser(ds data.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(keyID).(int32)
		user, err := ds.GetUser(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = json.MarshalWrite(w, &user, ds.GetOpts()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func deleteUser(ds data.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(keyID).(int32)
		if err := ds.DeleteUser(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func updateUser(ds data.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user db.UpdateUserQueryParams
		id := r.Context().Value(keyID).(int32)
		if err := json.UnmarshalRead(r.Body, &user, ds.GetOpts()); err != nil || user.ID != int32(id) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := ds.UpdateUser(r.Context(), user); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
