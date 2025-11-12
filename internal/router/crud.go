package router

import (
	"context"
	"encoding/json/v2"
	"log"
	"net/http"
	"strconv"

	"borg/internal/domain"

	"github.com/go-chi/chi/v5"
)

type contextKey string

const keyID contextKey = contextKey("keyID")

func idCtx(next http.Handler) http.Handler {
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

func create[T, C, U any, R domain.Repository[T, C, U]](repo R, opts json.Options) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item C
		if err := json.UnmarshalRead(r.Body, &item, opts); err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusBadRequest)
			return
		}
		if err := repo.Create(r.Context(), item); err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError)
			return
		}
		if err := writeSuccess(w, &item, http.StatusCreated, opts); err != nil {
			log.Println(err)
		}
	}
}

func getByID[T, C, U any, R domain.Repository[T, C, U]](
	repo R,
	opts json.Options,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(keyID).(int32)
		item, err := repo.GetByID(r.Context(), id)
		if err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError)
			return
		}
		if err = writeSuccess(w, &item, http.StatusOK, opts); err != nil {
			log.Println(err)
		}
	}
}

func getByUserID[T, C, U any, R domain.UserScopedRepository[T, C, U]](
	repo R,
	opts json.Options,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(keyID).(int32)
		items, err := repo.GetByUserID(r.Context(), id)
		if err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError)
			return
		}
		if err = writeSuccess(w, &items, http.StatusOK, opts); err != nil {
			log.Println(err)
		}
	}
}

func getByPostID[T, C, U any, R domain.PostScopedRepository[T, C, U]](
	repo R,
	opts json.Options,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(keyID).(int32)
		items, err := repo.GetByPostID(r.Context(), id)
		if err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError)
			return
		}
		if err = writeSuccess(w, &items, http.StatusOK, opts); err != nil {
			log.Println(err)
		}
	}
}

func delete[T, C, U any, R domain.Repository[T, C, U]](repo R) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(keyID).(int32)
		if err := repo.Delete(r.Context(), id); err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError)
			return
		}
		if err := writeSuccess(w, id, http.StatusOK); err != nil {
			log.Println(err)
		}
	}
}

func update[T, C, U any, R domain.Repository[T, C, U]](repo R, opts json.Options) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item U
		if err := json.UnmarshalRead(r.Body, &item, opts); err != nil {
			_ = writeError(w, err, http.StatusBadRequest)
			return
		}
		if err := repo.Update(r.Context(), item); err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError)
			return
		}
		if err := writeSuccess(w, &item, http.StatusOK, opts); err != nil {
			log.Println(err)
		}
	}
}

func getFollowers(repo domain.UserRepository, opts json.Options) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(keyID).(int32)
		users, err := repo.GetFollowers(r.Context(), id)
		if err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError, opts)
			return
		}
		if err = writeSuccess(w, users, int(http.StatusOK), opts); err != nil {
			log.Println(err)
		}
	}
}

func getFollowing(repo domain.UserRepository, opts json.Options) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(keyID).(int32)
		users, err := repo.GetFollowed(r.Context(), id)
		if err != nil {
			log.Println(err)
			_ = writeError(w, err, http.StatusInternalServerError, opts)
			return
		}
		if err = writeSuccess(w, users, int(http.StatusOK), opts); err != nil {
			log.Println(err)
		}
	}
}
