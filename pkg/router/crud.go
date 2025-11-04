package router

import (
	"context"
	"encoding/json/v2"
	"log"
	"net/http"
	"strconv"

	"borg/pkg/data"

	"github.com/go-chi/chi/v5"
)

type crudHandler[T any, C any, U any] struct {
	repo data.Repository[T, C, U]
	opts json.Options
}

func newCrudHandler[T any, C any, U any](repo data.Repository[T, C, U], opts json.Options) *crudHandler[T, C, U] {
	return &crudHandler[T, C, U]{repo: repo, opts: opts}
}

func (h *crudHandler[T, C, U]) registerRoutes(r chi.Router) {
	r.Get("/", h.getAll)
	r.Post("/", h.create)
	r.Put("/", h.update)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(idCtx)
		r.Get("/", h.getByID)
		r.Delete("/", h.delete)
	})
}

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

func (h *crudHandler[T, C, U]) create(w http.ResponseWriter, r *http.Request) {
	var item C
	if err := json.UnmarshalRead(r.Body, &item, h.opts); err != nil {
		log.Println(err)
		_ = writeError(w, err, http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(r.Context(), item); err != nil {
		log.Println(err)
		_ = writeError(w, err, http.StatusInternalServerError)
		return
	}
	if err := writeSuccess(w, &item, http.StatusCreated, h.opts); err != nil {
		log.Println(err)
	}
}

func (h *crudHandler[T, C, U]) getByID(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(keyID).(int32)
	item, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		_ = writeError(w, err, http.StatusInternalServerError)
		return
	}
	if err = writeSuccess(w, &item, http.StatusOK, h.opts); err != nil {
		log.Println(err)
	}
}

func (h *crudHandler[T, C, U]) getAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.repo.GetAll(r.Context())
	if err != nil {
		log.Println(err)
		_ = writeError(w, err, http.StatusInternalServerError)
		return
	}
	if err = writeSuccess(w, &items, http.StatusOK, h.opts); err != nil {
		log.Println(err)
	}
}

func (h *crudHandler[T, C, U]) delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(keyID).(int32)
	if err := h.repo.Delete(r.Context(), id); err != nil {
		log.Println(err)
		_ = writeError(w, err, http.StatusInternalServerError)
		return
	}
	if err := writeSuccess(w, id, http.StatusOK); err != nil {
		log.Println(err)
	}
}

func (h *crudHandler[T, C, U]) update(w http.ResponseWriter, r *http.Request) {
	var item U
	opts := h.opts
	if err := json.UnmarshalRead(r.Body, &item, opts); err != nil {
		_ = writeError(w, err, http.StatusBadRequest)
		return
	}
	if err := h.repo.Update(r.Context(), item); err != nil {
		log.Println(err)
		_ = writeError(w, err, http.StatusInternalServerError)
		return
	}
	if err := writeSuccess(w, &item, http.StatusOK, opts); err != nil {
		log.Println(err)
	}
}
