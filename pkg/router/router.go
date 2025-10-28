package router

import (
	"io/fs"
	"log"
	"net/http"

	"borg/pkg/db/postgres"
	"borg/web"
)

type Router struct {
	http.Handler
	db     *postgres.Queries
	assets fs.FS
}

func New(appEnv string, q *postgres.Queries) *Router {
	assets, err := web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	r := &Router{db: q, assets: assets}

	h := http.NewServeMux()

	if appEnv == "prod" {
		h.HandleFunc("/", r.handleRoot)
		h.HandleFunc("/static/", func(w http.ResponseWriter, req *http.Request) {
			http.StripPrefix("/", http.HandlerFunc(r.handleAssets)).ServeHTTP(w, req)
		})
	}

	r.Handler = h

	return r
}

func (h *Router) handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, h.assets, "index.html")
}

func (h *Router) handleAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServerFS(h.assets).ServeHTTP(w, r)
}
