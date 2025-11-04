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

func (h *Router) handleUserRoutes(r chi.Router) {
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
		var user db.AddUserParams
		if err := json.UnmarshalRead(r.Body, &user, ds.Opts()); err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusBadRequest)
			return
		}
		if err := ds.UserRepository().Create(r.Context(), user); err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError)
			return
		}
		if err := writeSuccess(w, &user, http.StatusCreated, ds.Opts()); err != nil {
			log.Println(err)
		}
	}
}

func getUsers(ds data.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := ds.UserRepository().GetAll(r.Context())
		if err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError)
			return
		}
		if err = writeSuccess(w, &users, http.StatusOK, ds.Opts()); err != nil {
			log.Println(err)
		}
	}
}

func getUser(ds data.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(keyID).(int32)
		user, err := ds.UserRepository().GetByID(r.Context(), id)
		if err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError)
			return
		}
		if err = writeSuccess(w, &user, http.StatusOK, ds.Opts()); err != nil {
			log.Println(err)
		}
	}
}

func deleteUser(ds data.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(keyID).(int32)
		if err := ds.UserRepository().Delete(r.Context(), id); err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError)
			return
		}
		if err := writeSuccess(w, id, http.StatusOK); err != nil {
			log.Println(err)
		}
	}
}

func updateUser(ds data.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user db.UpdateUserParams
		id := r.Context().Value(keyID).(int32)
		opts := ds.Opts()
		if err := json.UnmarshalRead(r.Body, &user, opts); err != nil || user.ID != int32(id) {
			_ = writeError(w, err, http.StatusBadRequest)
			return
		}
		if err := ds.UserRepository().Update(r.Context(), user); err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError)
			return
		}
		if err := writeSuccess(w, &user, http.StatusOK, opts); err != nil {
			log.Println(err)
		}
	}
}
