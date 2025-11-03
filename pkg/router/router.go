package router

import (
	"io/fs"
	"log"
	"net/http"

	"borg/pkg/data"
	"borg/web"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	http.Handler
	ds     data.DataStore
	assets fs.FS
}

func NewRouter(appEnv string, ds data.DataStore) *Router {
	assets, err := web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}
	r := &Router{ds: ds, assets: assets}

	h := chi.NewRouter()

	if appEnv == "prod" {
		mux := http.NewServeMux()
		mux.HandleFunc("/", r.handleRoot)
		mux.HandleFunc("/static/", func(w http.ResponseWriter, req *http.Request) {
			http.StripPrefix("/", http.HandlerFunc(r.handleAssets)).ServeHTTP(w, req)
		})
		h.Mount("/", mux)
	}
	h.Route("/api", r.addUserRoute)

	r.Handler = h

	return r
}

func (h *Router) handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, h.assets, "index.html")
}

func (h *Router) handleAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServerFS(h.assets).ServeHTTP(w, r)
}
