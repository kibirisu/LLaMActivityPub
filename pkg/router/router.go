package router

import (
	"io/fs"
	"log"
	"net/http"

	"borg/pkg/data"
	"borg/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	http.Handler
	ds     data.DataStore
	assets fs.FS
}

func NewRouter(ds data.DataStore) *Router {
	assets, err := web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}
	r := &Router{ds: ds, assets: assets}

	h := chi.NewRouter()
	h.Use(middleware.Logger)
	h.Route("/", func(h chi.Router) {
		h.Get("/*", r.handleRoot)
		h.Get("/static/*", r.handleAssets)
	})
	h.Route("/api", func(h chi.Router) {
		h.Route("/user", r.addUserRoute)
	})

	r.Handler = h
	return r
}

func (h *Router) handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, h.assets, "index.html")
}

func (h *Router) handleAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServerFS(h.assets).ServeHTTP(w, r)
}
