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
	h.Use(middleware.Recoverer)
	h.Route("/", func(h chi.Router) {
		h.Get("/*", r.handleRoot)
		h.Get("/static/*", r.handleAssets)
	})
	h.Mount("/api", r.apiRouter())

	r.Handler = h
	return r
}

func (h *Router) handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, h.assets, "index.html")
}

func (h *Router) handleAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServerFS(h.assets).ServeHTTP(w, r)
}

func (h *Router) apiRouter() http.Handler {
	r := chi.NewRouter()
	r.Mount("/user", h.userRouter())
	r.Mount("/post", h.postRouter())
	return r
}

func (h *Router) userRouter() http.Handler {
	repo, opts := h.ds.UserRepository(), h.ds.Opts()
	r := chi.NewRouter()
	r.Post("/", create(repo, opts))
	r.Put("/", update(repo, opts))
	r.Route("/{id}", func(r chi.Router) {
		r.Use(idCtx)
		r.Get("/", getByID(repo, opts))
		r.Delete("/", delete(repo))
	})
	r.With(idCtx).Get("/{id}/posts", getByUserID(h.ds.PostRepository(), opts))
	r.With(idCtx).Get("/{id}/shares", getByUserID(h.ds.ShareRepository(), opts))
	r.With(idCtx).Get("/{id}/likes", getByUserID(h.ds.LikeRepository(), opts))
	r.With(idCtx).Get("/{id}/following", getFollowing(repo, opts))
	r.With(idCtx).Get("/{id}/followers", getFollowers(repo, opts))
	return r
}

func (h *Router) postRouter() http.Handler {
	repo, opts := h.ds.PostRepository(), h.ds.Opts()
	r := chi.NewRouter()
	r.Post("/", create(repo, opts))
	r.Put("/", update(repo, opts))
	r.Route("/{id}", func(r chi.Router) {
		r.Use(idCtx)
		r.Get("/", getByID(repo, opts))
		r.Delete("/", delete(repo))
	})
	r.With(idCtx).Get("/{id}/shares", getByPostID(h.ds.ShareRepository(), opts))
	r.With(idCtx).Get("/{id}/likes", getByPostID(h.ds.LikeRepository(), opts))
	return r
}
