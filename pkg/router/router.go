package router

import (
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"net/http"

	"borg/pkg/datastore"
	"borg/pkg/db"
	"borg/web"
)

type Router struct {
	http.Handler
	ds     datastore.DataStore
	assets fs.FS
}

func NewRouter(appEnv string, q *db.Queries) *Router {
	assets, err := web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	ds := datastore.NewDataStore(q)
	r := &Router{ds: ds, assets: assets}

	h := http.NewServeMux()

	if appEnv == "prod" {
		h.HandleFunc("/", r.handleRoot)
		h.HandleFunc("/static/", func(w http.ResponseWriter, req *http.Request) {
			http.StripPrefix("/", http.HandlerFunc(r.handleAssets)).ServeHTTP(w, req)
		})
	}
	h.HandleFunc("/api/", func(w http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println(string(body))
		var user db.AddUserQueryParams
		json.Unmarshal(body, &user)
		log.Println(user)
		if err = r.ds.AddUser(req.Context(), user); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
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
